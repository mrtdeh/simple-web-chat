package server

import (
	"api-channel/proto"
	"fmt"
)

func (s *GRPCServer) CreateStream(pconn *proto.Connect, stream proto.ChatService_CreateStreamServer) error {
	// prepare new connection info
	conn := &Connection{
		stream: stream,
		user:   pconn.User,
		active: true,
		error:  make(chan error),
	}

	username := pconn.User.Username
	// deny requested username if it assigned and actived by another client
	if c, ok := s.Connections[username]; ok {
		err := fmt.Errorf("this username(%s) has already been used", pconn.User.Username)
		if c != nil && c.active {
			stream.Send(errorResponse(err))
			return err
		}
	}

	// add user to connection list
	s.Connections[username] = conn

	return <-conn.error
}
