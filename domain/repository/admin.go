package  repository

import (
	"trash_bot/domain/model"
)

type AdminRepository interface {
	GetAdmin(id int) model.Admin
	GetAdmins() []model.Admin
	Create(admin model.Admin)
	Update(admin model.Admin)
	Delete(admin model.Admin)
}