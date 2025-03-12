import 'dart:async';
import 'package:dashboard/utils/locker.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc_web.dart';

import '../proto/service.pbgrpc.dart';

TextStyle defaultTextStyle = const TextStyle(color: Colors.white, fontSize: 16, height: 2);

class Message {
  final MessageData data;
  final GlobalKey key;
  double? textHeight;
  double? boxHeight;
  bool? haveAvatar;
  bool? toLeft;
  MessageStatus? status;

  Message({required this.data, required this.key, this.textHeight, this.boxHeight, this.haveAvatar, this.toLeft, this.status});

  void setStatus(MessageStatus s) {
    status = s;
  }

  void setID(int id) {
    data.messageId = id;
  }
}

const double CHATS_WIDTH = 400;

enum MessageStatus {
  sending,
  sended,
  received,
  failed
}

enum RecordDirection {
  previous,
  next,
  last,
  none
}

class WebChat {
  WebChat._privateConstructor();
  static final WebChat instance = WebChat._privateConstructor();

  late ChatServiceClient _service;
  int _selectedChatIndex = 0;
  double jumpHeight = 0;
  int userID = 0;
  String token = "";
  List<Message> messages = [];
  List<ChatsResponse_ChatData> chats = [];
  // ChatsResponse_ChatData? chat;
  int getActiveChatID() {
    return chats[_selectedChatIndex].chatId;
  }

  void setChat(int index) {
    _selectedChatIndex = index;
  }

  void addMessage(Message msg) {
    messages.add(msg);
    _messageStream.add(messages);
  }

  void updateMessage(Message msg) {
    final updatedMessages = messages.map((m) {
      return m.data.messageId == msg.data.messageId ? msg : m;
    }).toList();
    messages = updatedMessages;
    _messageStream.sink.add(messages);
  }

  Message newMessage({String? content, int? senderId}) {
    var chat = chats[_selectedChatIndex];
    var msg = Message(
        data: MessageData(
          content: content,
          senderId: senderId,
        ),
        key: GlobalKey(),
        haveAvatar: chat.type == "public",
        toLeft: false);

    return msg;
  }

  final StreamController<List<ChatsResponse_ChatData>> _chatStream = StreamController<List<ChatsResponse_ChatData>>.broadcast();
  Stream<List<ChatsResponse_ChatData>> get chatStream => _chatStream.stream;

  StreamController<List<Message>> _messageStream = StreamController<List<Message>>.broadcast();
  Stream<List<Message>> get messageStream => _messageStream.stream;

  void init() {
    Future.delayed(Duration(microseconds: 1), () {
      _messageStream.sink.add([]);
    });
  }

  void resetStream() {
    messages = [];
    _messageStream = StreamController<List<Message>>.broadcast();
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
  final locker = Locker<String>();

  void getMessages(
    RecordDirection direction,
    int count,
    BuildContext context, {
    Function(int)? onComplete,
  }) async {
    if (locker.isLocked("getMessage")) {
      print("Task is locked!");
      return;
    }

    var readedMsgId = 0;

    switch (direction) {
      case RecordDirection.next || RecordDirection.last:
        // Set parameters for get next page
        readedMsgId = messages[messages.length - 1].data.messageId;
        if (readedMsgId == 0) {
          return;
        }
      // break;
      case RecordDirection.previous:
        // Set parameters for get previous page
        readedMsgId = messages[0].data.messageId;
      // break;
      default:
    }

    print("readedMsgId : $readedMsgId");

    // _lock = true;
    locker.lock("getMessage");

    print("get message : $direction $count from $readedMsgId");

    var chatId = chats[_selectedChatIndex].chatId;

    final request = GetMessagesRequest(
      chatId: chatId,
      direction: direction.index,
      count: count,
      readedMsgId: readedMsgId,
    );
    // double totalHeight = 0;
    List<Message> msgs = [];

    _service.getMessages(request).listen((response) {
      for (var msg in response.data) {
        // Create new Message class with calculate the box height
        var newMsg = newBoxMessage(msg, context);
        // Add message to temprary list
        msgs.add(newMsg);
      }
    }, onDone: () {
      try {
        if (msgs.isEmpty) {
          // call done callback if defined
          if (onComplete != null) {
            onComplete(msgs.length);
          }
          // Return on empty messages received
          return;
        }
        switch (direction) {
          case RecordDirection.next:
            // If do lazy-loading next message
            // Add messages to end of list
            messages.insertAll(messages.length, msgs);
            // Check if total messages reached maximum page size
            if (messages.length > pageSize) {
              // Calculate total height of messages from top of list
              var count = messages.length - pageSize;
              // for (var i = 0; i <= count - 1; i++) {
              //   totalHeight += messages[i].boxHeight!;
              // }
              // Removing messages from the top of list
              messages.removeRange(0, count.toInt());
            }
          // break;
          case RecordDirection.previous:
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
          // break;
          case RecordDirection.last:
            messages = msgs;
          // break;
          default:
            // If do not lazy-loading messages
            messages.insertAll(0, msgs);
        }
        // Update stream sink to update Listview of messages
        _messageStream.sink.add(messages);

        // call done callback if defined
        if (onComplete != null) {
          onComplete(msgs.length);
        }
      } finally {
        // reset incomming list
        msgs = [];
        // _lock = false;
        locker.unlock("getMessage");
      }
    }, onError: (error) {
      // _lock = false;
      locker.unlock("getMessage");
      print("Error in get messages: $error");
    });
  }

  void sendMessage({Message? message, int? chatId, Function()? onComplete}) async {
    final request = MessageRequest(
      chatId: chatId,
      content: message!.data.content,
      token: token,
    );

    try {
      var res = await _service.sendMessage(request);
      message.setStatus(MessageStatus.sended);
      message.setID(res.messageId);
      if (onComplete != null) {
        onComplete();
      }
    } catch (e) {
      message.setStatus(MessageStatus.failed);
      print(e);
    }
    //  res.messageId
  }

  void startMessageChannel() {
    print("start listening...");
    final request = NotificationRequest()..token = token;
    _service.notificationChannel(request).listen((response) {
      if (response.hasChats()) {
        for (var chat in response.chats.data) {
          chats.add(chat);
        }
        _chatStream.sink.add(chats);
      }
    }, onError: (error) {
      print("Error in message channel: $error");
    }, onDone: () {
      print('Closed connection to server.');
      _retryStream();
    });
  }

  Message newBoxMessage(MessageData msg, BuildContext context) {
    var text = msg.content;
    var attLen = msg.attachements.length;
    double textHeight = 0.0;

    TextPainter textPainter = TextPainter(
      text: TextSpan(text: text, style: defaultTextStyle),
      textDirection: TextDirection.ltr,
      maxLines: null,
    );

    double windowWidth = MediaQuery.of(context).size.width;
    double maxWidth = windowWidth;
    double textBoxPadding = 20; // text box (left + right) padding
    double viewInnerWidth = 600; // without margins
    double viewOutterWidth = 620; // with margins
    double avatarWidth = 70; // (radius * 2) + left/right margins
    double leftPanelWidth = CHATS_WIDTH;
    // double others = avatarWidth + viewHorizontalPadding + chatsWidth;
    double messagesViewWidth = windowWidth - leftPanelWidth;
    if (messagesViewWidth > viewOutterWidth) {
      maxWidth = viewInnerWidth - (avatarWidth + textBoxPadding);
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
      haveAvatar: chat.type == "public",
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
