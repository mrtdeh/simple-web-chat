package server

import (
	"api-channel/proto"
	"fmt"
)

var tm = &TokenManager{}
	
func (s *Server) MessageChannel(pc *proto.Connect, stream proto.ChatService_MessageChannelServer) error {
	// Check authorize user
	token := pc.Token
	ctx := stream.Context()

	// Check token is exist
	tokenData, err := checkToken(token)
	if err != nil {
		return err
	}

	// Check for username is registred or not
	// TODO: search in db for username and check exits and registred
	// TODO: check for username in redis cache
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
	defer func() {
		s.Sessions.Delete(username)
	}()

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

func receiveService(session *Session) {
	for {
		select {
		case <-session.close:
			fmt.Println("session killed: receive channel closed")
			return
		case <-session.OnReceive():
			fmt.Println("received")
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
