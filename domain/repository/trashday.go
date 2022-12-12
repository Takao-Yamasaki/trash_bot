package repository

import (
	"trash_bot/domain/model"
)

type TrashDayRepository interface {
	GetTrashDay(id int) model.TrashDay
	GetTrashDays() []model.TrashDay
	Create(td model.TrashDay)
	Update(td model.TrashDay)
	Delete(td model.TrashDay)
}