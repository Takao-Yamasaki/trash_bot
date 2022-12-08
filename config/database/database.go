package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TrashDay struct {
	gorm.Model
	Day  string
	Trash string
}

type Admin struct {
	gorm.Model
	Name string
	Email string
	Password string
}

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("trashbot.db"), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}
	// マイグレートの実行
	db.AutoMigrate(&TrashDay{})
	db.AutoMigrate(&Admin{})

	return db
}
