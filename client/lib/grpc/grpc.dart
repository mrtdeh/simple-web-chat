import 'dart:async';
import 'package:flutter/material.dart';
import 'package:grpc/grpc_web.dart';

import '../proto/service.pbgrpc.dart';

TextStyle defaultTextStyle = const TextStyle(color: Colors.white, fontSize: 16, height: 2);

class Message {
  final MessagesResponse_MessageData data;
  final GlobalKey key;
  double? textHeight;
  double? boxHeight;
  bool? haveAvatar;
  bool? toLeft;

  Message({
    required this.data,
    required this.key,
    this.textHeight,
    this.boxHeight,
    this.haveAvatar,
    this.toLeft,
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
  int userID = 0;
  // ChatsResponse_ChatData? chat;
  int _selectedChatIndex = 0;
  void setChat(int index) {
    _selectedChatIndex = index;
  }

  // ChatsResponse_ChatData getChat(int index){

  // }

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
      userID = response.userId;
      print("Login successful, token: $token");
      return true;
    } catch (error) {
      print("Login failed: $error");
      return false;
    }
  }

  final double pageSize = 150;

  void getMessages(RecordDirection direction, int count, BuildContext context, {Function(double)? onComplete}) async {
    var readedMsgId = 0;

    if (direction == RecordDirection.next) {
      // Set parameters for get next page
      readedMsgId = messages[messages.length - 1].data.messageId;
    } else if (direction == RecordDirection.previous) {
      // Set parameters for get previous page
      readedMsgId = messages[0].data.messageId;
    } else {
      readedMsgId = 150;
    }

    var chatId = chats[_selectedChatIndex].chatId;

    final request = GetMessagesRequest(chatId: chatId, direction: direction.toString(), count: count, readedMsgId: readedMsgId);
    double totalHeight = 0;
    List<Message> msgs = [];

    _service.getMessages(request).listen((response) {
      for (var msg in response.data) {
        // Create new Message class with calculate the box height
        var newMsg = newBoxMessage(msg, context);
        // Add message to temprary list
        msgs.add(newMsg);
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

  Message newBoxMessage(MessagesResponse_MessageData msg, BuildContext context) {
    var text = msg.content;
    var attLen = msg.attachements.length;
    double textHeight = 0.0;

    TextPainter textPainter = TextPainter(
      text: TextSpan(text: text, style: defaultTextStyle),
      textDirection: TextDirection.ltr,
      maxLines: null,
    );

    // double maxWidth = MediaQuery.of(context).size.width - (250 + 40);
    double avatarWidth = 45; // with margins
    double viewHorizontalPadding = 40;
    double chatsWidth = 250;
    double others = avatarWidth + viewHorizontalPadding + chatsWidth;
    double maxWidth = MediaQuery.of(context).size.width - others;
    if (maxWidth > 600) {
      maxWidth = 600;
    }

    textPainter.layout(maxWidth: maxWidth);

    var textVerticalPadding = 20;
    textHeight += textPainter.size.height + textVerticalPadding;

    var textVerticalMargin = 20;
    var atc = (attLen / 3).ceil();
    var attachHeight = 100;
    var boxHeight = textHeight + textVerticalMargin + (atc * attachHeight);

    var chat = chats[_selectedChatIndex];
    return Message(
      data: msg,
      key: GlobalKey(),
      textHeight: textHeight,
      boxHeight: boxHeight,
      haveAvatar: false,
      toLeft: msg.senderId != userID,
    );
  }

  void _retryStream() {
    print("Retrying the stream...");
    Future.delayed(Duration(seconds: 5), () {
      startMessageChannel();
    });
  }
}

final wc = WebChat.instance;
