//
//  Generated code. Do not modify.
//  source: proto/service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use closeDescriptor instead')
const Close$json = {
  '1': 'Close',
};

/// Descriptor for `Close`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List closeDescriptor = $convert.base64Decode(
    'CgVDbG9zZQ==');

@$core.Deprecated('Use errorResponseDescriptor instead')
const ErrorResponse$json = {
  '1': 'ErrorResponse',
  '2': [
    {'1': 'message', '3': 1, '4': 1, '5': 9, '10': 'message'},
    {'1': 'code', '3': 2, '4': 1, '5': 5, '10': 'code'},
  ],
};

/// Descriptor for `ErrorResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List errorResponseDescriptor = $convert.base64Decode(
    'Cg1FcnJvclJlc3BvbnNlEhgKB21lc3NhZ2UYASABKAlSB21lc3NhZ2USEgoEY29kZRgCIAEoBV'
    'IEY29kZQ==');

@$core.Deprecated('Use pingRequestDescriptor instead')
const PingRequest$json = {
  '1': 'PingRequest',
};

/// Descriptor for `PingRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pingRequestDescriptor = $convert.base64Decode(
    'CgtQaW5nUmVxdWVzdA==');

@$core.Deprecated('Use pingResponseDescriptor instead')
const PingResponse$json = {
  '1': 'PingResponse',
};

/// Descriptor for `PingResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pingResponseDescriptor = $convert.base64Decode(
    'CgxQaW5nUmVzcG9uc2U=');

@$core.Deprecated('Use emptyRequestDescriptor instead')
const EmptyRequest$json = {
  '1': 'EmptyRequest',
};

/// Descriptor for `EmptyRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List emptyRequestDescriptor = $convert.base64Decode(
    'CgxFbXB0eVJlcXVlc3Q=');

@$core.Deprecated('Use loginRequestDescriptor instead')
const LoginRequest$json = {
  '1': 'LoginRequest',
  '2': [
    {'1': 'username', '3': 1, '4': 1, '5': 9, '10': 'username'},
    {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
  ],
};

/// Descriptor for `LoginRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginRequestDescriptor = $convert.base64Decode(
    'CgxMb2dpblJlcXVlc3QSGgoIdXNlcm5hbWUYASABKAlSCHVzZXJuYW1lEhoKCHBhc3N3b3JkGA'
    'IgASgJUghwYXNzd29yZA==');

@$core.Deprecated('Use loginResponseDescriptor instead')
const LoginResponse$json = {
  '1': 'LoginResponse',
  '2': [
    {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
    {'1': 'userId', '3': 2, '4': 1, '5': 13, '10': 'userId'},
  ],
};

/// Descriptor for `LoginResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginResponseDescriptor = $convert.base64Decode(
    'Cg1Mb2dpblJlc3BvbnNlEhQKBXRva2VuGAEgASgJUgV0b2tlbhIWCgZ1c2VySWQYAiABKA1SBn'
    'VzZXJJZA==');

@$core.Deprecated('Use messageStreamDescriptor instead')
const MessageStream$json = {
  '1': 'MessageStream',
  '2': [
    {'1': 'msg_id', '3': 1, '4': 1, '5': 9, '10': 'msgId'},
    {'1': 'data', '3': 2, '4': 1, '5': 11, '6': '.proto.StreamData', '10': 'data'},
  ],
};

/// Descriptor for `MessageStream`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageStreamDescriptor = $convert.base64Decode(
    'Cg1NZXNzYWdlU3RyZWFtEhUKBm1zZ19pZBgBIAEoCVIFbXNnSWQSJQoEZGF0YRgCIAEoCzIRLn'
    'Byb3RvLlN0cmVhbURhdGFSBGRhdGE=');

