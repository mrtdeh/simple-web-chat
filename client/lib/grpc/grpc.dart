import 'dart:async';
import 'dart:html';
import '../proto/service.pbgrpc.dart';

class WebChat {
  final ChatServiceClient _service;
  String? token;

  WebChat(this._service);

  Future<void> login(String username, String password) async {
    try {
      final response = await _service.login(
        LoginRequest(username: username, password: password),
      );
      token = response.token;
      print("Login successful, token: $token");
    } catch (error) {
      _addRightMessage("Login failed: $error");
      print("Login failed: $error");
    }
  }

  void createMessageChannel(String token) {
    final request = MessageChannelRequest()..token = token;
    _service.messageChannel(request).listen((response) {
      // _addRightMessage(response.message);
    }, onError: (error) {
      _addRightMessage("Error: $error");
      print("Error in message channel: $error");
    }, onDone: () {
      print('Closed connection to server.');
    });
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
