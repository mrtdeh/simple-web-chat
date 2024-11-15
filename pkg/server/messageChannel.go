package server

import (
	"api-channel/proto"
	"context"
	"fmt"
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

	go sendChats(ctx, s, session)
	// This service for realtime checking receive channel for incomming messages
	go receiveService(session)

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

func sendChats(ctx context.Context, s *Server, session *Session) error {
	res, err := s.GetChats(ctx, &proto.GetChatsRequest{
		Token: session.token.Value,
	})
	if err != nil {
		return err
	}
	fmt.Println("res.Data : ", len(res.Data))
	return session.stream.Send(&proto.MessageChannelResponse{
		Data: &proto.MessageChannelResponse_Chats{
			Chats: &proto.ChatsResponse{
				Data: res.Data,
			},
		},
	})
}

func receiveService(session *Session) {
	t := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-session.close:
			fmt.Println("session killed: receive channel closed")
			return
		case <-session.OnReceive():
			fmt.Println("received")
		case <-t.C:
			// store the last readed message id to start from it later
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
