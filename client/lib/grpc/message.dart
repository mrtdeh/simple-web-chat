import 'package:dashboard/grpc/grpc.dart';
import 'package:dashboard/proto/service.pb.dart';
import 'package:flutter/material.dart';

class Message {
  MessageData data;
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
