package  repository

import (
	"trash_bot/domain/model"
)

type AdminRepository interface {
	GetAdmin(id int) (result *model.Admin, err error)
	GetAdmins() (result []model.Admin, err error)
	Create(admin model.Admin) error
	Update(admin model.Admin) error
	Delete(admin model.Admin) error
}