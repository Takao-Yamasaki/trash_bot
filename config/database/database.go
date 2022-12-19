package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TrashDay struct {
	gorm.Model
	TrashDayId string
	Day   string
	Trash string
}

type Admin struct {
	gorm.Model
	AdminId string
	Name     string
	Email    string
	Password string
}

type Comment struct {
	gorm.Model
	CommentId string
	Contents string
}

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("trashbot.db"), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}
	// マイグレートの実行
	db.AutoMigrate(&TrashDay{})
	db.AutoMigrate(&Admin{})
	db.AutoMigrate(&Comment{})

	return db
}
