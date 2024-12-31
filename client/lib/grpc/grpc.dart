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
      for (var msg in response.data) {
        // Calculate the box height
        var height = _calculateHeight(msg.content, context);
        // Calculate the box height with margin
        var boxHeight = height + 20 + 100;
        msgs.add(Message(
          data: msg,
          key: GlobalKey(),
          height: height,
          boxHeight: boxHeight,
        ));
      }
    }, onDone: () {
      if (direction == RecordDirection.next) {
        // If do lazy-loading next message
        // Add messages to end of list
        messages.insertAll(messages.length, msgs);
        // Check if total messages reached maximum page size
        if (messages.length > pageSize) {
          // Calculate total height of messages from top of list
          var count = messages.length - pageSize;
          for (var i = 0; i <= count - 1; i++) {
            totalHeight += messages[i].boxHeight!;
          }
          // Removing messages from the top of list
          messages.removeRange(0, count.toInt());
        }
      } else if (direction == RecordDirection.previous) {
        // If do lazy-loading previous message
        // Add messages to the top of list
        messages.insertAll(0, msgs);
        // Check if total messages reached maximum page size
        if (messages.length > pageSize) {
          // Removing the messages from the down of list
          var len = messages.length;
          var rmCounts = (messages.length - pageSize).toInt();
          messages.removeRange(len - rmCounts, len);
        }
        // Calculate the total height of messages to be added to the top of the list.
        for (var msg in msgs) {
          totalHeight += msg.boxHeight!;
        }
      } else {
        // If do not lazy-loading messages
        messages.insertAll(0, msgs);
      }
      // Update stream sink to update Listview of messages
      _messageController.sink.add(messages);
      // reset incomming list
      msgs = [];
      // call done callback if defined
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

    // double maxWidth = MediaQuery.of(context).size.width - (250 + 40);
    double avatarWidth = 45; // with margins
    double messagesViewPadding = 40;
    double chatsWidth = 250;
    double others = avatarWidth + messagesViewPadding + chatsWidth;
    double maxWidth = MediaQuery.of(context).size.width - others;
    if (maxWidth > 600) {
      maxWidth = 600;
    }

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
