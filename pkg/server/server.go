package server

import (
	"api-channel/proto"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func (a *Server) Serve() error {
	a.setDefaultOptions()

	if a.Id == "" {
		id := sha256.Sum256([]byte(time.Now().String()))
		serverId := hex.EncodeToString(id[:])
		a.Id = serverId
	}
	a.status = "unknown"
	chatServer := &Server{
		Id:          a.Id,
		Connections: make(map[string]*Connection),
	}

	return chatServer.start(a.Addr)
}

func (s *Server) start(addr string) error {
	gs := grpc.NewServer()
	s.grpcServer = gs

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}
	s.listener = l

	proto.RegisterChatServiceServer(gs, s)
	gs.Serve(l)
	return nil
}

func (a *Server) setDefaultOptions() {
	if a.Options == nil {
		a.Options = &ServerOptions{}
	}
	if a.Options.MaxConnectTry == 0 {
		a.Options.MaxConnectTry = 10
	}
	if a.Options.MaxTimeout == 0 {
		a.Options.MaxTimeout = time.Second * 5
	}
	if a.Options.MaxConnectTryTimeout == 0 {
		a.Options.MaxConnectTryTimeout = time.Second * 1
	}
}

func (a *Server) Status() string {
	return a.status
}
