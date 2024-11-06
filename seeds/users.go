package seeds

import (
	"api-channel/pkg/models"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func UserSeeds(db *gorm.DB) {
	now := time.Now()
	// Clear records of table
	truncateTable(db, models.ChatMember{})
	truncateTable(db, models.User{})
	truncateTable(db, models.Group{})
	truncateTable(db, models.Chat{})

	user1 := models.User{Username: "ali", Password: "12345", PhoneNumber: faker.Phonenumber(), Email: faker.Email(), DisplayName: faker.Name()}
	user2 := models.User{Username: "reza", Password: "12345", PhoneNumber: faker.Phonenumber(), Email: faker.Email(), DisplayName: faker.Name()}
	db.Create(&user1)
	db.Create(&user2)

	chat1 := models.Chat{IsGroup: true}
	chat2 := models.Chat{IsGroup: false}
	db.Create(&chat1)
	db.Create(&chat2)

	group1 := models.Group{Name: "test-group", OwnerID: user1.ID, ChatID: chat1.ID}
	db.Create(&group1)

	db.Create(&models.ChatMember{ChatID: chat1.ID, UserID: user1.ID, Role: "admin", JoinedAt: now})
	db.Create(&models.ChatMember{ChatID: chat1.ID, UserID: user2.ID, Role: "member", JoinedAt: now})

	db.Create(&models.ChatMember{ChatID: chat2.ID, UserID: user1.ID, Role: "member", JoinedAt: now})
	db.Create(&models.ChatMember{ChatID: chat2.ID, UserID: user2.ID, Role: "member", JoinedAt: now})
}

func truncateTable(db *gorm.DB, model interface{}) error {
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
	if err := db.Exec(fmt.Sprintf("SELECT pg_get_serial_sequence('%s', 'id');", table)).Error; err == nil {
		err = db.Exec(fmt.Sprintf("SELECT setval('%s_id_seq', 1, false);", table)).Error
		if err != nil {
			panic(err)
		}
	}
	return nil
}
