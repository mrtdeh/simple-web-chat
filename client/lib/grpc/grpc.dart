import 'dart:async';
import 'dart:html';

import '../proto/service.pbgrpc.dart';

class WebChat {
  final ChatServiceClient _service;

  WebChat(this._service);

  Future<void> sendMessage(String token, message) async {
    _addLeftMessage(message);

    try {
      // final response = await _service.echo(EchoRequest()..message = message);
      final response = await _service.sendMessage(
          MessageRequest(chatId: 1, content: message, token: token));
      _addRightMessage(response.message);
    } catch (error) {
      _addRightMessage(error.toString());
    }
  }

  void createMessageChannel(String token) async {
    final request = MessageChannelRequest()..token = token;

    _service.messageChannel(request).listen((response) {
      _addRightMessage(response.message);
    }, onError: (error) {
      _addRightMessage(error.toString());
    }, onDone: () => print('Closed connection to server.'));
  }


  void _addLeftMessage(String message) {
    _addMessage(message, 'label-primary pull-left');
  }

  void _addRightMessage(String message) {
    _addMessage(message, 'label-default pull-right');
  }

  void _addMessage(String message, String cssClass) {
    final classes = cssClass.split(' ');
    querySelector('#first')!.after(DivElement()
      ..classes.add('row')
      ..append(Element.tag('h2')
        ..append(SpanElement()
          ..classes.add('label')
          ..classes.addAll(classes)
          ..text = message)));
  }
}
