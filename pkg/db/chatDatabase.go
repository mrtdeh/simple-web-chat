package database

import (
	"database/sql"

	"gorm.io/gorm"
)

type ChatDatabase struct {
	gormDB *gorm.DB
}

func (db *ChatDatabase) ScanRows(rows *sql.Rows, dest interface{}) error {
	return db.gormDB.ScanRows(rows, dest)
}
func (db *ChatDatabase) GORM() *gorm.DB {
	return db.gormDB
}
