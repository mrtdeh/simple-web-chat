import 'dart:async';
import 'package:flutter/material.dart';
import 'package:grpc/grpc_web.dart';

import '../proto/service.pbgrpc.dart';

class Message {
  final MessagesResponse_MessageData data;
  final GlobalKey key;
  double? height;

  void setHeight(double h) {
    this.height = h;
  }

  Message({
    required this.data,
    required this.key,
    this.height,
  });
}

class WebChat {
  WebChat._privateConstructor();
  static final WebChat instance = WebChat._privateConstructor();

  late ChatServiceClient _service;
  String token = "";
  List<ChatsResponse_ChatData> chats = [];
  List<Message> messages = [];

  final StreamController<List<ChatsResponse_ChatData>> _chatController = StreamController<List<ChatsResponse_ChatData>>.broadcast();
  Stream<List<ChatsResponse_ChatData>> get chatStream => _chatController.stream;

  final StreamController<List<Message>> _messageController = StreamController<List<Message>>.broadcast();
  Stream<List<Message>> get messageStream => _messageController.stream;

  Future<void> start() async {
    try {
      final channel = GrpcWebClientChannel.xhr(Uri.parse('http://localhost:8081'));
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

  final double pageSize = 150;

  void getMessages(int chatId, int readedMsgId, int nextCount, int prevCount, Function(double) onComplete) async {
    final request = GetMessagesRequest(chatId: chatId, nextCount: nextCount, prevCount: prevCount, readedMsgId: readedMsgId);
    double totalHeight = 0;
    _service.getMessages(request).listen((response) {
      List<Message> msgs = [];
      response.data.forEach(
        (msg) {
          msgs.add(Message(data: msg, key: new GlobalKey()));
        },
      );

      totalHeight = 0;
      print("debug 1 : " + totalHeight.toString());

      if (nextCount > 0 && prevCount == 0) {
        messages.insertAll(messages.length, msgs);
        if (messages.length > pageSize) {
          var count = messages.length - pageSize;

          for (var i = 0; i <= count - 1; i++) {
            if (messages[i].height != null) {
              totalHeight += messages[i].height!;
            }
          }
          print("totalHeight : " + totalHeight.toString());

          messages.removeRange(0, count.toInt());
        }
      } else if (prevCount > 0 && nextCount == 0) {
        messages.insertAll(0, msgs);
        if (messages.length > pageSize) {
          // var count = messages.length ;

          // for (var i = messages.length - pageSize; i < count - 1; i++) {
          //   if (messages[i].height != null) {
          //     totalHeight += messages[i].height!;
          //   }
          // }
          // print("totalHeight : " + totalHeight.toString());

          messages.removeRange((messages.length - pageSize).toInt(), messages.length);
        }
      } else {
        messages.insertAll(0, msgs);
      }

      _messageController.sink.add(messages);
    }, onError: (error) {
      print("Error in get messages: $error");
    }, onDone: () {
      print('get messages closed');
      onComplete(totalHeight);
    });
  }

  void startMessageChannel() {
    print("start listening...");
    final request = MessageChannelRequest()..token = token;
    _service.messageChannel(request).listen((response) {
      if (response.hasChats()) {
        for (var chat in response.chats.data) {
          chats.add(chat);
        }
        _chatController.sink.add(chats);
      }
      // if (response.hasMessages()) {
      //   for (var msg in response.messages.data) {
      //     messages.add(msg);
      //   }
      //   _messageController.sink.add(messages);
      // }
    }, onError: (error) {
      print("Error in message channel: $error");
    }, onDone: () {
      print('Closed connection to server.');
      _retryStream();
    });
  }

  void _retryStream() {
    print("Retrying the stream...");
    Future.delayed(Duration(seconds: 5), () {
      startMessageChannel();
    });
  }
}

final $WebChat = WebChat.instance;
