package server

import (
	"api-channel/proto"
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

var tm = &TokenManager{
	tokens: make(map[string]*TokenData),
	l:      sync.RWMutex{},
}

func (s *Server) MessageChannel(pc *proto.MessageChannelRequest, stream proto.ChatService_MessageChannelServer) error {
	// Check authorize user
	token := pc.Token
	ctx := stream.Context()

	// Check token is exist
	tokenData, err := checkToken(token)
	if err != nil {
		return err
	}
	username := tokenData.Username
	// Deny requested username if it assigned and actived by another client
	s.Sessions.Close(username)
	// Prepare new connection info
	session := &Session{
		stream:  stream,
		token:   tokenData,
		receive: make(chan struct{}),
		close:   make(chan struct{}),
		error:   make(chan error),
	}
	s.Sessions.Add(username, session)
	// Delete the session from map after stream terminate.
	defer s.Sessions.Delete(username)

	if err := s.sendChats(username); err != nil {
		log.Fatal("error in sendChats: ", err)
	}
	// This service for realtime checking receive channel for incomming messages
	go s.receiveService(username)

	// This blocking select wait for interrupt signals (kill,error,context done) and close session
	select {
	case <-session.close:
		fmt.Println("session killed")
		return nil

	case err := <-session.error:
		fmt.Println("session error : ", err.Error())
		closeChannel(session.close)
		return err

	case <-ctx.Done():
		fmt.Println("session context done")
		closeChannel(session.close)
		return nil
	}
}

// =================================================================

func (s *Server) sendChats(username string) error {
	session := s.Sessions.Get(username)
	chats, err := s.getChats(session.token.UserID)
	if err != nil {
		return err
	}
	return session.stream.Send(&proto.MessageChannelResponse{
		Data: &proto.MessageChannelResponse_Chats{
			Chats: &proto.ChatsResponse{
				Data: chats,
			},
		},
	})
}

func (s *Server) sendMessages(username string) error {
	session := s.Sessions.Get(username)
	res, err := s.GetMessages(context.Background(), &proto.GetMessagesRequest{
		ChatId: session.activeChatId,
	})
	if err != nil {
		return err
	}
	return session.stream.Send(&proto.MessageChannelResponse{
		Data: &proto.MessageChannelResponse_Messages{
			Messages: &proto.MessagesResponse{
				Data: res.Data,
			},
		},
	})
}

func (s *Server) receiveService(username string) {
	session := s.Sessions.Get(username)

	t := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-session.close:
			fmt.Println("session killed: receive channel closed")
			return
		case <-session.OnReceive():
			if err := s.sendMessages(session.token.Username); err != nil {
				fmt.Println("error in send messages from channel : ", err.Error())
			}
		case <-t.C:
			// check for received messages from db

			// send messages in stream

		}
	}
}

func closeChannel(c chan struct{}) {
	select {
	case <-c:
	default:
		close(c)
	}
}