@$core.Deprecated('Use streamDataDescriptor instead')
const StreamData$json = {
  '1': 'StreamData',
  '2': [
    {'1': 'header', '3': 1, '4': 1, '5': 11, '6': '.proto.StreamData.Header', '9': 0, '10': 'header'},
    {'1': 'text', '3': 2, '4': 1, '5': 11, '6': '.proto.StreamData.Text', '9': 0, '10': 'text'},
    {'1': 'file', '3': 3, '4': 1, '5': 11, '6': '.proto.StreamData.File', '9': 0, '10': 'file'},
  ],
  '3': [StreamData_Header$json, StreamData_Text$json, StreamData_File$json],
  '8': [
    {'1': 'data'},
  ],
};

@$core.Deprecated('Use streamDataDescriptor instead')
const StreamData_Header$json = {
  '1': 'Header',
  '2': [
    {'1': 'timestamp', '3': 3, '4': 1, '5': 9, '10': 'timestamp'},
    {'1': 'from', '3': 4, '4': 1, '5': 9, '10': 'from'},
    {'1': 'type', '3': 5, '4': 1, '5': 9, '10': 'type'},
  ],
};

@$core.Deprecated('Use streamDataDescriptor instead')
const StreamData_Text$json = {
  '1': 'Text',
  '2': [
    {'1': 'content', '3': 1, '4': 1, '5': 9, '10': 'content'},
  ],
};

@$core.Deprecated('Use streamDataDescriptor instead')
const StreamData_File$json = {
  '1': 'File',
  '2': [
    {'1': 'info', '3': 1, '4': 1, '5': 11, '6': '.proto.FileInfo', '9': 0, '10': 'info'},
    {'1': 'chunk', '3': 2, '4': 1, '5': 12, '9': 0, '10': 'chunk'},
    {'1': 'end', '3': 3, '4': 1, '5': 8, '9': 0, '10': 'end'},
  ],
  '8': [
    {'1': 'file'},
  ],
};

/// Descriptor for `StreamData`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List streamDataDescriptor = $convert.base64Decode(
    'CgpTdHJlYW1EYXRhEjIKBmhlYWRlchgBIAEoCzIYLnByb3RvLlN0cmVhbURhdGEuSGVhZGVySA'
    'BSBmhlYWRlchIsCgR0ZXh0GAIgASgLMhYucHJvdG8uU3RyZWFtRGF0YS5UZXh0SABSBHRleHQS'
    'LAoEZmlsZRgDIAEoCzIWLnByb3RvLlN0cmVhbURhdGEuRmlsZUgAUgRmaWxlGk4KBkhlYWRlch'
    'IcCgl0aW1lc3RhbXAYAyABKAlSCXRpbWVzdGFtcBISCgRmcm9tGAQgASgJUgRmcm9tEhIKBHR5'
    'cGUYBSABKAlSBHR5cGUaIAoEVGV4dBIYCgdjb250ZW50GAEgASgJUgdjb250ZW50GmEKBEZpbG'
    'USJQoEaW5mbxgBIAEoCzIPLnByb3RvLkZpbGVJbmZvSABSBGluZm8SFgoFY2h1bmsYAiABKAxI'
    'AFIFY2h1bmsSEgoDZW5kGAMgASgISABSA2VuZEIGCgRmaWxlQgYKBGRhdGE=');

@$core.Deprecated('Use messageRequestDescriptor instead')
const MessageRequest$json = {
  '1': 'MessageRequest',
  '2': [
    {'1': 'token', '3': 2, '4': 1, '5': 9, '10': 'token'},
    {'1': 'content', '3': 1, '4': 1, '5': 9, '10': 'content'},
    {'1': 'type', '3': 3, '4': 1, '5': 9, '10': 'type'},
    {'1': 'chat_id', '3': 4, '4': 1, '5': 13, '10': 'chatId'},
    {'1': 'repliedMessages', '3': 5, '4': 3, '5': 11, '6': '.proto.MessageRequest.RepliedMessage', '10': 'repliedMessages'},
  ],
  '3': [MessageRequest_RepliedMessage$json],
};

