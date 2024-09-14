package server

import (
	"api-channel/proto"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type Config struct {
	Port    int
	Options ServerOptions
}

func NewServer(cnf ...Config) *Server {
	c := cnf[0]

	id := sha256.Sum256([]byte(time.Now().String()))
	serverId := hex.EncodeToString(id[:])

	chatServer := &Server{
		Id:      serverId,
		Addr:    fmt.Sprintf("127.0.0.1:%d", c.Port),
		status:  "unknown",
		Options: &c.Options,
		Sessions: &SessionManger{
			l:        &sync.RWMutex{},
			sessions: make(map[string]*Session),
		},
	}

	chatServer.setDefaultOptions()
	return chatServer
}

func (s *Server) Serve() error {
	gs := grpc.NewServer()
	s.grpcServer = gs

	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}
	s.listener = l

	proto.RegisterChatServiceServer(gs, s)
	gs.Serve(l)
	return nil
}

func (a *Server) setDefaultOptions() {
	if a.Addr == "" {
		a.Addr = "127.0.0.1:8082"
	}
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
