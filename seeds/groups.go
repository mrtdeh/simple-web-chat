package seeds

import (
	"api-channel/pkg/models"
	"fmt"

	"gorm.io/gorm"
)

func GroupSeeds(db *gorm.DB) {
	// Clear records of table
	db.Where("true").Unscoped().Delete(&models.Group{})
	// Create sample records
	groups := []models.Group{
		{Name: "group1", ChatID: 1, OwnerID: 1},
	}
	// Create records in the database
	for _, group := range groups {
		err := db.Create(&group).Error
		if err != nil {
			fmt.Printf("Error when create group: %s : %v\n", group.Name, err)
			return
		}
	}
}
