package model

import (
	sqlite "trash_bot/config/database"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Contents string
}

func GetComment(id int) Comment {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	var comment Comment
	db.First(&comment)
	connect.Close()
	return comment
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

// 登録
func (comment *Comment) Create() {
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Create(comment)
	connect.Close()
}

func (comment *Comment) Update() {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Save(comment)
	connect.Close()
}

func (comment *Comment) Delete() {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Delete(comment)
	connect.Close()
}
