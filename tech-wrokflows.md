
# Technical Workflows

## Frontend: Update Last Read Messages from Chats

1. Open a chat from the sidebar and fetch all messages by chat ID.  
2. In the messages view, each message box that appears in the viewport triggers a handler which sends the message ID of viewed messages to the server.  
3. The server validates and stores the received message IDs as read in the database.  
4. The frontend updates local chat data to reflect the correct read status and metrics in the chats view.