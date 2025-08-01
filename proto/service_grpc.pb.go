// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.15.8
// source: proto/service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChatService_StreamChannel_FullMethodName             = "/proto.ChatService/StreamChannel"
	ChatService_GetMessages_FullMethodName               = "/proto.ChatService/GetMessages"
	ChatService_UpdateLastReadedMessageID_FullMethodName = "/proto.ChatService/UpdateLastReadedMessageID"
	ChatService_AddUser_FullMethodName                   = "/proto.ChatService/AddUser"
	ChatService_AddMember_FullMethodName                 = "/proto.ChatService/AddMember"
	ChatService_Login_FullMethodName                     = "/proto.ChatService/Login"
	ChatService_Ping_FullMethodName                      = "/proto.ChatService/Ping"
	ChatService_CreateChat_FullMethodName                = "/proto.ChatService/CreateChat"
	ChatService_UpdateMemberInfo_FullMethodName          = "/proto.ChatService/UpdateMemberInfo"
	ChatService_CreateGroup_FullMethodName               = "/proto.ChatService/CreateGroup"
	ChatService_JoinGroup_FullMethodName                 = "/proto.ChatService/JoinGroup"
	ChatService_SendMessage_FullMethodName               = "/proto.ChatService/SendMessage"
	ChatService_UploadFile_FullMethodName                = "/proto.ChatService/UploadFile"
	ChatService_GetChats_FullMethodName                  = "/proto.ChatService/GetChats"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	StreamChannel(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamResponse], error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Messages], error)
	// Update Last Readed Message ID Method
	UpdateLastReadedMessageID(ctx context.Context, in *LrmRequest, opts ...grpc.CallOption) (*Close, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*AddMemberResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error)
	UpdateMemberInfo(ctx context.Context, in *MemberInfoRequest, opts ...grpc.CallOption) (*MemberInfoResponse, error)
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error)
	SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	UploadFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error)
	// Just for test
	GetChats(ctx context.Context, in *GetChatsRequest, opts ...grpc.CallOption) (*Chats, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) StreamChannel(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_StreamChannel_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamRequest, StreamResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_StreamChannelClient = grpc.ServerStreamingClient[StreamResponse]

func (c *chatServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Messages], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], ChatService_GetMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetMessagesRequest, Messages]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_GetMessagesClient = grpc.ServerStreamingClient[Messages]

func (c *chatServiceClient) UpdateLastReadedMessageID(ctx context.Context, in *LrmRequest, opts ...grpc.CallOption) (*Close, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Close)
	err := c.cc.Invoke(ctx, ChatService_UpdateLastReadedMessageID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, ChatService_AddUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*AddMemberResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddMemberResponse)
	err := c.cc.Invoke(ctx, ChatService_AddMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, ChatService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, ChatService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateChatResponse)
	err := c.cc.Invoke(ctx, ChatService_CreateChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateMemberInfo(ctx context.Context, in *MemberInfoRequest, opts ...grpc.CallOption) (*MemberInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberInfoResponse)
	err := c.cc.Invoke(ctx, ChatService_UpdateMemberInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, ChatService_CreateGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinGroupResponse)
	err := c.cc.Invoke(ctx, ChatService_JoinGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, ChatService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UploadFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, ChatService_UploadFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChats(ctx context.Context, in *GetChatsRequest, opts ...grpc.CallOption) (*Chats, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Chats)
	err := c.cc.Invoke(ctx, ChatService_GetChats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations should embed UnimplementedChatServiceServer
// for forward compatibility.
type ChatServiceServer interface {
	StreamChannel(*StreamRequest, grpc.ServerStreamingServer[StreamResponse]) error
	GetMessages(*GetMessagesRequest, grpc.ServerStreamingServer[Messages]) error
	// Update Last Readed Message ID Method
	UpdateLastReadedMessageID(context.Context, *LrmRequest) (*Close, error)
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	AddMember(context.Context, *AddMemberRequest) (*AddMemberResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error)
	UpdateMemberInfo(context.Context, *MemberInfoRequest) (*MemberInfoResponse, error)
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	JoinGroup(context.Context, *JoinGroupRequest) (*JoinGroupResponse, error)
	SendMessage(context.Context, *MessageRequest) (*MessageResponse, error)
	UploadFile(context.Context, *FileRequest) (*FileResponse, error)
	// Just for test
	GetChats(context.Context, *GetChatsRequest) (*Chats, error)
}

// UnimplementedChatServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServiceServer struct{}

func (UnimplementedChatServiceServer) StreamChannel(*StreamRequest, grpc.ServerStreamingServer[StreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamChannel not implemented")
}
func (UnimplementedChatServiceServer) GetMessages(*GetMessagesRequest, grpc.ServerStreamingServer[Messages]) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedChatServiceServer) UpdateLastReadedMessageID(context.Context, *LrmRequest) (*Close, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLastReadedMessageID not implemented")
}
func (UnimplementedChatServiceServer) AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedChatServiceServer) AddMember(context.Context, *AddMemberRequest) (*AddMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMember not implemented")
}
func (UnimplementedChatServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedChatServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedChatServiceServer) CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatServiceServer) UpdateMemberInfo(context.Context, *MemberInfoRequest) (*MemberInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMemberInfo not implemented")
}
func (UnimplementedChatServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedChatServiceServer) JoinGroup(context.Context, *JoinGroupRequest) (*JoinGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroup not implemented")
}
func (UnimplementedChatServiceServer) SendMessage(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) UploadFile(context.Context, *FileRequest) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedChatServiceServer) GetChats(context.Context, *GetChatsRequest) (*Chats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (UnimplementedChatServiceServer) testEmbeddedByValue() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_StreamChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).StreamChannel(m, &grpc.GenericServerStream[StreamRequest, StreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_StreamChannelServer = grpc.ServerStreamingServer[StreamResponse]

func _ChatService_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).GetMessages(m, &grpc.GenericServerStream[GetMessagesRequest, Messages]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_GetMessagesServer = grpc.ServerStreamingServer[Messages]

func _ChatService_UpdateLastReadedMessageID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LrmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateLastReadedMessageID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_UpdateLastReadedMessageID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateLastReadedMessageID(ctx, req.(*LrmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_AddMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddMember(ctx, req.(*AddMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateMemberInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateMemberInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_UpdateMemberInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateMemberInfo(ctx, req.(*MemberInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_JoinGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).JoinGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_JoinGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).JoinGroup(ctx, req.(*JoinGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UploadFile(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChats(ctx, req.(*GetChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateLastReadedMessageID",
			Handler:    _ChatService_UpdateLastReadedMessageID_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _ChatService_AddUser_Handler,
		},
		{
			MethodName: "AddMember",
			Handler:    _ChatService_AddMember_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _ChatService_Login_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ChatService_Ping_Handler,
		},
		{
			MethodName: "CreateChat",
			Handler:    _ChatService_CreateChat_Handler,
		},
		{
			MethodName: "UpdateMemberInfo",
			Handler:    _ChatService_UpdateMemberInfo_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _ChatService_CreateGroup_Handler,
		},
		{
			MethodName: "JoinGroup",
			Handler:    _ChatService_JoinGroup_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _ChatService_UploadFile_Handler,
		},
		{
			MethodName: "GetChats",
			Handler:    _ChatService_GetChats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamChannel",
			Handler:       _ChatService_StreamChannel_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMessages",
			Handler:       _ChatService_GetMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}
