package database

import (
	"api-channel/pkg/models"
)

func (db *ChatDatabase) CreateGroup(groupName string, chatID, userID uint32) (uint32, error) {

	g := &models.Group{
		Name:    groupName,
		OwnerID: userID,
		ChatID:  chatID,
	}
	err := db.gormDB.Create(g).Error
	if err != nil {
		return 0, err
	}

	return uint32(g.ID), nil
}
