import 'dart:async';
import 'dart:ffi';
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

enum RecordDirection {
  previous,
  next,
  none
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

  StreamController<List<Message>> _messageController = StreamController<List<Message>>.broadcast();
  Stream<List<Message>> get messageStream => _messageController.stream;

  void init() {
    Future.delayed(Duration(microseconds: 1), () {
    _messageController.sink.add([]);
    });
  }

//   Stream<List<Message>> getMessageStream() async* {
//   yield null; 
//   await Future.delayed(Duration(seconds: 2)); 
//   yield await messageStream.;
// }

  void resetStream() {
    _messageController = StreamController<List<Message>>.broadcast();
  }

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

  void getMessages(int chatId, RecordDirection direction, int count, BuildContext context, {Function(double)? onComplete}) async {
    var readedMsgId = 0;
    var nextCount = 0;
    var prevCount = 0;

    if (direction == RecordDirection.next) {
      // Set parameters for get next page
      nextCount = count;
      readedMsgId = messages[messages.length - 1].data.messageId;
    } else if (direction == RecordDirection.previous) {
      // Set parameters for get previous page
      prevCount = count;
      readedMsgId = messages[0].data.messageId;
    } else {
      // Set parameters for get
      nextCount = count;
      prevCount = count;
    }

    final request = GetMessagesRequest(chatId: chatId, nextCount: nextCount, prevCount: prevCount, readedMsgId: readedMsgId);
    double totalHeight = 0;
    List<Message> msgs = [];
    _service.getMessages(request).listen((response) {
          for(var msg in response.data){
        var height = _calculateHeight(msg.content, context);
        var boxHeight = height + 20;
        msgs.add(Message(
          data: msg,
          key:  GlobalKey(),
          height: height,
          boxHeight: boxHeight,
        ));
      }

    }, onDone: () async {
      // totalHeight = 0;
      if (nextCount > 0 && prevCount == 0) {
        messages.insertAll(messages.length, msgs);
        if (messages.length > pageSize) {
          var count = messages.length - pageSize;
          for (var i = 0; i <= count - 1; i++) {
            totalHeight += messages[i].boxHeight!;
          }
          messages.removeRange(0, count.toInt());
        }
      } else if (prevCount > 0 && nextCount == 0) {
        messages.insertAll(0, msgs);
        if (messages.length > pageSize) {
          var len = messages.length;
          var rmCounts = (messages.length - pageSize).toInt();
          messages.removeRange(len - rmCounts, len);
        }
        for(var msg in msgs){
            totalHeight += msg.boxHeight!;
        }
      } else {
        messages.insertAll(0, msgs);
      }

      _messageController.sink.add(messages);
      msgs = [];
      if (onComplete != null) {
        onComplete(totalHeight);
      }
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

final wc = WebChat.instance;
