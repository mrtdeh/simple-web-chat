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

	grpcServer  *grpc.Server
	listener    net.Listener
	Connections map[string]*Connection
}

type ServerOptions struct {
	MaxTimeout           time.Duration
	MaxConnectTry        int
	MaxConnectTryTimeout time.Duration
}

type Connection struct {
	stream proto.ChatService_CreateStreamServer
	user   *proto.User
	active bool
	status string
	error  chan error
}
