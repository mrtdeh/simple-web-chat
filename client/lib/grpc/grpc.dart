import 'dart:async';
// import 'dart:ffi';
// import 'package:dashboard/utils/locker.dart';
import 'package:dashboard/grpc/chat.dart';
import 'package:dashboard/grpc/message.dart';
import 'package:dashboard/utils/locker.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc_web.dart';
import 'package:synchronized/synchronized.dart';

import '../proto/service.pbgrpc.dart';

TextStyle defaultTextStyle =
    const TextStyle(color: Colors.white, fontSize: 16, height: 2);

const double CHATS_WIDTH = 400;

enum MessageStatus { sending, sended, received, failed }

const dirNext = GetMessagesRequest_Direction.NextPage;
const dirPrev = GetMessagesRequest_Direction.PrevPage;
const dirBoth = GetMessagesRequest_Direction.BothPage;
const dirLast = GetMessagesRequest_Direction.LastPage;
const dirNone = GetMessagesRequest_Direction.None;

class WebChat {
  WebChat._privateConstructor();
  static final WebChat instance = WebChat._privateConstructor();

  late ChatServiceClient _service;
  int _selectedChatIndex = -1;
  double jumpHeight = 0;
  int userID = 0;
  String token = "";
  List<Message> messages = [];
  List<Chat> chats = [];
  
  final double pageMax = 150;
  var msglock = Lock();
  final locker = Locker<String>();

  final ValueNotifier<int> unreadedMessagesCount = ValueNotifier<int>(0);


  final StreamController<List<Chat>> _chatStream =
      StreamController<List<Chat>>.broadcast();
  Stream<List<Chat>> get chatStream => _chatStream.stream;

  StreamController<List<Message>> _messageStream =
      StreamController<List<Message>>.broadcast();
  Stream<List<Message>> get messageStream => _messageStream.stream;



  // int getUnreadedMessagesCount() {
  //   var i = 0;
  //   for (var m in messages) {
  //     if (m.data.messageId >
  //         chats[_selectedChatIndex].data.lastReadedMessageId) {
  //       i++;
  //     }
  //   }
  //   return i;
  // }


  int getActiveChatID() {
    return chats[_selectedChatIndex].data.chatId;
  }

  Chat getActiveChat() {
    return chats[_selectedChatIndex];
  }

  int getActiveChatIndex() {
    return _selectedChatIndex;
  }

  void setChat(int index) {
    _selectedChatIndex = index;
  }

  void setLastReadedMessageID(int id) {
    print("set last readed id :$id");
    chats[_selectedChatIndex].data.lastReadedMessageId = id;

    // Calculate Unreaded messages count
    var i = 0;
    for (var m in messages) {
      if (m.data.messageId >
          chats[_selectedChatIndex].data.lastReadedMessageId) {
        i++;
      }
    }
    chats[_selectedChatIndex].unreadedMessagesCount = i;
    unreadedMessagesCount.value = i;

    if (locker.isLocked("setLastReadedMessageID")) {
      return;
    }
    locker.lock("setLastReadedMessageID");

    Future.delayed(Duration(seconds: 1), () {
      // wc.updateLastReadedMessageID();
      var readedID = chats[_selectedChatIndex].data.lastReadedMessageId;
      _service.updateLastReadedMessageID(LrmRequest(
          token: token, chatId: getActiveChatID(), readedMsgId: readedID));
      locker.unlock("setLastReadedMessageID");
    });
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
        haveAvatar: chat.data.type == "public",
        toLeft: false);

    return msg;
  }

  void init() {
    Future.delayed(Duration(microseconds: 1), () {
      _messageStream.sink.add([]);
    });
  }

  void resetStream() {
    messages = [];
    _messageStream = StreamController<List<Message>>.broadcast();
  }

  Future<void> connectToServer() async {
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
      userID = response.userId;
      print("Login successful, token: $token");
      return true;
    } catch (error) {
      print("Login failed: $error");
      return false;
    }
  }

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

    var chatId = chats[_selectedChatIndex].data.chatId;
    int? lastMsgId;
    if (messages.isNotEmpty) {
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

  void sendMessage(
      {Message? message, int? chatId, Function()? onComplete}) async {
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

  bool updateMessages(List<Message> messages, MessageData recv) {
    for (var curr in messages) {
      if (curr.data.messageId == recv.messageId) {
        // Check updated date
        // if (recv.updatedAt < curr.data.updatedAt) {
        //   return false;
        // }
        // Successfully update chat data
        curr.data = recv;
        return true;
      }
    }
    // Append a new chat
    messages.add(newBoxMessage(recv));
    return true;
  }

  bool updateChats(List<Chat> mychats, ChatData recv) {
    for (var curr in mychats) {
      if (curr.data.chatId == recv.chatId) {
        // Check last reader id
        if (recv.lastReadedMessageId < curr.data.lastReadedMessageId) {
          return false;
        }
        // Successfully update chat data
        curr.data = recv;
        return true;
      }
    }
    // Append a new chat
    mychats.add(Chat(data: recv));
    return true;
  }

  void startMessageChannel() {
    print("start listening...");
    final request = StreamRequest()..token = token;
    _service.streamChannel(request).listen((response) {
      if (response.hasChats()) {
        for (var chat in response.chats.data) {
          // Check and add/update new chat to chats list
          updateChats(chats, chat);
        }
        _chatStream.sink.add(chats);
      }

      if (response.hasMessages()) {
        List<Message> msgs = [];
        for (var msg in response.messages.data) {
          // Check and add/update new message to messages list
          updateMessages(msgs, msg);
        }
        processReceivedMessages(msgs, dirNext);
      }
    }, onError: (error) {
      print("Error in message channel: $error");
    }, onDone: () {
      print('Closed connection to server.');
      print("Retrying the stream...");
      Future.delayed(Duration(seconds: 5), () {
        startMessageChannel();
      });
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
    var chat = chats[_selectedChatIndex];
    return Message(
      data: msg,
      key: GlobalKey(),
      haveAvatar: chat.data.type == "public",
      toLeft: msg.senderId != userID,
    );
  }


}

final wc = WebChat.instance;
