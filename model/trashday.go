package model

import (
	sqlite "trash_bot/config/database"

	"gorm.io/gorm"
)

type Trashday struct {
	gorm.Model
	Dayofweek   string
	Typeoftrash string
}

func GetTrashday(id int) Trashday {
	db := sqlite.New()
	
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}

	var trashday Trashday
	db.First(&trashday)

	connect.Close()
	
	return trashday
}

// 一覧の取得
func GetTrashdays() []Trashday {
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}

	var trashdays []Trashday
	db.Find(&trashdays)

	connect.Close()

	return trashdays
}

// 登録
func (c *Trashday) Create() {
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}

	db.Create(c)

	connect.Close()
}
