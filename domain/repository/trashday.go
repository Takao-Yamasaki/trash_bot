package repository

import (
	"trash_bot/domain/model"
)

type TrashDayRepository interface {
	GetTrashDay(id int) (result *model.TrashDay, err error)
	GetTrashDays() (result []model.TrashDay, err error)
	Create(td model.TrashDay) error
	Update(td model.TrashDay) error
	Delete(td model.TrashDay) error
}