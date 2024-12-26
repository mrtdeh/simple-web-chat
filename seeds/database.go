package seeds

import (
	"api-channel/pkg/models"
	"fmt"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {
	// Clear records of table
	truncate_table(db, models.Attachment{}, true)
	truncate_table(db, models.Reply{}, true)
	truncate_table(db, models.Message{}, true)
	truncate_table(db, models.ChatMember{}, false)
	truncate_table(db, models.User{}, true)
	truncate_table(db, models.Group{}, true)
	truncate_table(db, models.Chat{}, true)
	// Create initial users
	user1 := new_user(db, "ali", "12345")
	user2 := new_user(db, "reza", "12345")
	// Create group with admin user1 and member user2
	chatId := new_group(db, "test-group", user1)
	join_chat(db, chatId, user2)
	// Create private chat with user1 , user2
	chatId2 := new_chat(db, false)
	join_chat(db, chatId2, user1)
	join_chat(db, chatId2, user2)
	// Create message for private chat
	msgId1 := new_message(db, chatId2, user1, "hello user2")

	assign_attachment_to_user(db, msgId1, "file1.jpg", "image")

	new_replay(db, chatId2, user2, msgId1, "thanks and you?")

	for i := 0; i < 500; i++ {
		new_message(db, chatId, user1, fmt.Sprintf("message %d", i+1))
	}

}
