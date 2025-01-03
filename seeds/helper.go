package seeds

import (
	"api-channel/pkg/conf"
	"api-channel/pkg/helper"
	"api-channel/pkg/models"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"reflect"
	"time"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func new_user(db *gorm.DB, user, pass string) uint32 {
	u := models.User{Username: user, Password: pass, PhoneNumber: faker.Phonenumber(), Email: faker.Email(), DisplayName: faker.Name()}
	db.Create(&u)
	return u.ID
}

func join_chat(db *gorm.DB, chatId, member uint32) {
	now := time.Now()
	db.Create(&models.ChatMember{ChatID: chatId, UserID: member, Role: "member", JoinedAt: now})
}

func new_chat(db *gorm.DB, chatType string) (chatId uint32) {
	c := models.Chat{Type: chatType}
	db.Create(&c)

	chatId = c.ID
	return
}

func new_personal_chat(db *gorm.DB) (chatId uint32) {
	return new_chat(db, "personal")
}

func new_group(db *gorm.DB, name, groupType string, owner uint32) (chatId uint32) {
	now := time.Now()

	chatId = new_chat(db, groupType)

	g := models.Group{Name: name, OwnerID: owner, ChatID: chatId}
	db.Create(&g)

	db.Create(&models.ChatMember{ChatID: chatId, UserID: owner, Role: "admin", JoinedAt: now})

	return
}

func new_message(db *gorm.DB, chatId, senderId uint32, text string) (msgId uint32) {
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

func new_replay(db *gorm.DB, chatId, senderId, repliedMsgId uint32, text string) (msgId uint32) {
	msgId = new_message(db, chatId, senderId, text)
	replyMsg := &models.Reply{
		MessageID:      msgId,
		ReplyMessageId: repliedMsgId,
	}
	db.Create(&replyMsg)
	return msgId
}

func truncate_table(db *gorm.DB, model interface{}, resetId bool) error {
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

func assign_attachment_to_user(db *gorm.DB, msgId uint32, filepath, fileType string) {
	url, _ := url.JoinPath(conf.UPLOAD_SERVER, path.Base(filepath))
	// Write file path in Attachments table
	attachment := models.Attachment{
		MessageID: msgId,
		FilePath:  url,
		FileType:  fileType,
		FileSize:  3200,
	}
	err := db.Create(&attachment).Error
	if err != nil {
		log.Fatal(err)
	}

	img, err := helper.OpenImage(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		log.Fatal(err)
	}

	// generate small size 64x64
	thum64 := helper.Thumbnail(img, 64)
	err = db.Create(&models.Thumbnail{
		AttachmentID: attachment.ID,
		Base64:       helper.Base64Image(thum64, 1),
		Type:         "small",
	}).Error
	if err != nil {
		log.Fatal(err)
	}

	// generate mini size 32x32
	thum32 := helper.Thumbnail(img, 32)
	err = db.Create(&models.Thumbnail{
		AttachmentID: attachment.ID,
		Base64:       helper.Base64Image(thum32, 1),
		Type:         "mini",
	}).Error
	if err != nil {
		log.Fatal(err)
	}

	placeholder := helper.Placeholder(img)
	err = db.Create(&models.Thumbnail{
		AttachmentID: attachment.ID,
		Base64:       helper.Base64Image(placeholder, 1),
		Type:         "placeholder",
	}).Error
	if err != nil {
		log.Fatal(err)
	}

}
