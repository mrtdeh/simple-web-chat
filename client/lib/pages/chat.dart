import 'package:dashboard/grpc/grpc.dart';
import 'package:dashboard/proto/service.pb.dart';
import 'package:flutter/material.dart';

class ChatScreen extends StatelessWidget {
  const ChatScreen({Key? key}) : super(key: key);

  // یک Stream برای دریافت لیست چت‌ها به صورت بلادرنگ
  Stream<List<String>> get chatStream async* {
    // شبیه‌سازی داده‌های چت به صورت بلادرنگ
    await Future.delayed(Duration(seconds: 1));
    yield ["Chat 1", "Chat 2", "Chat 3"];
    await Future.delayed(Duration(seconds: 3));
    yield ["Chat 1", "Chat 2", "Chat 3", "Chat 4"];
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Chat App"),
      ),
      body: Row(
        children: [
          // بخش سایدبار برای نمایش لیست چت‌ها
          Container(
            width: 250, // عرض سایدبار
            color: Colors.grey[200],
            child: Column(
              children: [
                Container(
                  padding: EdgeInsets.all(16),
                  alignment: Alignment.centerLeft,
                  child: Text(
                    "Chats",
                    style: TextStyle(
                      fontSize: 24,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),
                Expanded(
                  child: StreamBuilder<List<ChatsResponse_ChatData>>(
                    stream: $WebChat.chatStream,
                    builder: (context, snapshot) {
                      if (snapshot.connectionState == ConnectionState.waiting) {
                        return Center(child: CircularProgressIndicator());
                      }
                      if (snapshot.hasError) {
                        return Center(child: Text("Error: ${snapshot.error}"));
                      }
                      if (!snapshot.hasData || snapshot.data!.isEmpty) {
                        return Center(child: Text("No chats available"));
                      }

                      // نمایش لیست چت‌ها
                      final chats = snapshot.data!;
                      return ListView.builder(
                        itemCount: chats.length,
                        itemBuilder: (context, index) {
                          return ListTile(
                            leading: CircleAvatar(
                              child: Icon(Icons.person),
                            ),
                            title: Text(chats[index].chatTitle),
                            subtitle: Text("Last message preview..."),
                            onTap: () {
                              // عمل کلیک بر روی چت
                            },
                          );
                        },
                      );
                    },
                  ),
                ),
              ],
            ),
          ),
          // بخش اصلی چت
          Expanded(
            child: Container(
              color: Colors.white,
              child: Center(
                child: Text(
                  "Select a chat to start messaging",
                  style: TextStyle(fontSize: 18, color: Colors.grey),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
