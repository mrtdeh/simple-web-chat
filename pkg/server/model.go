package server

import (
	database "api-channel/pkg/db"
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
	db         *database.ChatDatabase
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
	stream       proto.ChatService_NotificationChannelServer
	token        *TokenData
	receive      chan struct{}
	close        chan struct{}
	activeChatId uint32
	error        chan error
}

// Token manager struct
type TokenManager struct {
	tokens map[string]*TokenData
	l      sync.RWMutex
}
type TokenData struct {
	Value      string
	UserID     uint32
	Username   string
	ExpireTime time.Time
}
