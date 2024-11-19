package seeds

import (
	"api-channel/pkg/models"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {
	// Clear records of table
	truncateTable(db, models.Reply{}, true)
	truncateTable(db, models.Message{}, true)
	truncateTable(db, models.ChatMember{}, false)
	truncateTable(db, models.User{}, true)
	truncateTable(db, models.Group{}, true)
	truncateTable(db, models.Chat{}, true)
	// Create initial users
	user1 := newUser(db, "ali", "12345")
	user2 := newUser(db, "reza", "12345")
	// Create group with admin user1 and member user2
	chatId := newGroup(db, "test-group", user1)
	joinChat(db, chatId, user2)
	// Create private chat with user1 , user2
	chatId2 := newChat(db, false)
	joinChat(db, chatId2, user1)
	joinChat(db, chatId2, user2)
	// Create message for private chat
	msgId1 := newMessage(db, chatId2, user1, "hello user2")
	newReplay(db, chatId2, user2, msgId1, "thanks and you?")

	newMessage(db, chatId, user1, "hello every one...")

}
