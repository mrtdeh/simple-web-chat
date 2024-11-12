import 'package:dashboard/grpc/grpc.dart';
import 'package:flutter/material.dart';
import '../proto/service.pbgrpc.dart';
import 'package:grpc/grpc_web.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Web Chat',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: ChatScreen(),
    );
  }
}

class ChatScreen extends StatefulWidget {
  @override
  _ChatScreenState createState() => _ChatScreenState();
}

class _ChatScreenState extends State<ChatScreen> {
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  final List<String> _messages = [];
  // late WebChat _webChat;
  late ChatServiceClient _service;
  String? token;
  bool _isLoggedIn = false;

  @override
  void initState() {
    super.initState();
    _initializeGrpcClient();
  }

  void _initializeGrpcClient() {
    final channel =
        GrpcWebClientChannel.xhr(Uri.parse('http://localhost:8080'));

    _service = ChatServiceClient(channel);
  }

  Future<void> _login() async {
    final username = _usernameController.text;
    final password = _passwordController.text;

    try {
      var res = await _service
          .login(LoginRequest(password: password, username: username));
      if (res.hasToken()) {
        setState(() {
          token = res.token;
          _isLoggedIn = true;
        });
        _startMessageChannel();
      } else {
        _showError("Login failed. Please check your credentials.");
      }
    } catch (error) {
      _showError("Login error: $error");
    }
  }

  void _startMessageChannel() {
    _service.messageChannel(MessageChannelRequest(token: token)).listen(
        (message) {
      if (message.hasMessages()) {
        for (var msg in message.messages) {}
      }
    }, onError: (error) {
      _showError("Message channel error: $error");
    });
  }

  void _showError(String error) {
    ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(error)));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Web Chat'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            if (!_isLoggedIn) ...[
              TextField(
                controller: _usernameController,
                decoration: InputDecoration(labelText: 'Username'),
              ),
              TextField(
                controller: _passwordController,
                decoration: InputDecoration(labelText: 'Password'),
                obscureText: true,
              ),
              SizedBox(height: 16),
              ElevatedButton(
                onPressed: _login,
                child: Text('Login'),
              ),
            ] else ...[
              Expanded(
                child: ListView.builder(
                  itemCount: _messages.length,
                  itemBuilder: (context, index) {
                    return ListTile(
                      title: Text(_messages[index]),
                    );
                  },
                ),
              ),
              SizedBox(height: 16),
              Text("You are logged in and receiving messages."),
            ],
          ],
        ),
      ),
    );
  }

  @override
  void dispose() {
    _usernameController.dispose();
    _passwordController.dispose();
    super.dispose();
  }
}
