package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null"`
	Password       string `gorm:"not null"`
	DisplayName    string
	PhoneNumber    string `gorm:"unique"`
	Email          string `gorm:"unique"`
	ProfilePicture string
	Bio            string
	CreatedAt      time.Time
	LastLogin      time.Time
	Chats          []Chat `gorm:"many2many:chat_members;"`
}

type Chat struct {
	gorm.Model
	Name     string
	IsGroup  bool   `gorm:"default:false"`
	Members  []User `gorm:"many2many:chat_members;"`
	Messages []Message
}

type ChatMember struct {
	ChatID   uint   `gorm:"primaryKey"`
	UserID   uint   `gorm:"primaryKey"`
	Role     string `gorm:"default:'member'"`
	JoinedAt time.Time
}

type Message struct {
	gorm.Model
	ChatID      uint `gorm:"not null"`
	SenderID    uint `gorm:"not null"`
	Sender      User `gorm:"foreignKey:SenderID"`
	Content     string
	MessageType string `gorm:"default:'text'"` // text|file
	Attachments []Attachment
	Replies     []Reply
	Stickers    []Sticker
}

type LastMessageRead struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	ChatID    uint `gorm:"not null"`
	MessageID uint `gorm:"not null"`
}

type Sticker struct {
	gorm.Model
	MessageID uint    `gorm:"not null"` // message ID
	SenderID  uint    `gorm:"not null"`
	Gesture   Gesture `gorm:"not null"`
}

type Reply struct {
	gorm.Model
	MessageID      uint    `gorm:"not null"` // message ID
	ReplyMessageId uint    `gorm:"not null"` // replied message id
	ReplyMessage   Message `gorm:"foreignKey:ReplyMessageId"`
	Thumbnails     []ReplyThumbnails
}

type ReplyThumbnails struct {
	gorm.Model
	ReplyID     uint      `gorm:"not null"` // replied message id
	ThumbnailId uint      `gorm:"not null"`
	Thumbnail   Thumbnail `gorm:"foreignKey:ThumbnailId"`
}

type Attachment struct {
	gorm.Model
	MessageID  uint `gorm:"not null"`
	FilePath   string
	FileType   string // image|video|pdf|...
	FileSize   int
	Thumbnails []Thumbnail
}

type Thumbnail struct {
	gorm.Model
	AttachmentID uint   `gorm:"not null"`
	Base64       string // base64 of image file
	Type         string // mini|small|placeholder
}

type Group struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	OwnerID     uint      `gorm:"not null"`
	ChatID      uint      `gorm:"not null"`
	Chat        Chat      `gorm:"foreignKey:ChatID"`
	Messages    []Message `gorm:"foreignKey:ChatID"`
	Members     []User    `gorm:"many2many:chat_members;foreignKey:ChatID;joinForeignKey:ChatID;"`
}

// type GroupMember struct {
// 	GroupID  uint   `gorm:"primaryKey"`
// 	UserID   uint   `gorm:"primaryKey"`
// 	Role     string `gorm:"default:'member'"`
// 	JoinedAt time.Time
// }
