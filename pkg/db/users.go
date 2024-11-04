package database

import (
	"api-channel/pkg/models"
)

func (db *ChatDatabase) CreateUser(user models.User) (uint32, error) {

	err := db.gormDB.Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil

}

func (db *ChatDatabase) CheckUser(username, password string) (uint32, error) {

	var result models.User

	err := db.gormDB.Where("username = ? AND password = ?", username, password).First(&result).Error
	if err != nil {
		return 0, err
	}
	return result.ID, nil
}
