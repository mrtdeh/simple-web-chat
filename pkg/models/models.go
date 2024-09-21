package models

import "time"

type User struct {
	ID             uint   `gorm:"primaryKey"`
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
	ID        uint `gorm:"primaryKey"`
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
	ID          uint `gorm:"primaryKey"`
	ChatID      uint `gorm:"not null"`
	SenderID    uint `gorm:"not null"`
	Content     string
	MessageType string `gorm:"default:'text'"` // text|file
	CreatedAt   time.Time
}

type Attachment struct {
	ID         uint `gorm:"primaryKey"`
	MessageID  uint `gorm:"not null"`
	FilePath   string
	FileType   string
	FileSize   int
	UploadedAt time.Time
}

type Group struct {
	ID          uint   `gorm:"primaryKey"`
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