@$core.Deprecated('Use messageRequestDescriptor instead')
const MessageRequest_RepliedMessage$json = {
  '1': 'RepliedMessage',
  '2': [
    {'1': 'messageId', '3': 1, '4': 1, '5': 13, '10': 'messageId'},
    {'1': 'wholeMessage', '3': 2, '4': 1, '5': 8, '10': 'wholeMessage'},
    {'1': 'selectedAttachments', '3': 3, '4': 3, '5': 13, '10': 'selectedAttachments'},
  ],
};

/// Descriptor for `MessageRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageRequestDescriptor = $convert.base64Decode(
    'Cg5NZXNzYWdlUmVxdWVzdBIUCgV0b2tlbhgCIAEoCVIFdG9rZW4SGAoHY29udGVudBgBIAEoCV'
    'IHY29udGVudBISCgR0eXBlGAMgASgJUgR0eXBlEhcKB2NoYXRfaWQYBCABKA1SBmNoYXRJZBJO'
    'Cg9yZXBsaWVkTWVzc2FnZXMYBSADKAsyJC5wcm90by5NZXNzYWdlUmVxdWVzdC5SZXBsaWVkTW'
    'Vzc2FnZVIPcmVwbGllZE1lc3NhZ2VzGoQBCg5SZXBsaWVkTWVzc2FnZRIcCgltZXNzYWdlSWQY'
    'ASABKA1SCW1lc3NhZ2VJZBIiCgx3aG9sZU1lc3NhZ2UYAiABKAhSDHdob2xlTWVzc2FnZRIwCh'
    'NzZWxlY3RlZEF0dGFjaG1lbnRzGAMgAygNUhNzZWxlY3RlZEF0dGFjaG1lbnRz');

@$core.Deprecated('Use messageResponseDescriptor instead')
const MessageResponse$json = {
  '1': 'MessageResponse',
  '2': [
    {'1': 'message_id', '3': 1, '4': 1, '5': 13, '10': 'messageId'},
  ],
};

/// Descriptor for `MessageResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageResponseDescriptor = $convert.base64Decode(
    'Cg9NZXNzYWdlUmVzcG9uc2USHQoKbWVzc2FnZV9pZBgBIAEoDVIJbWVzc2FnZUlk');

@$core.Deprecated('Use fileInfoDescriptor instead')
const FileInfo$json = {
  '1': 'FileInfo',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    {'1': 'size', '3': 2, '4': 1, '5': 13, '10': 'size'},
    {'1': 'type', '3': 3, '4': 1, '5': 9, '10': 'type'},
  ],
};

/// Descriptor for `FileInfo`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List fileInfoDescriptor = $convert.base64Decode(
    'CghGaWxlSW5mbxISCgRuYW1lGAEgASgJUgRuYW1lEhIKBHNpemUYAiABKA1SBHNpemUSEgoEdH'
    'lwZRgDIAEoCVIEdHlwZQ==');

@$core.Deprecated('Use fileRequestDescriptor instead')
const FileRequest$json = {
  '1': 'FileRequest',
  '2': [
    {'1': 'request_id', '3': 5, '4': 1, '5': 9, '10': 'requestId'},
    {'1': 'message_id', '3': 1, '4': 1, '5': 13, '10': 'messageId'},
    {'1': 'info', '3': 2, '4': 1, '5': 11, '6': '.proto.FileInfo', '9': 0, '10': 'info'},
    {'1': 'chunk', '3': 3, '4': 1, '5': 9, '9': 0, '10': 'chunk'},
    {'1': 'done', '3': 4, '4': 1, '5': 8, '9': 0, '10': 'done'},
  ],
  '8': [
    {'1': 'data'},
  ],
};

/// Descriptor for `FileRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List fileRequestDescriptor = $convert.base64Decode(
    'CgtGaWxlUmVxdWVzdBIdCgpyZXF1ZXN0X2lkGAUgASgJUglyZXF1ZXN0SWQSHQoKbWVzc2FnZV'
    '9pZBgBIAEoDVIJbWVzc2FnZUlkEiUKBGluZm8YAiABKAsyDy5wcm90by5GaWxlSW5mb0gAUgRp'
    'bmZvEhYKBWNodW5rGAMgASgJSABSBWNodW5rEhQKBGRvbmUYBCABKAhIAFIEZG9uZUIGCgRkYX'
    'Rh');

