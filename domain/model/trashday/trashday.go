package trashday

import (
	"github.com/google/uuid"
	"trash_bot/domain/model/vo"
)

type TrashDay struct {
	trashDayId vo.UuId
	day        vo.Day
	trash      vo.Trash
}

type trashDayId string
type day string
type trash string

func New(trashDayId string, day string, trash string) (*TrashDay, error) {
	createdTrashDayId, err := vo.NewUuid(trashDayId)
	if err != nil {
		return nil, err
	}

	createdDay, err := vo.NewDay(day)
	if err != nil {
		return nil, err
	}

	createdTrash, err := vo.NewTrash(trash)
	if err != nil {
		return nil, err
	}

	trashDay := TrashDay{
		trashDayId: *createdTrashDayId,
		day:        *createdDay,
		trash:      *createdTrash,
	}

	return &trashDay, nil
}

// コンストラクタ：構造体を返す
func Create(day string, trash string) (*TrashDay, error) {
	trashDayId := uuid.New().String()
	trashDay, err := New(trashDayId, day, trash)

	if err != nil {
		return nil, err
	}

	return trashDay, err
}

// Getter
func (td TrashDay) GetTrashDayId() string {
	return string(td.trashDayId)
}

func (td TrashDay) GetDay() string {
	return string(td.day)
}

func (td TrashDay) GetTrash() string {
	return string(td.trash)
}
