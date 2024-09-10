package server

import (
	"api-channel/proto"
	"fmt"
)

func (s *Server) CreateStream(pconn *proto.Connect, stream proto.ChatService_CreateStreamServer) error {
	// Prepare new connection info
	conn := &Connection{
		stream: stream,
		user:   pconn.User,
		active: true,
		error:  make(chan error),
	}

	username := pconn.User.Username

	// Check for username is registred or not
	// TODO: search in db for username and check exits and registred
	// TODO: check for username in redis cache

	// Deny requested username if it assigned and actived by another client
	if c, ok := s.Connections[username]; ok {
		err := fmt.Errorf("this username(%s) has already been used", pconn.User.Username)
		if c != nil && c.active {
			stream.Send(errorResponse(err))
			return err
		}
	}

	defer func() {
		// Close connection and remove it from s.Connections
		conn.status = "closed"
	}()

	// Add user to connection list
	s.Connections[username] = conn
	conn.status = "connected"

	return <-conn.error
}
