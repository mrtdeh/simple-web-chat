package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint32 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Model
	Username       string `gorm:"unique;not null"`
	Password       string `gorm:"not null"`
	DisplayName    string
	PhoneNumber    string `gorm:"unique"`
	Email          string `gorm:"unique"`
	ProfilePicture string
	Bio            string
	Chats          []Chat `gorm:"many2many:chat_members;"`
}

type Chat struct {
	Model
	Name          string
	Type          string `gorm:"not null"` // e.g., "channel", "private", "public", "personal"
	Members       []User `gorm:"many2many:chat_members;"`
	Messages      []Message
	Group         *Group `gorm:"foreignKey:ChatID"` // رابطه با جدول گروه
	LastMessageID uint
}

type Group struct {
	Model
	Name        string `gorm:"not null"`
	Description string
	AvatarUrl   string
	OwnerID     uint32 `gorm:"not null"`
	ChatID      uint32 `gorm:"not null"`
}

type Token struct {
	Model
	UserID     uint32    `gorm:"not null"`
	Value      string    `gorm:"not null"`
	ExpireDate time.Time `gorm:"not null"`
}

type ChatMember struct {
	ChatID              uint32 `gorm:"primaryKey"`
	Chat                Chat   `gorm:"foreignKey:ChatID"`
	UserID              uint32 `gorm:"primaryKey"`
	User                User   `gorm:"foreignKey:UserID"`
	Role                string `gorm:"default:'member'"`
	Mute                bool   `gorm:"default:false"`
	LastReadedMessageID uint32 `gorm:"not null"`
	JoinedAt            time.Time
}

type Message struct {
	Model
	ChatID      uint32 `gorm:"not null"`
	SenderID    uint32 `gorm:"not null"`
	Sender      User   `gorm:"foreignKey:SenderID"`
	Content     string
	MessageType string `gorm:"default:'text'"` // text|file
	Attachments []Attachment
	Replies     []Reply
	Stickers    []Sticker
}

type Sticker struct {
	Model
	MessageID uint32  `gorm:"not null"` // message ID
	SenderID  uint32  `gorm:"not null"`
	Gesture   Gesture `gorm:"not null"`
}

type Reply struct {
	Model
	MessageID      uint32  `gorm:"not null"` // message ID
	ReplyMessageId uint32  `gorm:"not null"` // replied message id
	ReplyMessage   Message `gorm:"foreignKey:ReplyMessageId"`
	Thumbnails     []ReplyThumbnails
}

type ReplyThumbnails struct {
	Model
	ReplyID     uint32    `gorm:"not null"` // replied message id
	ThumbnailId uint32    `gorm:"not null"`
	Thumbnail   Thumbnail `gorm:"foreignKey:ThumbnailId"`
}

type Attachment struct {
	Model
	MessageID  uint32 `gorm:"not null"`
	FilePath   string
	FileType   string // image|video|pdf|...
	FileSize   int
	Thumbnails []Thumbnail
}

type Thumbnail struct {
	Model
	AttachmentID uint32 `gorm:"not null"`
	Base64       string // base64 of image file
	Type         string // mini|small|placeholder
}
