package model

import (
	sqlite "trash_bot/config/database"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func GetAdmin(id int) Admin {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	var admin Admin
	db.First(&admin)
	connect.Close()
	return admin
}

func GetAdmins() []Admin {
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

func (admin *Admin) Update() {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Save(admin)
	connect.Close()
}

func (admin *Admin) Delete() {
	db := sqlite.New()
	connect, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Delete(admin)
	connect.Close()
}
