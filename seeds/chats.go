package seeds

import (
	"api-channel/pkg/models"
	"fmt"

	"gorm.io/gorm"
)

func ChatSeeds(db *gorm.DB) {
	// Clear records of table
	db.Where("true").Unscoped().Delete(&models.Chat{})
	// Create sample records
	chats := []models.Chat{
		{Name: "group", IsGroup: true},
		{Name: "private", IsGroup: false},
	}
	// Create records in the database
	for _, chat := range chats {
		err := db.Create(&chat).Error
		if err != nil {
			fmt.Printf("Error when create chat: %s : %v\n", chat.Name, err)
			return
		}
	}
}
