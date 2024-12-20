import 'dart:async';
import 'package:flutter/material.dart';
import 'package:grpc/grpc_web.dart';

import '../proto/service.pbgrpc.dart';

TextStyle defaultTextStyle = const TextStyle(color: Colors.white, fontSize: 16, height: 2);

class Message {
  final MessagesResponse_MessageData data;
  final GlobalKey key;
  double? height;
  double? boxHeight;

  Message({
    required this.data,
    required this.key,
    this.height,
    this.boxHeight,
  });
}

class WebChat {
  WebChat._privateConstructor();
  static final WebChat instance = WebChat._privateConstructor();

  late ChatServiceClient _service;
  String token = "";

  List<ChatsResponse_ChatData> chats = [];
  List<Message> messages = [];
  double jumpHeight = 0;

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

  void getMessages(int chatId, int readedMsgId, int nextCount, int prevCount, BuildContext context, Function(double) onComplete) async {
    final request = GetMessagesRequest(chatId: chatId, nextCount: nextCount, prevCount: prevCount, readedMsgId: readedMsgId);
    double totalHeight = 0;
    List<Message> msgs = [];
    _service.getMessages(request).listen((response) {
      print("receive ${response.data.length} record");

      response.data.forEach((msg) {
        var height = _calculateHeight(msg.content, context);
        var boxHeight = height + 20;
        msgs.add(Message(
          data: msg,
          key: new GlobalKey(),
          height: height,
          boxHeight: boxHeight,
        ));
        // totalHeight += boxHeight;
      });
    }, onDone: () async {
      // totalHeight = 0;
      if (nextCount > 0 && prevCount == 0) {
        messages.insertAll(messages.length, msgs);
        if (messages.length > pageSize) {
          var count = messages.length - pageSize;
          for (var i = 0; i <= count - 1; i++) {
            // var ctx = messages[i].key.currentContext!;
            // final RenderBox box = ctx.findRenderObject() as RenderBox;

            totalHeight += messages[i].boxHeight!;
          }
          // print("remove from 0 to $count");
          messages.removeRange(0, count.toInt());
        }
      } else if (prevCount > 0 && nextCount == 0) {
        messages.insertAll(0, msgs);
        if (messages.length > pageSize) {
          var len = messages.length;
          var rmCounts = (messages.length - pageSize).toInt();
          messages.removeRange(len - rmCounts, len);
        }
        msgs.forEach(
          (msg) {
            totalHeight += msg.boxHeight!;
          },
        );
      } else {
        messages.insertAll(0, msgs);
      }

      _messageController.sink.add(messages);
      msgs = [];
      onComplete(totalHeight);
      print('get messages closed');
    }, onError: (error) {
      print("Error in get messages: $error");
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
    }, onError: (error) {
      print("Error in message channel: $error");
    }, onDone: () {
      print('Closed connection to server.');
      _retryStream();
    });
  }

  double _calculateHeight(String text, BuildContext context) {
    double totalHeight = 0.0;

    TextPainter textPainter = TextPainter(
      text: TextSpan(text: text, style: defaultTextStyle),
      textDirection: TextDirection.ltr,
      maxLines: null,
    );

    double maxWidth = MediaQuery.of(context).size.width - (250 + 40);
    textPainter.layout(maxWidth: maxWidth);

    totalHeight += textPainter.size.height;

    return totalHeight + 20;
  }

  void _retryStream() {
    print("Retrying the stream...");
    Future.delayed(Duration(seconds: 5), () {
      startMessageChannel();
    });
  }
}

final $WebChat = WebChat.instance;
