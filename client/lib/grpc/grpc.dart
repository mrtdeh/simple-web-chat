import 'dart:async';
import 'package:grpc/grpc_web.dart';

import '../proto/service.pbgrpc.dart';

class WebChat {
  WebChat._privateConstructor();
  static final WebChat instance = WebChat._privateConstructor();

  late ChatServiceClient _service;
  String? token;

  // WebChat(this._service);

  final StreamController<List<ChatsResponse_ChatData>> _chatController =
      StreamController<List<ChatsResponse_ChatData>>.broadcast();
  Stream<List<ChatsResponse_ChatData>> get chatStream => _chatController.stream;

  Future<void> start() async {
    try {
      final channel =
          GrpcWebClientChannel.xhr(Uri.parse('http://localhost:8081'));
      _service = ChatServiceClient(channel);
    } catch (err) {
      print("connect to server failed : " + err.toString());
    }
  }

  Future<bool> login(String username, String password) async {
    try {
      final response = await _service.login(
        LoginRequest(username: username, password: password),
      );
      token = response.token;
      print("Login successful, token: $token");
      return true;
    } catch (error) {
      print("Login failed: $error");
      return false;
    }
  }

  void startMessageChannel(String token) {
    final request = MessageChannelRequest()..token = token;

    _service.messageChannel(request).listen((response) {
      if (response.hasChats()) {
        for (var chat in response.chats.data) {
          // put json to stream
          _chatController.sink.add([chat]);

          print("Add to stream : " + chat.chatTitle);
        }
      }
    }, onError: (error) {
      print("Error in message channel: $error");
    }, onDone: () {
      print('Closed connection to server.');
    });
  }
}

final $WebChat = WebChat.instance;
