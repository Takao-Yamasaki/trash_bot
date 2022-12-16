package repository

import (
	"trash_bot/domain/model/trashday"
)

type TrashDayRepository interface {
	GetTrashDay(id string) (result *trashday.TrashDay, err error)
	GetTrashDays() (result []trashday.TrashDay, err error)
	InsertTrashDay(td *trashday.TrashDay) error
	UpdateTrashDay(td *trashday.TrashDay) error
	DeleteTrashDay(td *trashday.TrashDay) error
}