@$core.Deprecated('Use fileResponseDescriptor instead')
const FileResponse$json = {
  '1': 'FileResponse',
  '2': [
    {'1': 'status', '3': 1, '4': 1, '5': 9, '10': 'status'},
    {'1': 'attachment_id', '3': 2, '4': 1, '5': 13, '10': 'attachmentId'},
  ],
};

/// Descriptor for `FileResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List fileResponseDescriptor = $convert.base64Decode(
    'CgxGaWxlUmVzcG9uc2USFgoGc3RhdHVzGAEgASgJUgZzdGF0dXMSIwoNYXR0YWNobWVudF9pZB'
    'gCIAEoDVIMYXR0YWNobWVudElk');

@$core.Deprecated('Use addUserRequestDescriptor instead')
const AddUserRequest$json = {
  '1': 'AddUserRequest',
  '2': [
    {'1': 'username', '3': 1, '4': 1, '5': 9, '10': 'username'},
    {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
    {'1': 'displayName', '3': 3, '4': 1, '5': 9, '10': 'displayName'},
    {'1': 'phoneNumber', '3': 4, '4': 1, '5': 9, '10': 'phoneNumber'},
    {'1': 'email', '3': 5, '4': 1, '5': 9, '10': 'email'},
  ],
};

/// Descriptor for `AddUserRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addUserRequestDescriptor = $convert.base64Decode(
    'Cg5BZGRVc2VyUmVxdWVzdBIaCgh1c2VybmFtZRgBIAEoCVIIdXNlcm5hbWUSGgoIcGFzc3dvcm'
    'QYAiABKAlSCHBhc3N3b3JkEiAKC2Rpc3BsYXlOYW1lGAMgASgJUgtkaXNwbGF5TmFtZRIgCgtw'
    'aG9uZU51bWJlchgEIAEoCVILcGhvbmVOdW1iZXISFAoFZW1haWwYBSABKAlSBWVtYWls');

@$core.Deprecated('Use addUserResponseDescriptor instead')
const AddUserResponse$json = {
  '1': 'AddUserResponse',
  '2': [
    {'1': 'userId', '3': 1, '4': 1, '5': 13, '10': 'userId'},
    {'1': 'status', '3': 2, '4': 1, '5': 9, '10': 'status'},
  ],
};

/// Descriptor for `AddUserResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addUserResponseDescriptor = $convert.base64Decode(
    'Cg9BZGRVc2VyUmVzcG9uc2USFgoGdXNlcklkGAEgASgNUgZ1c2VySWQSFgoGc3RhdHVzGAIgAS'
    'gJUgZzdGF0dXM=');

@$core.Deprecated('Use createChatRequestDescriptor instead')
const CreateChatRequest$json = {
  '1': 'CreateChatRequest',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    {'1': 'type', '3': 2, '4': 1, '5': 9, '10': 'type'},
  ],
};

/// Descriptor for `CreateChatRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createChatRequestDescriptor = $convert.base64Decode(
    'ChFDcmVhdGVDaGF0UmVxdWVzdBISCgRuYW1lGAEgASgJUgRuYW1lEhIKBHR5cGUYAiABKAlSBH'
    'R5cGU=');

@$core.Deprecated('Use createChatResponseDescriptor instead')
const CreateChatResponse$json = {
  '1': 'CreateChatResponse',
  '2': [
    {'1': 'chat_id', '3': 1, '4': 1, '5': 13, '10': 'chatId'},
  ],
};

/// Descriptor for `CreateChatResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createChatResponseDescriptor = $convert.base64Decode(
    'ChJDcmVhdGVDaGF0UmVzcG9uc2USFwoHY2hhdF9pZBgBIAEoDVIGY2hhdElk');

