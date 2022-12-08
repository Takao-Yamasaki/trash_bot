package model

import (
	sqlite "trash_bot/config/database"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name string
	Email string
	Password string
}

func GetAdmins() ([]Admin) {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	var admins []Admin
	db.Find(&admins)
	connect.Close()
	return admins
}

func (admin *Admin) Create() {
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Create(admin)
	connect.Close()
}