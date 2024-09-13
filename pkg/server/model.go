package server

import (
	"api-channel/proto"
	"net"
	"time"

	"google.golang.org/grpc"
)

// var grpcLog glog.LoggerV2

type Server struct {
	Id      string
	Addr    string
	status  string
	Options *ServerOptions

	grpcServer *grpc.Server
	listener   net.Listener
	Sessions   map[string]*Session
}

type ServerOptions struct {
	MaxTimeout           time.Duration
	MaxConnectTry        int
	MaxConnectTryTimeout time.Duration
}

type Session struct {
	stream proto.ChatService_MessageChannelServer
	// user   *proto.User
	token   string
	active  bool
	status  string
	receive chan struct{}
	kill    chan struct{}
	error   chan error
}
