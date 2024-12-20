import 'package:dashboard/grpc/grpc.dart';
import 'package:dashboard/proto/service.pb.dart';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';

class ChatScreen extends StatefulWidget {
  const ChatScreen({super.key});

  @override
  _ChatScreenState createState() => _ChatScreenState();
}

class _ChatScreenState extends State<ChatScreen> {
  @override
  void initState() {
    $WebChat.startMessageChannel();
    _scrollController.addListener(_onScroll);
    super.initState();
  }

  int lastMsgId = 0;
  int firstMsgId = 0;
  int chatId = 0;
  // double currentOffset = 0;

  void _onScroll() {
    if (_scrollController.position.pixels == _scrollController.position.maxScrollExtent) {
      $WebChat.getMessages(chatId, lastMsgId, 50, 0, context, (totalHeight) {
        // WidgetsBinding.instance.addPostFrameCallback((_) {
        _scrollController.jumpTo(_scrollController.position.pixels - totalHeight);
        // });
      });
    }
    if (_scrollController.position.pixels == _scrollController.position.minScrollExtent) {
      $WebChat.getMessages(chatId, firstMsgId, 0, 50, context, (totalHeight) {
        // WidgetsBinding.instance.addPostFrameCallback((_) {
        _scrollController.jumpTo(_scrollController.position.pixels + totalHeight);
        // });
      });
    }
  }

  final _scrollController = ScrollController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Chat App"),
        scrolledUnderElevation: 0,
      ),
      body: Row(
        children: [
          // بخش سایدبار برای نمایش لیست چت‌ها
          Container(
            width: 250, // عرض سایدبار
            color: Colors.black54,
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
                            trailing: Badge(
                              backgroundColor: Colors.red,
                              textColor: Colors.white,
                              label: Text(chats[index].unreadedMessagesCount),
                            ),
                            leading: CircleAvatar(
                              child: Icon(Icons.person),
                            ),
                            title: Text(chats[index].chatTitle),
                            subtitle: Text(chats[index].lastMessage, style: TextStyle(color: Colors.white)),
                            onTap: () {
                              chatId = chats[index].chatId;
                              $WebChat.getMessages(chats[index].chatId, chats[index].lastReadedMessageId, 50, 50, context, (t) {});
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
              child: RawScrollbar(
                controller: _scrollController,
                thumbColor: Colors.redAccent,
                radius: Radius.circular(20),
                thickness: 15,
                child: StreamBuilder<List<Message>>(
                  stream: $WebChat.messageStream,
                  builder: (context, snapshot) {
                    if (snapshot.connectionState == ConnectionState.waiting) {
                      return Center(child: CircularProgressIndicator());
                    }
                    if (snapshot.hasError) {
                      return Center(child: Text("Error: ${snapshot.error}"));
                    }
                    if (!snapshot.hasData || snapshot.data!.isEmpty) {
                      return Center(
                          child: Text(
                        "Select a chat to start messaging",
                        style: TextStyle(fontSize: 18, color: Colors.grey),
                      ));
                    }

                    // نمایش لیست چت‌ها
                    final messages = snapshot.data!;
                    // return Text("");
                    return ListView.builder(
                      cacheExtent: double.maxFinite,
                      controller: _scrollController,
                      scrollDirection: Axis.vertical,
                      itemCount: messages.length,
                      itemBuilder: (context, index) {
                        if (index == 0) {
                          firstMsgId = messages[0].data.messageId;
                        }
                        if (index == messages.length - 1) {
                          lastMsgId = messages[messages.length - 1].data.messageId;
                        }
                        final msg = messages[index];
                        return Container(
                          key: msg.key,
                          margin: const EdgeInsets.symmetric(vertical: 10),
                          padding: const EdgeInsets.all(10),
                          decoration: BoxDecoration(
                            color: Colors.blueAccent,
                            borderRadius: BorderRadius.circular(10),
                          ),
                          width: double.infinity,
                          height: msg.height,
                          child: RichText(
                            text: TextSpan(
                              text: msg.data.content,
                              style: defaultTextStyle,
                            ),
                          ),
                        );
                      },
                      // children: List.generate(messages.length, (index) {

                      // }),
                    );
                  },
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