@$core.Deprecated('Use notificationRequestDescriptor instead')
const NotificationRequest$json = {
  '1': 'NotificationRequest',
  '2': [
    {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `NotificationRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List notificationRequestDescriptor = $convert.base64Decode(
    'ChNOb3RpZmljYXRpb25SZXF1ZXN0EhQKBXRva2VuGAEgASgJUgV0b2tlbg==');

@$core.Deprecated('Use notificationResponseDescriptor instead')
const NotificationResponse$json = {
  '1': 'NotificationResponse',
  '2': [
    {'1': 'messages', '3': 1, '4': 1, '5': 11, '6': '.proto.MessagesResponse', '9': 0, '10': 'messages'},
    {'1': 'chats', '3': 2, '4': 1, '5': 11, '6': '.proto.ChatsResponse', '9': 0, '10': 'chats'},
  ],
  '8': [
    {'1': 'data'},
  ],
};

/// Descriptor for `NotificationResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List notificationResponseDescriptor = $convert.base64Decode(
    'ChROb3RpZmljYXRpb25SZXNwb25zZRI1CghtZXNzYWdlcxgBIAEoCzIXLnByb3RvLk1lc3NhZ2'
    'VzUmVzcG9uc2VIAFIIbWVzc2FnZXMSLAoFY2hhdHMYAiABKAsyFC5wcm90by5DaGF0c1Jlc3Bv'
    'bnNlSABSBWNoYXRzQgYKBGRhdGE=');

@$core.Deprecated('Use getMessagesRequestDescriptor instead')
const GetMessagesRequest$json = {
  '1': 'GetMessagesRequest',
  '2': [
    {'1': 'chat_id', '3': 1, '4': 1, '5': 13, '10': 'chatId'},
    {'1': 'readed_msg_id', '3': 2, '4': 1, '5': 13, '10': 'readedMsgId'},
    {'1': 'direction', '3': 3, '4': 1, '5': 5, '10': 'direction'},
    {'1': 'count', '3': 4, '4': 1, '5': 5, '10': 'count'},
  ],
};

/// Descriptor for `GetMessagesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMessagesRequestDescriptor = $convert.base64Decode(
    'ChJHZXRNZXNzYWdlc1JlcXVlc3QSFwoHY2hhdF9pZBgBIAEoDVIGY2hhdElkEiIKDXJlYWRlZF'
    '9tc2dfaWQYAiABKA1SC3JlYWRlZE1zZ0lkEhwKCWRpcmVjdGlvbhgDIAEoBVIJZGlyZWN0aW9u'
    'EhQKBWNvdW50GAQgASgFUgVjb3VudA==');

@$core.Deprecated('Use followChatRequestDescriptor instead')
const FollowChatRequest$json = {
  '1': 'FollowChatRequest',
  '2': [
    {'1': 'chat_id', '3': 1, '4': 1, '5': 13, '10': 'chatId'},
    {'1': 'last_msg_id', '3': 2, '4': 1, '5': 13, '10': 'lastMsgId'},
  ],
};

/// Descriptor for `FollowChatRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List followChatRequestDescriptor = $convert.base64Decode(
    'ChFGb2xsb3dDaGF0UmVxdWVzdBIXCgdjaGF0X2lkGAEgASgNUgZjaGF0SWQSHgoLbGFzdF9tc2'
    'dfaWQYAiABKA1SCWxhc3RNc2dJZA==');

@$core.Deprecated('Use messagesResponseDescriptor instead')
const MessagesResponse$json = {
  '1': 'MessagesResponse',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.proto.MessageData', '10': 'data'},
    {'1': 'follow', '3': 2, '4': 1, '5': 8, '10': 'follow'},
  ],
};

/// Descriptor for `MessagesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messagesResponseDescriptor = $convert.base64Decode(
    'ChBNZXNzYWdlc1Jlc3BvbnNlEiYKBGRhdGEYASADKAsyEi5wcm90by5NZXNzYWdlRGF0YVIEZG'
    'F0YRIWCgZmb2xsb3cYAiABKAhSBmZvbGxvdw==');

