import 'package:dashboard/grpc/grpc.dart';
import 'package:dashboard/proto/service.pb.dart';
import 'package:flutter/material.dart';

class ChatSidebar extends StatelessWidget {
  final VoidCallback? onChange;
  final String? title;
  final double? width;
  const ChatSidebar({super.key, this.title, this.width, this.onChange});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: width,
      color: Colors.black54,
      child: Column(
        children: [
          Container(
            padding: EdgeInsets.all(16),
            alignment: Alignment.centerLeft,
            child: Text(
              title ?? "Chats",
              style: TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.bold,
              ),
            ),
          ),
          Expanded(
            child: StreamBuilder<List<ChatsResponse_ChatData>>(
              stream: wc.chatStream,
              builder: (context, snapshot) {
                if (snapshot.connectionState == ConnectionState.waiting) {
                  return Center(child: CircularProgressIndicator());
                }
                if (snapshot.hasError) {
                  return Center(child: Text("Error: \${snapshot.error}"));
                }
                if (!snapshot.hasData || snapshot.data!.isEmpty) {
                  return Center(child: Text("No chats available"));
                }

                final chats = snapshot.data!;
                return ListView.builder(
                  itemCount: chats.length,
                  itemBuilder: (context, index) {
                    final c = chats[index];
                    return ChatListItem(
                      chat: c,
                      index: index,
                      onTap: onChange,
                    );
                  },
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}

class ChatListItem extends StatelessWidget {
  final ChatsResponse_ChatData chat;
  final int index;
  final VoidCallback? onTap;

  const ChatListItem({super.key, required this.chat, required this.index, this.onTap});

  @override
  Widget build(BuildContext context) {
    return ListTile(
      trailing: Badge(
        backgroundColor: Colors.red,
        textColor: Colors.white,
        label: Text(chat.unreadedMessagesCount),
      ),
      leading: CircleAvatar(
        child: Icon(Icons.person),
      ),
      title: Text(chat.chatTitle),
      subtitle: Text(chat.lastMessage, style: TextStyle(color: Colors.white)),
      onTap: () async {
        if (onTap != null) {
          onTap!();
        }
        wc.setChat(index);
        wc.getMessages(GetMessagesRequest_Direction.None, 50, context);
      },
    );
  }
}
