package trashday

import (
	"fmt"
	"github.com/google/uuid"
)

type TrashDay struct {
	trashDayId trashDayId
	day        day
	trash      trash
}

type trashDayId string
type day string
type trash string

func New(trashDayId string, day string, trash string) (*TrashDay, error) {
	createdTrashDayId, err := newTrashDayId(trashDayId)
	if err != nil {
		return nil, err
	}

	createdDay, err := newDay(day)
	if err != nil {
		return nil, err
	}

	createdTrash, err := newTrash(trash)
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

// valueコンストラクタ
func newTrashDayId(value string) (*trashDayId, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:trashDayId NewTrashDayId()")
		return nil, err
	}

	trashDayId := trashDayId(value)

	return &trashDayId, nil
}

func newDay(value string) (*day, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:day newDay()")
		return nil, err
	}

	day := day(value)

	return &day, nil
}

func newTrash(value string) (*trash, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:trash newTrash()")
		return nil, err
	}
	trash := trash(value)

	return &trash, nil
}