@$core.Deprecated('Use messageDataDescriptor instead')
const MessageData$json = {
  '1': 'MessageData',
  '2': [
    {'1': 'senderId', '3': 1, '4': 1, '5': 13, '10': 'senderId'},
    {'1': 'messageId', '3': 2, '4': 1, '5': 13, '10': 'messageId'},
    {'1': 'content', '3': 3, '4': 1, '5': 9, '10': 'content'},
    {'1': 'attachements', '3': 4, '4': 3, '5': 11, '6': '.proto.Attachment', '10': 'attachements'},
    {'1': 'repliedMessages', '3': 5, '4': 3, '5': 11, '6': '.proto.RepliedMessage', '10': 'repliedMessages'},
    {'1': 'sendAt', '3': 6, '4': 1, '5': 3, '10': 'sendAt'},
  ],
};

/// Descriptor for `MessageData`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageDataDescriptor = $convert.base64Decode(
    'CgtNZXNzYWdlRGF0YRIaCghzZW5kZXJJZBgBIAEoDVIIc2VuZGVySWQSHAoJbWVzc2FnZUlkGA'
    'IgASgNUgltZXNzYWdlSWQSGAoHY29udGVudBgDIAEoCVIHY29udGVudBI1CgxhdHRhY2hlbWVu'
    'dHMYBCADKAsyES5wcm90by5BdHRhY2htZW50UgxhdHRhY2hlbWVudHMSPwoPcmVwbGllZE1lc3'
    'NhZ2VzGAUgAygLMhUucHJvdG8uUmVwbGllZE1lc3NhZ2VSD3JlcGxpZWRNZXNzYWdlcxIWCgZz'
    'ZW5kQXQYBiABKANSBnNlbmRBdA==');

@$core.Deprecated('Use attachmentDescriptor instead')
const Attachment$json = {
  '1': 'Attachment',
  '2': [
    {'1': 'placeholder', '3': 1, '4': 1, '5': 9, '10': 'placeholder'},
    {'1': 'type', '3': 2, '4': 1, '5': 9, '10': 'type'},
    {'1': 'url', '3': 3, '4': 1, '5': 9, '10': 'url'},
  ],
};

/// Descriptor for `Attachment`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List attachmentDescriptor = $convert.base64Decode(
    'CgpBdHRhY2htZW50EiAKC3BsYWNlaG9sZGVyGAEgASgJUgtwbGFjZWhvbGRlchISCgR0eXBlGA'
    'IgASgJUgR0eXBlEhAKA3VybBgDIAEoCVIDdXJs');

@$core.Deprecated('Use repliedMessageDescriptor instead')
const RepliedMessage$json = {
  '1': 'RepliedMessage',
  '2': [
    {'1': 'messageId', '3': 1, '4': 1, '5': 13, '10': 'messageId'},
    {'1': 'content', '3': 2, '4': 1, '5': 9, '10': 'content'},
    {'1': 'thumbnails', '3': 3, '4': 3, '5': 9, '10': 'thumbnails'},
  ],
};

/// Descriptor for `RepliedMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List repliedMessageDescriptor = $convert.base64Decode(
    'Cg5SZXBsaWVkTWVzc2FnZRIcCgltZXNzYWdlSWQYASABKA1SCW1lc3NhZ2VJZBIYCgdjb250ZW'
    '50GAIgASgJUgdjb250ZW50Eh4KCnRodW1ibmFpbHMYAyADKAlSCnRodW1ibmFpbHM=');

