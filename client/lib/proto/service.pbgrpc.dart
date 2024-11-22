//
//  Generated code. Do not modify.
//  source: proto/service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'service.pb.dart' as $0;

export 'service.pb.dart';

@$pb.GrpcServiceName('proto.ChatService')
class ChatServiceClient extends $grpc.Client {
  static final _$messageChannel = $grpc.ClientMethod<$0.MessageChannelRequest, $0.MessageChannelResponse>(
      '/proto.ChatService/MessageChannel',
      ($0.MessageChannelRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.MessageChannelResponse.fromBuffer(value));
  static final _$addUser = $grpc.ClientMethod<$0.AddUserRequest, $0.AddUserResponse>(
      '/proto.ChatService/AddUser',
      ($0.AddUserRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.AddUserResponse.fromBuffer(value));
  static final _$addMember = $grpc.ClientMethod<$0.AddMemberRequest, $0.AddMemberResponse>(
      '/proto.ChatService/AddMember',
      ($0.AddMemberRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.AddMemberResponse.fromBuffer(value));
  static final _$login = $grpc.ClientMethod<$0.LoginRequest, $0.LoginResponse>(
      '/proto.ChatService/Login',
      ($0.LoginRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.LoginResponse.fromBuffer(value));
  static final _$ping = $grpc.ClientMethod<$0.PingRequest, $0.PingResponse>(
      '/proto.ChatService/Ping',
      ($0.PingRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.PingResponse.fromBuffer(value));
  static final _$createChat = $grpc.ClientMethod<$0.CreateChatRequest, $0.CreateChatResponse>(
      '/proto.ChatService/CreateChat',
      ($0.CreateChatRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.CreateChatResponse.fromBuffer(value));
  static final _$updateMemberInfo = $grpc.ClientMethod<$0.MemberInfoRequest, $0.MemberInfoResponse>(
      '/proto.ChatService/UpdateMemberInfo',
      ($0.MemberInfoRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.MemberInfoResponse.fromBuffer(value));
  static final _$createGroup = $grpc.ClientMethod<$0.CreateGroupRequest, $0.CreateGroupResponse>(
      '/proto.ChatService/CreateGroup',
      ($0.CreateGroupRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.CreateGroupResponse.fromBuffer(value));
  static final _$joinGroup = $grpc.ClientMethod<$0.JoinGroupRequest, $0.JoinGroupResponse>(
      '/proto.ChatService/JoinGroup',
      ($0.JoinGroupRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.JoinGroupResponse.fromBuffer(value));
  static final _$sendMessage = $grpc.ClientMethod<$0.MessageRequest, $0.MessageResponse>(
      '/proto.ChatService/SendMessage',
      ($0.MessageRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.MessageResponse.fromBuffer(value));
  static final _$uploadFile = $grpc.ClientMethod<$0.FileRequest, $0.FileResponse>(
      '/proto.ChatService/UploadFile',
      ($0.FileRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.FileResponse.fromBuffer(value));
  static final _$selectChat = $grpc.ClientMethod<$0.SelectChatRequest, $0.EmptyRequest>(
      '/proto.ChatService/SelectChat',
      ($0.SelectChatRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.EmptyRequest.fromBuffer(value));
  static final _$getMessages = $grpc.ClientMethod<$0.GetMessagesRequest, $0.MessagesResponse>(
      '/proto.ChatService/GetMessages',
      ($0.GetMessagesRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.MessagesResponse.fromBuffer(value));
  static final _$getChats = $grpc.ClientMethod<$0.GetChatsRequest, $0.ChatsResponse>(
      '/proto.ChatService/GetChats',
      ($0.GetChatsRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ChatsResponse.fromBuffer(value));

  ChatServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseStream<$0.MessageChannelResponse> messageChannel($0.MessageChannelRequest request, {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$messageChannel, $async.Stream.fromIterable([request]), options: options);
  }

  $grpc.ResponseFuture<$0.AddUserResponse> addUser($0.AddUserRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$addUser, request, options: options);
  }

  $grpc.ResponseFuture<$0.AddMemberResponse> addMember($0.AddMemberRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$addMember, request, options: options);
  }

  $grpc.ResponseFuture<$0.LoginResponse> login($0.LoginRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$login, request, options: options);
  }

  $grpc.ResponseFuture<$0.PingResponse> ping($0.PingRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$ping, request, options: options);
  }

  $grpc.ResponseFuture<$0.CreateChatResponse> createChat($0.CreateChatRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createChat, request, options: options);
  }

  $grpc.ResponseFuture<$0.MemberInfoResponse> updateMemberInfo($0.MemberInfoRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$updateMemberInfo, request, options: options);
  }

  $grpc.ResponseFuture<$0.CreateGroupResponse> createGroup($0.CreateGroupRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createGroup, request, options: options);
  }

  $grpc.ResponseFuture<$0.JoinGroupResponse> joinGroup($0.JoinGroupRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$joinGroup, request, options: options);
  }

  $grpc.ResponseFuture<$0.MessageResponse> sendMessage($0.MessageRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$sendMessage, request, options: options);
  }

  $grpc.ResponseFuture<$0.FileResponse> uploadFile($0.FileRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$uploadFile, request, options: options);
  }

  $grpc.ResponseFuture<$0.EmptyRequest> selectChat($0.SelectChatRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$selectChat, request, options: options);
  }

  $grpc.ResponseFuture<$0.MessagesResponse> getMessages($0.GetMessagesRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getMessages, request, options: options);
  }

  $grpc.ResponseFuture<$0.ChatsResponse> getChats($0.GetChatsRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getChats, request, options: options);
  }
}

@$pb.GrpcServiceName('proto.ChatService')
abstract class ChatServiceBase extends $grpc.Service {
  $core.String get $name => 'proto.ChatService';

  ChatServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.MessageChannelRequest, $0.MessageChannelResponse>(
        'MessageChannel',
        messageChannel_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.MessageChannelRequest.fromBuffer(value),
        ($0.MessageChannelResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.AddUserRequest, $0.AddUserResponse>(
        'AddUser',
        addUser_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.AddUserRequest.fromBuffer(value),
        ($0.AddUserResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.AddMemberRequest, $0.AddMemberResponse>(
        'AddMember',
        addMember_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.AddMemberRequest.fromBuffer(value),
        ($0.AddMemberResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.LoginRequest, $0.LoginResponse>(
        'Login',
        login_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.LoginRequest.fromBuffer(value),
        ($0.LoginResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.PingRequest, $0.PingResponse>(
        'Ping',
        ping_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.PingRequest.fromBuffer(value),
        ($0.PingResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CreateChatRequest, $0.CreateChatResponse>(
        'CreateChat',
        createChat_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateChatRequest.fromBuffer(value),
        ($0.CreateChatResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.MemberInfoRequest, $0.MemberInfoResponse>(
        'UpdateMemberInfo',
        updateMemberInfo_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.MemberInfoRequest.fromBuffer(value),
        ($0.MemberInfoResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CreateGroupRequest, $0.CreateGroupResponse>(
        'CreateGroup',
        createGroup_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateGroupRequest.fromBuffer(value),
        ($0.CreateGroupResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.JoinGroupRequest, $0.JoinGroupResponse>(
        'JoinGroup',
        joinGroup_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.JoinGroupRequest.fromBuffer(value),
        ($0.JoinGroupResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.MessageRequest, $0.MessageResponse>(
        'SendMessage',
        sendMessage_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.MessageRequest.fromBuffer(value),
        ($0.MessageResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.FileRequest, $0.FileResponse>(
        'UploadFile',
        uploadFile_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.FileRequest.fromBuffer(value),
        ($0.FileResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.SelectChatRequest, $0.EmptyRequest>(
        'SelectChat',
        selectChat_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.SelectChatRequest.fromBuffer(value),
        ($0.EmptyRequest value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetMessagesRequest, $0.MessagesResponse>(
        'GetMessages',
        getMessages_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetMessagesRequest.fromBuffer(value),
        ($0.MessagesResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetChatsRequest, $0.ChatsResponse>(
        'GetChats',
        getChats_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetChatsRequest.fromBuffer(value),
        ($0.ChatsResponse value) => value.writeToBuffer()));
  }

  $async.Stream<$0.MessageChannelResponse> messageChannel_Pre($grpc.ServiceCall call, $async.Future<$0.MessageChannelRequest> request) async* {
    yield* messageChannel(call, await request);
  }

  $async.Future<$0.AddUserResponse> addUser_Pre($grpc.ServiceCall call, $async.Future<$0.AddUserRequest> request) async {
    return addUser(call, await request);
  }

  $async.Future<$0.AddMemberResponse> addMember_Pre($grpc.ServiceCall call, $async.Future<$0.AddMemberRequest> request) async {
    return addMember(call, await request);
  }

  $async.Future<$0.LoginResponse> login_Pre($grpc.ServiceCall call, $async.Future<$0.LoginRequest> request) async {
    return login(call, await request);
  }

  $async.Future<$0.PingResponse> ping_Pre($grpc.ServiceCall call, $async.Future<$0.PingRequest> request) async {
    return ping(call, await request);
  }

  $async.Future<$0.CreateChatResponse> createChat_Pre($grpc.ServiceCall call, $async.Future<$0.CreateChatRequest> request) async {
    return createChat(call, await request);
  }

  $async.Future<$0.MemberInfoResponse> updateMemberInfo_Pre($grpc.ServiceCall call, $async.Future<$0.MemberInfoRequest> request) async {
    return updateMemberInfo(call, await request);
  }

  $async.Future<$0.CreateGroupResponse> createGroup_Pre($grpc.ServiceCall call, $async.Future<$0.CreateGroupRequest> request) async {
    return createGroup(call, await request);
  }

  $async.Future<$0.JoinGroupResponse> joinGroup_Pre($grpc.ServiceCall call, $async.Future<$0.JoinGroupRequest> request) async {
    return joinGroup(call, await request);
  }

  $async.Future<$0.MessageResponse> sendMessage_Pre($grpc.ServiceCall call, $async.Future<$0.MessageRequest> request) async {
    return sendMessage(call, await request);
  }

  $async.Future<$0.FileResponse> uploadFile_Pre($grpc.ServiceCall call, $async.Future<$0.FileRequest> request) async {
    return uploadFile(call, await request);
  }

  $async.Future<$0.EmptyRequest> selectChat_Pre($grpc.ServiceCall call, $async.Future<$0.SelectChatRequest> request) async {
    return selectChat(call, await request);
  }

  $async.Future<$0.MessagesResponse> getMessages_Pre($grpc.ServiceCall call, $async.Future<$0.GetMessagesRequest> request) async {
    return getMessages(call, await request);
  }

  $async.Future<$0.ChatsResponse> getChats_Pre($grpc.ServiceCall call, $async.Future<$0.GetChatsRequest> request) async {
    return getChats(call, await request);
  }

  $async.Stream<$0.MessageChannelResponse> messageChannel($grpc.ServiceCall call, $0.MessageChannelRequest request);
  $async.Future<$0.AddUserResponse> addUser($grpc.ServiceCall call, $0.AddUserRequest request);
  $async.Future<$0.AddMemberResponse> addMember($grpc.ServiceCall call, $0.AddMemberRequest request);
  $async.Future<$0.LoginResponse> login($grpc.ServiceCall call, $0.LoginRequest request);
  $async.Future<$0.PingResponse> ping($grpc.ServiceCall call, $0.PingRequest request);
  $async.Future<$0.CreateChatResponse> createChat($grpc.ServiceCall call, $0.CreateChatRequest request);
  $async.Future<$0.MemberInfoResponse> updateMemberInfo($grpc.ServiceCall call, $0.MemberInfoRequest request);
  $async.Future<$0.CreateGroupResponse> createGroup($grpc.ServiceCall call, $0.CreateGroupRequest request);
  $async.Future<$0.JoinGroupResponse> joinGroup($grpc.ServiceCall call, $0.JoinGroupRequest request);
  $async.Future<$0.MessageResponse> sendMessage($grpc.ServiceCall call, $0.MessageRequest request);
  $async.Future<$0.FileResponse> uploadFile($grpc.ServiceCall call, $0.FileRequest request);
  $async.Future<$0.EmptyRequest> selectChat($grpc.ServiceCall call, $0.SelectChatRequest request);
  $async.Future<$0.MessagesResponse> getMessages($grpc.ServiceCall call, $0.GetMessagesRequest request);
  $async.Future<$0.ChatsResponse> getChats($grpc.ServiceCall call, $0.GetChatsRequest request);
}
