import 'package:dashboard/grpc/grpc.dart';
import 'package:dashboard/proto/service.pb.dart';
import 'package:flutter/material.dart';

class ChatScreen extends StatefulWidget {
  const ChatScreen({super.key});

  @override
  State<ChatScreen> createState() => _ChatScreenState();
}

class _ChatScreenState extends State<ChatScreen> {
  @override
  void initState() {
    wc.startMessageChannel();
    _scrollController.addListener(_onScroll);
    wc.init();

    super.initState();
  }

  void switchToWaiting() {
    setState(() {
      wc.resetStream();
    });
  }

  int chatId = 0;

  void _onScroll() {
    if (_scrollController.position.pixels == _scrollController.position.maxScrollExtent) {
      wc.getMessages(chatId, RecordDirection.next, 50, context, onComplete: (totalHeight) {
        _scrollController.jumpTo(_scrollController.position.pixels - totalHeight);
      });
    }
    if (_scrollController.position.pixels == _scrollController.position.minScrollExtent) {
      wc.getMessages(chatId, RecordDirection.previous, 50, context, onComplete: (totalHeight) {
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
          Container(
            width: 250,
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
                    stream: wc.chatStream,
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
                                wc.getMessages(chats[index].chatId, RecordDirection.none, 50, context);
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
          Expanded(
            child: Container(
              color: Colors.white,
              child: RawScrollbar(
                controller: _scrollController,
                thumbColor: Colors.redAccent,
                radius: Radius.circular(20),
                thickness: 15,
                child: StreamBuilder<List<Message>?>(
                  stream: wc.messageStream,
                  initialData: const [],
                  builder: (context, snapshot) {
                    switch (snapshot.connectionState) {
                      case ConnectionState.waiting:
                        // While waiting for the data to load, show a loading spinner.
                        return Center(child: CircularProgressIndicator());
                      default:
                        if (snapshot.hasError) {
                          return Center(child: Text('Error: ${snapshot.error}'));
                        } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
                          return Center(
                              child: Text(
                            "Select a chat to start messaging",
                            style: TextStyle(fontSize: 18, color: Colors.grey),
                          ));
                        }
                    }

                    final messages = snapshot.data!;
                    return ScrollConfiguration(
                      behavior: ScrollConfiguration.of(context).copyWith(scrollbars: false),
                      child: ListView.builder(
                        controller: _scrollController,
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
                            child: Column(
                              children: [
                                msg.data.attachements.isNotEmpty
                                    ? SizedBox(
                                        height: 100,
                                        width: 100,
                                        child: Image.network(msg.data.attachements[0].url),
                                      )
                                    : SizedBox(),
                                RichText(
                                  text: TextSpan(
                                    text: msg.data.content,
                                    style: defaultTextStyle,
                                  ),
                                ),
                              ],
                            ),
                          );
                        },
                      ),
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
