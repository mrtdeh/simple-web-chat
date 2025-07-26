import 'package:dashboard/proto/service.pb.dart';
import 'package:flutter/material.dart';

class Chat {
  late final int id;
  late final String title;
  late final String type;
  late final bool mute;
  late final String avatarUrl;
  late final String lastMessage;
  late final ValueNotifier<int> lastReadedMessageId;
  late final ValueNotifier<int> unreadedMessagesCount;

  Chat(ChatData data) {
    id = data.chatId;
    title = data.chatTitle;
    type = data.type;
    mute = data.mute;
    avatarUrl = data.avatarUrl;
    lastMessage = data.lastMessage;
    lastReadedMessageId = ValueNotifier<int>(data.lastReadedMessageId);
    unreadedMessagesCount = ValueNotifier<int>(data.unreadedMessagesCount);
  }

  bool setLastReadedMessageID(int id, int readedCount) {
    if (id > lastReadedMessageId.value) {
      lastReadedMessageId.value = id;
      unreadedMessagesCount.value -= readedCount;
      return true;
    }
    return false;
  }
}
