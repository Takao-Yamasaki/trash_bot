package model

import (
	sqlite "trash_bot/config/database"

	"gorm.io/gorm"
)

type TrashDay struct {
	gorm.Model
	Week  string
	Trash string
}

// １件の取得
func GetTrashDay(id int) TrashDay {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	var td TrashDay
	db.First(&td)
	connect.Close()
	return td
}

// 一覧の取得
func GetTrashDays() []TrashDay {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	var tds []TrashDay
	db.Find(&tds)
	connect.Close()
	return tds
}

// 登録
func (td *TrashDay) Create() {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Create(td)
	connect.Close()
}

// 更新
func (td *TrashDay) Update() {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Save(td)
	connect.Close()
}

// 削除
func (td *TrashDay) Delete() {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Delete(td)
	connect.Close()
}
