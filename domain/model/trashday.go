package model

import (
	"gorm.io/gorm"
)

// admin_idの追加
type TrashDay struct {
	gorm.Model
	Day  string
	Trash string
}