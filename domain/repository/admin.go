package  repository

import (
	"trash_bot/domain/model/admin"
)

type AdminRepository interface {
	GetAdmin(id string) (result *admin.Admin, err error)
	GetAdmins() (result []admin.Admin, err error)
	InsertAdmin(admin *admin.Admin) error
	UpdateAdmin(admin *admin.Admin) error
	DeleteAdmin(admin *admin.Admin) error
}