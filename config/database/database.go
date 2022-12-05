package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TrashDay struct {
	gorm.Model
	Dayofweek string
	Typeoftrash string
}

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("trashday.db"), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}
	// マイグレートの実行
	db.AutoMigrate(&TrashDay{})

	return db
}