@$core.Deprecated('Use getChatsRequestDescriptor instead')
const GetChatsRequest$json = {
  '1': 'GetChatsRequest',
  '2': [
    {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `GetChatsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getChatsRequestDescriptor = $convert.base64Decode(
    'Cg9HZXRDaGF0c1JlcXVlc3QSFAoFdG9rZW4YASABKAlSBXRva2Vu');

@$core.Deprecated('Use chatsResponseDescriptor instead')
const ChatsResponse$json = {
  '1': 'ChatsResponse',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.proto.ChatsResponse.ChatData', '10': 'data'},
  ],
  '3': [ChatsResponse_ChatData$json],
};

@$core.Deprecated('Use chatsResponseDescriptor instead')
const ChatsResponse_ChatData$json = {
  '1': 'ChatData',
  '2': [
    {'1': 'chat_id', '3': 1, '4': 1, '5': 13, '10': 'chatId'},
    {'1': 'chatTitle', '3': 2, '4': 1, '5': 9, '10': 'chatTitle'},
    {'1': 'type', '3': 3, '4': 1, '5': 9, '10': 'type'},
    {'1': 'mute', '3': 4, '4': 1, '5': 8, '10': 'mute'},
    {'1': 'avatarUrl', '3': 5, '4': 1, '5': 9, '10': 'avatarUrl'},
    {'1': 'lastMessage', '3': 6, '4': 1, '5': 9, '10': 'lastMessage'},
    {'1': 'unreadedMessagesCount', '3': 7, '4': 1, '5': 9, '10': 'unreadedMessagesCount'},
    {'1': 'lastReadedMessageId', '3': 8, '4': 1, '5': 13, '10': 'lastReadedMessageId'},
  ],
};

/// Descriptor for `ChatsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List chatsResponseDescriptor = $convert.base64Decode(
    'Cg1DaGF0c1Jlc3BvbnNlEjEKBGRhdGEYASADKAsyHS5wcm90by5DaGF0c1Jlc3BvbnNlLkNoYX'
    'REYXRhUgRkYXRhGpECCghDaGF0RGF0YRIXCgdjaGF0X2lkGAEgASgNUgZjaGF0SWQSHAoJY2hh'
    'dFRpdGxlGAIgASgJUgljaGF0VGl0bGUSEgoEdHlwZRgDIAEoCVIEdHlwZRISCgRtdXRlGAQgAS'
    'gIUgRtdXRlEhwKCWF2YXRhclVybBgFIAEoCVIJYXZhdGFyVXJsEiAKC2xhc3RNZXNzYWdlGAYg'
    'ASgJUgtsYXN0TWVzc2FnZRI0ChV1bnJlYWRlZE1lc3NhZ2VzQ291bnQYByABKAlSFXVucmVhZG'
    'VkTWVzc2FnZXNDb3VudBIwChNsYXN0UmVhZGVkTWVzc2FnZUlkGAggASgNUhNsYXN0UmVhZGVk'
    'TWVzc2FnZUlk');

@$core.Deprecated('Use createGroupRequestDescriptor instead')
const CreateGroupRequest$json = {
  '1': 'CreateGroupRequest',
  '2': [
    {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
    {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `CreateGroupRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createGroupRequestDescriptor = $convert.base64Decode(
    'ChJDcmVhdGVHcm91cFJlcXVlc3QSFAoFdG9rZW4YASABKAlSBXRva2VuEhIKBG5hbWUYAiABKA'
    'lSBG5hbWU=');

@$core.Deprecated('Use createGroupResponseDescriptor instead')
const CreateGroupResponse$json = {
  '1': 'CreateGroupResponse',
  '2': [
    {'1': 'groupId', '3': 1, '4': 1, '5': 13, '10': 'groupId'},
    {'1': 'chat_id', '3': 2, '4': 1, '5': 13, '10': 'chatId'},
  ],
};

/// Descriptor for `CreateGroupResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createGroupResponseDescriptor = $convert.base64Decode(
    'ChNDcmVhdGVHcm91cFJlc3BvbnNlEhgKB2dyb3VwSWQYASABKA1SB2dyb3VwSWQSFwoHY2hhdF'
    '9pZBgCIAEoDVIGY2hhdElk');

@$core.Deprecated('Use joinGroupRequestDescriptor instead')
const JoinGroupRequest$json = {
  '1': 'JoinGroupRequest',
  '2': [
    {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
    {'1': 'chat_id', '3': 2, '4': 1, '5': 13, '10': 'chatId'},
    {'1': 'role', '3': 3, '4': 1, '5': 9, '10': 'role'},
  ],
};

/// Descriptor for `JoinGroupRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List joinGroupRequestDescriptor = $convert.base64Decode(
    'ChBKb2luR3JvdXBSZXF1ZXN0EhQKBXRva2VuGAEgASgJUgV0b2tlbhIXCgdjaGF0X2lkGAIgAS'
    'gNUgZjaGF0SWQSEgoEcm9sZRgDIAEoCVIEcm9sZQ==');

@$core.Deprecated('Use joinGroupResponseDescriptor instead')
const JoinGroupResponse$json = {
  '1': 'JoinGroupResponse',
};

/// Descriptor for `JoinGroupResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List joinGroupResponseDescriptor = $convert.base64Decode(
    'ChFKb2luR3JvdXBSZXNwb25zZQ==');

@$core.Deprecated('Use addMemberRequestDescriptor instead')
const AddMemberRequest$json = {
  '1': 'AddMemberRequest',
  '2': [
    {'1': 'chat_id', '3': 1, '4': 1, '5': 13, '10': 'chatId'},
    {'1': 'membersId', '3': 2, '4': 3, '5': 13, '10': 'membersId'},
  ],
};

/// Descriptor for `AddMemberRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addMemberRequestDescriptor = $convert.base64Decode(
    'ChBBZGRNZW1iZXJSZXF1ZXN0EhcKB2NoYXRfaWQYASABKA1SBmNoYXRJZBIcCgltZW1iZXJzSW'
    'QYAiADKA1SCW1lbWJlcnNJZA==');

@$core.Deprecated('Use addMemberResponseDescriptor instead')
const AddMemberResponse$json = {
  '1': 'AddMemberResponse',
  '2': [
    {'1': 'status', '3': 1, '4': 1, '5': 9, '10': 'status'},
  ],
};

/// Descriptor for `AddMemberResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addMemberResponseDescriptor = $convert.base64Decode(
    'ChFBZGRNZW1iZXJSZXNwb25zZRIWCgZzdGF0dXMYASABKAlSBnN0YXR1cw==');

@$core.Deprecated('Use memberInfoRequestDescriptor instead')
const MemberInfoRequest$json = {
  '1': 'MemberInfoRequest',
  '2': [
    {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
    {'1': 'chat_id', '3': 2, '4': 1, '5': 13, '10': 'chatId'},
    {'1': 'mute', '3': 3, '4': 1, '5': 8, '9': 0, '10': 'mute'},
    {'1': 'lastReadedMessageId', '3': 4, '4': 1, '5': 13, '9': 0, '10': 'lastReadedMessageId'},
  ],
  '8': [
    {'1': 'info'},
  ],
};

/// Descriptor for `MemberInfoRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List memberInfoRequestDescriptor = $convert.base64Decode(
    'ChFNZW1iZXJJbmZvUmVxdWVzdBIUCgV0b2tlbhgBIAEoCVIFdG9rZW4SFwoHY2hhdF9pZBgCIA'
    'EoDVIGY2hhdElkEhQKBG11dGUYAyABKAhIAFIEbXV0ZRIyChNsYXN0UmVhZGVkTWVzc2FnZUlk'
    'GAQgASgNSABSE2xhc3RSZWFkZWRNZXNzYWdlSWRCBgoEaW5mbw==');

@$core.Deprecated('Use memberInfoResponseDescriptor instead')
const MemberInfoResponse$json = {
  '1': 'MemberInfoResponse',
};

/// Descriptor for `MemberInfoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List memberInfoResponseDescriptor = $convert.base64Decode(
    'ChJNZW1iZXJJbmZvUmVzcG9uc2U=');

@$core.Deprecated('Use selectChatRequestDescriptor instead')
const SelectChatRequest$json = {
  '1': 'SelectChatRequest',
  '2': [
    {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
    {'1': 'chat_id', '3': 2, '4': 1, '5': 13, '10': 'chatId'},
  ],
};

/// Descriptor for `SelectChatRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List selectChatRequestDescriptor = $convert.base64Decode(
    'ChFTZWxlY3RDaGF0UmVxdWVzdBIUCgV0b2tlbhgBIAEoCVIFdG9rZW4SFwoHY2hhdF9pZBgCIA'
    'EoDVIGY2hhdElk');

