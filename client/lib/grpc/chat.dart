// import 'package:dashboard/grpc/grpc.dart';
import 'package:dashboard/proto/service.pb.dart';
// import 'package:flutter/material.dart';

class Chat {
  final ChatData data;
  int unreadedMessagesCount;

  Chat({required this.data, this.unreadedMessagesCount = 0});
}
