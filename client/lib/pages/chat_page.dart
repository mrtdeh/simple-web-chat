import 'package:dashboard/grpc/grpc.dart';
import 'package:dashboard/pages/messages.dart';
import 'package:dashboard/pages/sidebar.dart';
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
    if (_scrollController.position.pixels ==
        _scrollController.position.maxScrollExtent) {
      wc.getMessages(RecordDirection.next, 50, context, onComplete: () {
        var t = wc.messages.length - 50;
        print("t : $t");
        _listController.jumpToItem(
          index: t - 1,
          scrollController: _scrollController,
          alignment: 1,
        );

        // WidgetsBinding.instance.addPostFrameCallback((_) {
        //   double targetAlignment = 1 - 0.5;
        //   _listController.animateToItem(
        //     index: t,
        //     scrollController: _scrollController,
        //     alignment: targetAlignment,
        //     duration: (estimatedDistance) => Duration(milliseconds: 1000),
        //     curve: (estimatedDistance) => Curves.easeInOut,
        //   );
        // });
      });
    }

    if (_scrollController.position.pixels ==
        _scrollController.position.minScrollExtent) {
      wc.getMessages(RecordDirection.previous, 50, context, onComplete: () {
        _listController.jumpToItem(
          index: 50,
          scrollController: _scrollController,
          alignment: 0,
        );

        // WidgetsBinding.instance.addPostFrameCallback((_) {
        //   double targetAlignment = 1 + 0.5;
        //   _listController.animateToItem(
        //     index: 49,
        //     scrollController: _scrollController,
        //     alignment: targetAlignment,
        //     duration: (estimatedDistance) => Duration(milliseconds: 1000),
        //     curve: (estimatedDistance) => Curves.easeInOut,
        //   );
        // });
      });
    }
  }

  final _scrollController = ScrollController();
  final _listController = ListController();
  TextEditingController _textController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Row(
        children: [
          ChatSidebar(
            title: "Chats View",
            width: CHATS_WIDTH,
            onChange: () => switchToWaiting(),
          ),
          Expanded(
            child: Column(
              children: [
                Expanded(
                  child: ChatMessages(
                    listController: _listController,
                    scrollController: _scrollController,
                  ),
                ),
                SizedBox(
                  height: 50,
                  child: Row(
                    children: [
                      Expanded(
                        child: TextFormField(
                          maxLines: 5,
                          controller: _textController,
                        ),
                      ),
                      ElevatedButton(
                        onPressed: () async {
                          // Create new message object
                          var msg = wc.newMessage(
                              content: _textController.text,
                              senderId: wc.userID);
                          // Add new message bubble to messages list with sending status
                          wc.addMessage(msg);
                          // Send new message to server
                          var chatId = wc.getActiveChatID(); 
                          print("chatId : $chatId");

                          wc.sendMessage(
                            chatId: chatId,
                            message: msg,
                          );
                          setState(() {
                            
                          });
                        },
                        child: Text("Send"),
                      )
                    ],
                  ),
                )
              ],
            ),
          ),
        ],
      ),
    );
  }
}
