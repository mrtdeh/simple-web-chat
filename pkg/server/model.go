package server

import (
	"api-channel/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

var grpcLog glog.LoggerV2

type GRPCServer struct {
	grpcServer *grpc.Server
	listener   net.Listener

	Id          string
	Connections map[string]*Connection
}

type Connection struct {
	stream proto.ChatService_CreateStreamServer
	user   *proto.User
	active bool
	error  chan error
}

func NewChat(sid string) *GRPCServer {
	server := &GRPCServer{
		Id:          sid,
		Connections: make(map[string]*Connection),
	}
	return server
}

func (s *GRPCServer) Serve(addr string) error {
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
