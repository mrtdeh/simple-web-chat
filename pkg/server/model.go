package server

import (
	"api-channel/proto"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

// Chat Server struct
type Server struct {
	Id      string
	Addr    string
	status  string
	Options *ServerOptions

	grpcServer *grpc.Server
	listener   net.Listener
	Sessions   *SessionManger
}
type ServerOptions struct {
	MaxTimeout           time.Duration
	MaxConnectTry        int
	MaxConnectTryTimeout time.Duration
}

// Session manager struct
type SessionManger struct {
	l        *sync.RWMutex
	sessions map[string]*Session
}
type Session struct {
	stream  proto.ChatService_MessageChannelServer
	token   *TokenData
	receive chan struct{}
	close   chan struct{}
	error   chan error
}

// Token manager struct
type TokenManager struct {
	tokens map[string]*TokenData
	l      sync.RWMutex
}
type TokenData struct {
	UserID     uint
	Username   string
	ExpireTime time.Time
}
