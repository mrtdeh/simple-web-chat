import 'dart:convert';

import 'package:dashboard/grpc/grpc.dart';
import 'package:extended_image/extended_image.dart';
import 'package:flutter/material.dart';
import 'package:super_sliver_list/super_sliver_list.dart';

class ChatMessages extends StatelessWidget {
  final ScrollController? scrollController;
  final ListController? listController;

  const ChatMessages({super.key, this.listController, this.scrollController});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: const Color.fromARGB(255, 19, 19, 19),
      child: RawScrollbar(
        controller: scrollController,
        thumbColor: Colors.redAccent,
        radius: Radius.circular(20),
        thickness: 15,
        child: StreamBuilder<List<Message>?>(
          stream: wc.messageStream,
          initialData: const [],
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.waiting) {
              return Center(child: CircularProgressIndicator());
            }
            if (snapshot.hasError) {
              return Center(child: Text('Error: ${snapshot.error}'));
            }
            if (!snapshot.hasData || snapshot.data!.isEmpty) {
              return Center(
                child: Text(
                  "Select a chat to start messaging",
                  style: TextStyle(fontSize: 18, color: Colors.grey),
                ),
              );
            }
            final messages = snapshot.data!;
            return MessagesList(
              messages: messages,
              listController: listController,
              scrollController: scrollController,
            );
          },
        ),
      ),
    );
  }
}

class MessagesList extends StatelessWidget {
  final List<Message> messages;
  final ScrollController? scrollController;
  final ListController? listController;

  const MessagesList({super.key, required this.messages, this.listController, this.scrollController});

  @override
  Widget build(BuildContext context) {
    return ScrollConfiguration(
      behavior: ScrollConfiguration.of(context).copyWith(scrollbars: false),
      child: SuperListView.builder(
        shrinkWrap: true,
        controller: scrollController,
        listController: listController,
        itemCount: messages.length,
        itemBuilder: (context, index) {
          return MessageBubble(message: messages[index], index: index);
        },
      ),
    );
  }
}

class MessageBubble extends StatelessWidget {
  final Message message;
  final int index;

  const MessageBubble({super.key, required this.message, required this.index});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Container(
        // key: message.key,
        width: 600,
        margin: EdgeInsets.symmetric(horizontal: 10),
        child: Directionality(
          textDirection: message.toLeft! ? TextDirection.ltr : TextDirection.rtl,
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.end,
            children: [
              message.haveAvatar!
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
                    message.data.attachements.isNotEmpty
                        ? SizedBox(
                            width: 300,
                            child: GridView.builder(
                              gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                                crossAxisCount: 3,
                                crossAxisSpacing: 5,
                                mainAxisSpacing: 5,
                              ),
                              controller: ScrollController(keepScrollOffset: false),
                              shrinkWrap: true,
                              itemCount: message.data.attachements.length,
                              itemBuilder: (context, index) {
                                // return ExtendedImage.memory(
                                //   base64Decode(message.data.attachements[index].placeholder),
                                //        gaplessPlayback: true,
                                // );
                                return Image(
                                  image: MemoryImage(base64Decode(message.data.attachements[index].placeholder)),
                                  gaplessPlayback: true,
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
                          bottomLeft: message.toLeft! ? Radius.zero : Radius.circular(10),
                          bottomRight: message.toLeft! ? Radius.circular(10) : Radius.zero,
                          topLeft: Radius.circular(10),
                          topRight: Radius.circular(10),
                        ),
                      ),
                      child: RichText(
                        text: TextSpan(
                          text: "${message.data.content} - ${message.status}",
                          style: defaultTextStyle,
                        ),
                      ),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
