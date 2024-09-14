package server

import (
	"api-channel/proto"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

// var grpcLog glog.LoggerV2
var tokens = map[string]TokenData{}

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

type SessionManger struct {
	l        *sync.RWMutex
	sessions map[string]*Session
}

func (s *SessionManger) Add(key string, session *Session) {
	s.l.Lock()
	s.sessions[key] = session
	s.l.Unlock()
}

func (s *SessionManger) Delete(key string) {
	s.l.Lock()
	delete(s.sessions, key)
	s.l.Unlock()
}

func (s *SessionManger) Get(key string) *Session {
	s.l.RLock()
	defer s.l.RUnlock()
	if s, ok := s.sessions[key]; ok {
		return s
	}
	return nil
}
func (s *SessionManger) Close(key string) {
	s.l.Lock()
	defer s.l.Unlock()
	if s, ok := s.sessions[key]; ok {
		closeChannel(s.close)
	}
}

type Session struct {
	stream  proto.ChatService_MessageChannelServer
	token   TokenData
	receive chan struct{}
	close   chan struct{}
	error   chan error
}
type TokenData struct {
	Username   string
	ExpireTime time.Time
}

func (s *Session) FireReceive() {
	close(s.receive)
	s.receive = make(chan struct{})
}

func (s *Session) OnReceive() <-chan struct{} {
	return s.receive
}
