import 'package:dashboard/grpc/grpc.dart';
import 'package:dashboard/pages/messages.dart';
import 'package:dashboard/pages/sidebar.dart';
import 'package:flutter/material.dart';
import 'package:super_sliver_list/super_sliver_list.dart';
import '../proto/service.pbgrpc.dart';

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

  void scrollToBottom() {
    // if (withAnimate) {
    //   _listController.animateToItem(
    //     index: wc.messages.length - 2,
    //     scrollController: _scrollController,
    //     alignment: 0,
    //     duration: (estimatedDistance) => Duration(milliseconds: 400),
    //     curve: (estimatedDistance) => Curves.easeInOut,
    //   );
    // } else {
    _listController.jumpToItem(
      // index: wc.messages.length - 1,
      index: _listController.numberOfItems - 1,
      scrollController: _scrollController,
      alignment: 0,
    );
    // }
  }

  void _onScroll() {
    if (_scrollController.position.pixels == _scrollController.position.maxScrollExtent) {
      wc.getMessages(GetMessagesRequest_Direction.NextPage, 50, context, onComplete: (x) {
        if (wc.messages.length > 50 && x > 0) {
          var t = wc.messages.length - 50;

          _listController.jumpToItem(
            index: t - 1,
            scrollController: _scrollController,
            alignment: 1,
          );
        }

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

    if (_scrollController.position.pixels == _scrollController.position.minScrollExtent) {
      wc.getMessages(GetMessagesRequest_Direction.PrevPage, 50, context, onComplete: (x) {
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
                          try {
                            // Create new message object
                            var msg = wc.newMessage(content: _textController.text, senderId: wc.userID);
                            // Check page position
                            wc.getMessages(GetMessagesRequest_Direction.LastPage, 50, context, onComplete: (x) {
                              _textController.text = "";
                              // Add new message bubble to messages list with sending status
                              wc.addMessage(msg);
                              // Scroll to bottom of screen
                              scrollToBottom();
                              // Send new message to server
                              wc.sendMessage(
                                chatId: wc.getActiveChatID(),
                                message: msg,
                                onComplete: () {
                                  Future.delayed(Duration(seconds: 5), () {
                                    wc.updateMessage(msg);
                                    setState(() {});
                                  });
                                },
                              );
                            });
                          } catch (e) {
                            print("error : $e");
                          }
                          // setState(() {});
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
