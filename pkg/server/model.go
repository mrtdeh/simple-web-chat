package server

import (
	"api-channel/proto"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

// var grpcLog glog.LoggerV2
type TokenManager struct {
	tokens map[string]*TokenData
	l      sync.RWMutex
}

func (tm *TokenManager) Create(userId uint, username string) string {
	tm.l.Lock()
	defer tm.l.Unlock()
	token := tokenGenerator()
	tm.tokens[token] = &TokenData{
		UserID:     userId,
		Username:   username,
		ExpireTime: time.Now().Add(time.Minute * 5),
	}
	return token
}

func (tm *TokenManager) Get(key string) (*TokenData, error) {
	tm.l.RLock()
	defer tm.l.RUnlock()

	tokenData, ok := tm.tokens[key]
	if !ok {
		return nil, fmt.Errorf("token is not found")
	}
	return tokenData, nil
}

func (tm *TokenManager) Delete(key string) {
	tm.l.Lock()
	defer tm.l.Unlock()
	delete(tm.tokens, key)
}

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
	token   *TokenData
	receive chan struct{}
	close   chan struct{}
	error   chan error
}
type TokenData struct {
	UserID     uint
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
