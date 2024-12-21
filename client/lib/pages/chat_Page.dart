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
    Future.delayed(Duration(microseconds: 1), () {
      $WebChat.Init();
    });

    super.initState();
  }

  void switchToWaiting() {
    setState(() {
      $WebChat.ResetStream(); // استریم خالی برای حالت waiting
    });
  }

  int chatId = 0;

  void _onScroll() {
    if (_scrollController.position.pixels == _scrollController.position.maxScrollExtent) {
      $WebChat.getMessages(chatId, RecordDirection.next, 50, context, onComplete: (totalHeight) {
        _scrollController.jumpTo(_scrollController.position.pixels - totalHeight);
      });
    }
    if (_scrollController.position.pixels == _scrollController.position.minScrollExtent) {
      $WebChat.getMessages(chatId, RecordDirection.previous, 50, context, onComplete: (totalHeight) {
        _scrollController.jumpTo(_scrollController.position.pixels + totalHeight);
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
                            onTap: () async {
                              switchToWaiting();
                              chatId = chats[index].chatId;
                              Future.delayed(Duration(seconds: 1), () {
                                $WebChat.getMessages(chats[index].chatId, RecordDirection.none, 50, context);
                              });
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
                child: StreamBuilder<List<Message>?>(
                  stream: $WebChat.messageStream,
                  initialData: [],
                  builder: (context, snapshot) {
                    switch (snapshot.connectionState) {
                      case ConnectionState.waiting:
                        // While waiting for the data to load, show a loading spinner.
                        return Center(child: CircularProgressIndicator());
                      default:
                        if (snapshot.hasError)
                          return Center(child: Text('Error: ${snapshot.error}'));
                        else if (!snapshot.hasData || snapshot.data!.isEmpty)
                          return Center(
                              child: Text(
                            "Select a chat to start messaging",
                            style: TextStyle(fontSize: 18, color: Colors.grey),
                          ));
                    }

                    // if (!snapshot.hasData || snapshot.data!.isEmpty) {
                    // return Center(
                    //     child: Text(
                    //   "Select a chat to start messaging",
                    //   style: TextStyle(fontSize: 18, color: Colors.grey),
                    // ));
                    // } else if (snapshot.connectionState == ConnectionState.waiting) {
                    //   return Center(child: CircularProgressIndicator());
                    // } else if (snapshot.hasError) {
                    //   return Center(child: Text("Error: ${snapshot.error}"));
                    // }

                    final messages = snapshot.data!;
                    return ListView.builder(
                      controller: _scrollController,
                      // scrollDirection: Axis.vertical,
                      shrinkWrap: true,
                      itemCount: messages.length,
                      itemBuilder: (context, index) {
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
