package model

import (
	sqlite "trash_bot/config/database"

	"gorm.io/gorm"
)

// admin_idの追加
type TrashDay struct {
	gorm.Model
	Day  string
	Trash string
}

// １件の取得
func GetTrashDay(id int) TrashDay {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	var trashday TrashDay
	db.First(&trashday)
	connect.Close()
	return trashday
}

// 一覧の取得
func GetTrashDays() ([]TrashDay) {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	var trashdays []TrashDay
	db.Find(&trashdays)
	connect.Close()
	return trashdays
}

// 登録
func (trashday *TrashDay) Create() {
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Create(trashday)
	connect.Close()
}

// 更新
func (trashday *TrashDay) Update() {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Save(trashday)
	connect.Close()
}

// 削除
func (trashday *TrashDay) Delete() {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Delete(trashday)
	connect.Close()
}
