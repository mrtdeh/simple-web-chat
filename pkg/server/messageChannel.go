package server

import (
	"api-channel/proto"
	"fmt"
	"log"
	"sync"
	"time"
)

var tm = &TokenManager{
	tokens: make(map[string]*TokenData),
	l:      sync.RWMutex{},
}

func (s *Server) StreamChannel(pc *proto.StreamRequest, stream proto.ChatService_StreamChannelServer) error {
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

	// Prepare new connection info
	session := &Session{
		sl:      &sync.Mutex{},
		stream:  stream,
		token:   tokenData,
		receive: make(chan *proto.StreamResponse, 1),
		close:   make(chan struct{}),
		error:   make(chan error),
	}
	s.Sessions.Add(username, session)
	// Delete the session from map after stream terminate.
	defer func() {
		s.Sessions.Close(username)
		s.Sessions.Delete(username)
	}()

	if err := s.sendChats(username); err != nil {
		log.Fatal("error in sendChats: ", err)
	}
	// This service for realtime checking receive channel for incomming messages
	go s.receiveService(username)

	// This blocking select wait for interrupt signals (kill,error,context done) and close session
	select {
	case <-session.close:
		log.Println("session killed")
		return nil

	case err := <-session.error:
		log.Println("session error : ", err.Error())
		closeChannel(session.close)
		return err

	case <-ctx.Done():
		log.Println("session context done")
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
	return session.stream.Send(&proto.StreamResponse{
		Data: &proto.StreamResponse_Chats{
			Chats: &proto.Chats{
				Data: chats,
			},
		},
	})
}

func (s *Server) noticeUser(username string) error {
	if err := s.sendChats(username); err != nil {
		return fmt.Errorf("error in sendChats: %s", err)
	}
	return nil

}

func (s *Server) receiveService(username string) {
	session := s.Sessions.Get(username)

	t := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-session.close:
			log.Println("session killed: receive channel closed")
			return
		case m, ok := <-session.onReceive():
			if ok {
				session.send(m)
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
