package dto

import (
	"time"
	"trash_bot/domain/model/trashday"
)

type TrashDay struct {
	ID         int
	TrashDayId string
	Day        string
	Trash      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func ConvertTrashDay(td *trashday.TrashDay) *TrashDay {
	return &TrashDay{
		TrashDayId: string(td.GetTrashDayId()),
		Day:        string(td.GetDay()),
		Trash:      string(td.GetTrash()),
	}
}

func AdaptTrashDay(converted_trashDay *TrashDay) (*trashday.TrashDay, error) {
	trashDay, err := trashday.New(
		converted_trashDay.TrashDayId,
		converted_trashDay.Day,
		converted_trashDay.Trash,
	)

	if err != nil {
		return nil, err
	}

	return trashDay, nil
}

func AdaptTrashDays(converted_trashDays []*TrashDay) ([]trashday.TrashDay, error) {
	var trashDays []trashday.TrashDay

	for _, converted_trashDay := range converted_trashDays {
		trashDay, err := trashday.New(
			converted_trashDay.TrashDayId,
			converted_trashDay.Day,
			converted_trashDay.Trash,
		)

		if err != nil {
			return nil, err
		}

		trashDays = append(trashDays, *trashDay)
	}

	return trashDays, nil
}
