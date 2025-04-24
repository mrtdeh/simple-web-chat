import 'dart:async';
// import 'package:dashboard/utils/locker.dart';
import 'package:dashboard/utils/locker.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc_web.dart';
import 'package:synchronized/synchronized.dart';

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

const dirNext = GetMessagesRequest_Direction.NextPage;
const dirPrev = GetMessagesRequest_Direction.PrevPage;
const dirBoth = GetMessagesRequest_Direction.BothPage;
const dirLast = GetMessagesRequest_Direction.LastPage;
const dirNone = GetMessagesRequest_Direction.None;

class WebChat {
  WebChat._privateConstructor();
  static final WebChat instance = WebChat._privateConstructor();

  late ChatServiceClient _service;
  int _selectedChatIndex = 0;
  double jumpHeight = 0;
  int userID = 0;
  String token = "";
  List<Message> messages = [];
  List<ChatData> chats = [];
  // ChatData? chat;
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

  final StreamController<List<ChatData>> _chatStream = StreamController<List<ChatData>>.broadcast();
  Stream<List<ChatData>> get chatStream => _chatStream.stream;

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

  final double pageMax = 150;
  var msglock = new Lock();
  final locker = Locker<String>();

  void getMessages(
    GetMessagesRequest_Direction direction,
    int count,
    BuildContext context, {
    Function(int)? onComplete,
  }) async {
    if (locker.isLocked("getMessage")) {
      print("Task is locked!");
      return;
    }
    locker.lock("getMessage");

    var fromMsgId = 0;

    switch (direction) {
      case dirNext || dirLast:
        // Set parameters for get next page
        fromMsgId = messages.last.data.messageId;
        if (fromMsgId == 0) {
          return;
        }
      // break;
      case dirPrev:
        // Set parameters for get previous page
        fromMsgId = messages.first.data.messageId;
      // break;
      default:
    }

    print("get message : $direction $count from $fromMsgId");

    var chatId = chats[_selectedChatIndex].chatId;
    int? lastMsgId;
    if (messages.length > 0) {
      lastMsgId = messages.last.data.messageId;
    }
    var currentPageSize = messages.length;

    final request = GetMessagesRequest(
      chatId: chatId,
      direction: direction,
      count: count,
      fromMsgId: fromMsgId,
      lastMsgId: lastMsgId,
      pageSize: currentPageSize,
      pageMax: pageMax.toInt(),
      token: token,
    );

    List<Message> msgs = [];

    _service.getMessages(request).listen((response) {
      for (var msg in response.data) {
        // Create new Message class with calculate the box height
        var newMsg = newBoxMessage(msg);
        // Add message to temprary list
        msgs.add(newMsg);
      }
    }, onDone: () {
      // Process recieved messages and put to listview
      processReceivedMessages(msgs, direction);
      if (onComplete != null) {
        onComplete(msgs.length);
      }
      locker.unlock("getMessage");
    }, onError: (error) {
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
    final request = StreamRequest()..token = token;
    _service.streamChannel(request).listen((response) {
      if (response.hasChats()) {
        for (var chat in response.chats.data) {
          chats.add(chat);
        }
        _chatStream.sink.add(chats);
      }

      if (response.hasMessages()) {
        List<Message> msgs = [];
        for (var msg in response.messages.data) {
          // Create new Message class with calculate the box height
          var newMsg = newBoxMessage(msg);
          // Add message to temprary list
          msgs.add(newMsg);
        }

        processReceivedMessages(msgs, dirNext);
      }
    }, onError: (error) {
      print("Error in message channel: $error");
    }, onDone: () {
      print('Closed connection to server.');
      _retryStream();
    });
  }

  void processReceivedMessages(
    List<Message> msgs,
    GetMessagesRequest_Direction direction, {
    Function(int)? onComplete,
  }) {
    msglock.synchronized(() async {
      try {
        if (msgs.isEmpty) {
          // call done callback if defined
          if (onComplete != null) {
            onComplete(0);
          }
          // Return on empty messages received
          return;
        }
        switch (direction) {
          case dirNext:
            // If do lazy-loading next message
            // Add messages to end of list
            messages.insertAll(messages.length, msgs);
            // Check if total messages reached maximum page size
            if (messages.length > pageMax) {
              // Calculate total height of messages from top of list
              var count = messages.length - pageMax;
              // Removing messages from the top of list
              messages.removeRange(0, count.toInt());
            }
          // break;
          case dirPrev:
            // If do lazy-loading previous message
            // Add messages to the top of list
            messages.insertAll(0, msgs);
            // Check if total messages reached maximum page size
            if (messages.length > pageMax) {
              // Removing the messages from the down of list
              var len = messages.length;
              var rmCounts = (messages.length - pageMax).toInt();
              messages.removeRange(len - rmCounts, len);
            }
          // break;
          case dirLast:
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
      } catch (e) {
        print("error in process messages :" + e.toString());
      }
    });
  }

  Message newBoxMessage(MessageData msg) {
    // var text = msg.content;
    // var attLen = msg.attachements.length;
    // double textHeight = 0.0;

    // TextPainter textPainter = TextPainter(
    //   text: TextSpan(text: text, style: defaultTextStyle),
    //   textDirection: TextDirection.ltr,
    //   maxLines: null,
    // );

    // double windowWidth = MediaQuery.of(context).size.width;
    // double maxWidth = windowWidth;
    // double textBoxPadding = 20; // text box (left + right) padding
    // double viewInnerWidth = 600; // without margins
    // double viewOutterWidth = 620; // with margins
    // double avatarWidth = 70; // (radius * 2) + left/right margins
    // double leftPanelWidth = CHATS_WIDTH;
    // // double others = avatarWidth + viewHorizontalPadding + chatsWidth;
    // double messagesViewWidth = windowWidth - leftPanelWidth;
    // if (messagesViewWidth > viewOutterWidth) {
    //   maxWidth = viewInnerWidth - (avatarWidth + textBoxPadding);
    // }

    // textPainter.layout(maxWidth: maxWidth);

    // var textVerticalPadding = 20;
    // textHeight += textPainter.size.height + textVerticalPadding;

    // var textVerticalMargin = 20;
    // var atc = (attLen / 3).ceil();
    // var attachHeight = 100;
    // var boxHeight = textHeight + textVerticalMargin + (atc * attachHeight);

    var chat = chats[_selectedChatIndex];
    return Message(
      data: msg,
      key: GlobalKey(),
      // textHeight: textHeight,
      // boxHeight: boxHeight,
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
