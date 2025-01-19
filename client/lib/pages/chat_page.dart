import 'dart:convert';

import 'package:dashboard/grpc/grpc.dart';
import 'package:dashboard/proto/service.pb.dart';
import 'package:flutter/material.dart';
import 'package:super_sliver_list/super_sliver_list.dart';

class ChatScreen extends StatefulWidget {
  const ChatScreen({super.key});

  @override
  State<ChatScreen> createState() => _ChatScreenState();
}

class _ChatScreenState extends State<ChatScreen> {
  double wh = 0;
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

  void _onScroll() {
    // print("${_scrollController.position.pixels} : ${_scrollController.position.maxScrollExtent}");
    if (_scrollController.position.pixels == _scrollController.position.maxScrollExtent) {
      wc.getMessages(RecordDirection.next, 50, context, onComplete: () {
        print("okkkkkkkkkkkk");
        if (_listController.numberOfItems > 100) {
          _listController.jumpToItem(
            // duration: (estimatedDistance) => Duration(milliseconds: 1000),
            // curve: (estimatedDistance) => Curves.easeInOut,
            index: 99,
            scrollController: _scrollController,
            alignment: 1 - wh,
          );
        }
      });
    }
    if (_scrollController.position.pixels == _scrollController.position.minScrollExtent) {
      wc.getMessages(RecordDirection.previous, 50, context, onComplete: () {
        // _scrollController.jumpTo(_scrollController.position.pixels + totalHeight);
        _listController.jumpToItem(index: 50, scrollController: _scrollController, alignment: 0 + (40 / MediaQuery.of(context).size.height));
      });
    }
  }

  final _scrollController = ScrollController();
  final _listController = ListController();

  @override
  Widget build(BuildContext context) {
    wh = 40 / (MediaQuery.of(context).size.height);
    print("wh : ${wh}");

    return Scaffold(
      // appBar: AppBar(
      //   title: Text("Chat App"),
      //   scrolledUnderElevation: 0,

      // ),
      body: Row(
        children: [
          Container(
            width: CHATS_WIDTH,
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
                          final c = chats[index];
                          return ListTile(
                            trailing: Badge(
                              backgroundColor: Colors.red,
                              textColor: Colors.white,
                              label: Text(c.unreadedMessagesCount),
                            ),
                            leading: CircleAvatar(
                              child: Icon(Icons.person),
                            ),
                            title: Text(c.chatTitle),
                            subtitle: Text(c.lastMessage, style: TextStyle(color: Colors.white)),
                            onTap: () async {
                              switchToWaiting();
                              wc.setChat(index);
                              wc.getMessages(RecordDirection.none, 50, context);
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
              color: const Color.fromARGB(255, 19, 19, 19),
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
                      child: SuperListView.builder(
                        // cacheExtent: double.infinity,

                        shrinkWrap: true,
                        controller: _scrollController,
                        listController: _listController,
                        itemCount: messages.length,
                        itemBuilder: (context, index) {
                          final msg = messages[index];
                          return Center(
                            child: Container(
                              key: msg.key,
                              width: 600,
                              margin: EdgeInsets.symmetric(horizontal: 10),
                              child: Directionality(
                                textDirection: msg.toLeft! ? TextDirection.ltr : TextDirection.rtl,
                                child: Row(
                                  crossAxisAlignment: CrossAxisAlignment.end,
                                  children: [
                                    msg.haveAvatar!
                                        ? Container(
                                            margin: EdgeInsets.symmetric(horizontal: 10),
                                            child: CircleAvatar(
                                              backgroundColor: Colors.amber,
                                              radius: 25,
                                            ),
                                          )
                                        : SizedBox(),
                                    Flexible(
                                      child: Column(
                                        crossAxisAlignment: CrossAxisAlignment.start,
                                        children: [
                                          index == 0
                                              ? SizedBox(
                                                  height: 40,
                                                )
                                              : SizedBox(),
                                          msg.data.attachements.isNotEmpty
                                              ? SizedBox(
                                                  width: 300,
                                                  child: GridView.builder(
                                                    gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(crossAxisCount: 3, crossAxisSpacing: 5, mainAxisSpacing: 5),
                                                    controller: new ScrollController(keepScrollOffset: false),
                                                    shrinkWrap: true,
                                                    scrollDirection: Axis.vertical,
                                                    itemCount: msg.data.attachements.length,
                                                    itemBuilder: (context, index) {
                                                      return SizedBox(
                                                        child: Image.memory(
                                                          base64Decode(msg.data.attachements[index].placeholder),
                                                        ),
                                                      );
                                                    },
                                                  ),
                                                )
                                              : SizedBox(),
                                          Container(
                                            margin: const EdgeInsets.symmetric(vertical: 10),
                                            padding: const EdgeInsets.all(10),
                                            decoration: BoxDecoration(
                                              color: const Color.fromARGB(255, 43, 43, 43),
                                              borderRadius: BorderRadius.only(
                                                bottomLeft: msg.toLeft! ? Radius.zero : Radius.circular(10),
                                                bottomRight: msg.toLeft! ? Radius.circular(10) : Radius.zero,
                                                topLeft: Radius.circular(10),
                                                topRight: Radius.circular(10),
                                              ),
                                            ),
                                            height: msg.textHeight,
                                            child: RichText(
                                              text: TextSpan(
                                                text: msg.data.content,
                                                style: defaultTextStyle,
                                              ),
                                            ),
                                          ),
                                          index == messages.length - 1
                                              ? SizedBox(
                                                  height: 40,
                                                )
                                              : SizedBox(),
                                        ],
                                      ),
                                    ),
                                  ],
                                ),
                              ),
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
