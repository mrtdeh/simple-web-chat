import 'package:dashboard/grpc/grpc.dart';
import 'package:dashboard/pages/chat.dart';
import 'package:flutter/material.dart';

class LoginPage extends StatelessWidget {
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();

  LoginPage({super.key});

  Future<bool> _handleLogin() async {
    final username = _usernameController.text;
    final password = _passwordController.text;

    // Call the login function from the singleton instance
    return await $WebChat.login(username, password);
  }

  void _showError(BuildContext context, String error) {
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text(error)),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Login")),
      body: FutureBuilder<void>(
        future: $WebChat.start(),
        builder: (context, snapshot) {
          // حالت در حال بارگذاری
          if (snapshot.connectionState == ConnectionState.waiting) {
            return Center(child: CircularProgressIndicator());
          }

          // حالت خطا
          if (snapshot.hasError) {
            return Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text("Error: ${snapshot.error}"),
                  ElevatedButton(
                    onPressed: () {
                      $WebChat.start();
                    },
                    child: Text("Retry"),
                  ),
                ],
              ),
            );
          }

          // اگر همه چیز درست باشد، فرم لاگین را نمایش بده
          return _buildLoginForm(context);
        },
      ),
    );
  }

  Widget _buildLoginForm(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          TextField(
            decoration: InputDecoration(labelText: "Username"),
            controller: _usernameController,
          ),
          TextField(
            decoration: InputDecoration(labelText: "Password"),
            obscureText: true,
            controller: _passwordController,
          ),
          SizedBox(height: 20),
          ElevatedButton(
            onPressed: () async {
              if (await _handleLogin()) {
                print("login success");
                Navigator.pushReplacement(
                    context,
                    MaterialPageRoute(
                      builder: (context) => ChatScreen(),
                    ));
              } else {
                _showError(context, "username or password incorrect");
              }
            },
            child: Text("Login"),
          ),
        ],
      ),
    );
  }
}
