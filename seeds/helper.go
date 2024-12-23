package seeds

import (
	"api-channel/pkg/models"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func newUser(db *gorm.DB, user, pass string) uint32 {
	u := models.User{Username: user, Password: pass, PhoneNumber: faker.Phonenumber(), Email: faker.Email(), DisplayName: faker.Name()}
	db.Create(&u)
	return u.ID
}

func joinChat(db *gorm.DB, chatId, member uint32) {
	now := time.Now()
	db.Create(&models.ChatMember{ChatID: chatId, UserID: member, Role: "member", JoinedAt: now})
}

func newChat(db *gorm.DB, isGroup bool) (chatId uint32) {
	c := models.Chat{IsGroup: isGroup}
	db.Create(&c)

	chatId = c.ID
	return
}

func newGroup(db *gorm.DB, name string, owner uint32) (chatId uint32) {
	now := time.Now()

	chatId = newChat(db, true)

	g := models.Group{Name: name, OwnerID: owner, ChatID: chatId}
	db.Create(&g)

	db.Create(&models.ChatMember{ChatID: chatId, UserID: owner, Role: "admin", JoinedAt: now})

	return
}

func newMessage(db *gorm.DB, chatId, senderId uint32, text string) (msgId uint32) {
	msg := models.Message{
		ChatID:      chatId,
		SenderID:    senderId,
		Content:     text,
		MessageType: "text",
	}
	db.Create(&msg)
	db.Table("chats").
		Where("id = ?", chatId).
		Update("last_message_id", msg.ID)
	return msg.ID
}

func newReplay(db *gorm.DB, chatId, senderId, repliedMsgId uint32, text string) (msgId uint32) {
	msgId = newMessage(db, chatId, senderId, text)
	replyMsg := &models.Reply{
		MessageID:      msgId,
		ReplyMessageId: repliedMsgId,
	}
	db.Create(&replyMsg)
	return msgId
}

func truncateTable(db *gorm.DB, model interface{}, resetId bool) error {
	// Get struct name
	sn := reflect.TypeOf(model).Name()
	// Find table name from struct name of model
	table := db.NamingStrategy.TableName(sn)
	// Delete all records of table
	err := db.Exec(fmt.Sprintf("DELETE FROM %s WHERE true", table)).Error
	if err != nil {
		panic(err)
	}
	// Reset auto increament history
	if !resetId {
		return nil
	}
	if err := db.Exec(fmt.Sprintf("SELECT pg_get_serial_sequence('%s', 'id');", table)).Error; err == nil {
		err = db.Exec(fmt.Sprintf("SELECT setval('%s_id_seq', 1, false);", table)).Error
		if err != nil {
			panic(err)
		}
	}
	return nil
}
