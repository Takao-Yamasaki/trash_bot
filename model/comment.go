package model

import (
	"gorm.io/gorm"
	sqlite "trash_bot/config/database"
)

type Comment struct {
	gorm.Model
	Contents string
}

func GetComments() []Comment {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	var comments []Comment
	db.Find(&comments)
	connect.Close()
	return comments
}

func (comment *Comment) Create() {
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Create(comment)
	connect.Close()
}
