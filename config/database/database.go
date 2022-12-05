package database

import (
	"github.com/mattn/go-sqlite3"
	"github.com/jinzhu/gorm"
)

type TrushDay struct {
	gorm.Model
	DayOfWeek string
	TypeOfTrash string
}

func New() {
	db, err := gorm.Open(sqlite.Open("trashday.db"), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	db.AutoMigrate(&TrushDay{})

	return db
}