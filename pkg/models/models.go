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
}

type Chat struct {
	gorm.Model
	Name      string
	IsGroup   bool `gorm:"default:false"`
	CreatedAt time.Time
	Members   []User `gorm:"many2many:chat_members;"`
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
	Content     string
	MessageType string `gorm:"default:'text'"` // text|file
	CreatedAt   time.Time
}

type Reply struct {
	gorm.Model
	MessageID uint   `gorm:"not null"` // message ID
	ReplyToID uint   `gorm:"not null"` // message|attachment ID
	Type      string `gorm:"not null"` // text|file
}

type Attachment struct {
	gorm.Model
	MessageID uint `gorm:"not null"`
	FilePath  string
	FileType  string // image|video|pdf|...
	FileSize  int
}

type Thumbnail struct {
	gorm.Model
	AttachmentID uint   `gorm:"not null"`
	Base64       string // base64 of image file
	Type         string // mini|small|blur
}

type Group struct {
	gorm.Model
	Name        string `gorm:"not null"`
	OwnerID     uint   `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	Members     []User `gorm:"many2many:group_members;"`
}

type GroupMember struct {
	GroupID  uint   `gorm:"primaryKey"`
	UserID   uint   `gorm:"primaryKey"`
	Role     string `gorm:"default:'member'"`
	JoinedAt time.Time
}
