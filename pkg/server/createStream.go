package server

import (
	"api-channel/proto"
	"fmt"
	"time"
)

type TokenData struct {
	Username   string
	ExpireTime time.Time
}

var tokens = map[string]TokenData{}

func (s *Server) MessageChannel(pc *proto.Connect, stream proto.ChatService_MessageChannelServer) error {
	// Check authorize user
	token := pc.Token
	ctx := stream.Context()

	// Check token expiration
	t, ok := tokens[token]
	if !ok {
		return fmt.Errorf("token is not found")
	}

	now := time.Now().Unix()
	tokenTime := t.ExpireTime.Unix()
	if ok && tokenTime < now {
		return fmt.Errorf("token is expired")
	}

	// Check for username is registred or not
	// TODO: search in db for username and check exits and registred
	// TODO: check for username in redis cache
	username := t.Username

	// Deny requested username if it assigned and actived by another client
	if s, ok := s.Sessions[username]; ok {
		close(s.kill)
	}
	// Prepare new connection info
	session := &Session{
		stream:  stream,
		token:   token,
		active:  true,
		receive: make(chan struct{}),
		kill:    make(chan struct{}),
		error:   make(chan error),
	}

	s.Sessions[username] = session
	session.status = "connected"

	select {
	case err := <-session.error:
		fmt.Println("session error : ", err.Error())
		return err
	case <-session.kill:
		fmt.Println("session killed")
		return nil
	case <-ctx.Done():
		fmt.Println("session context done")
		return nil
	}
}
