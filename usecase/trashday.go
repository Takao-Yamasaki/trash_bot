package usecase

import (
	"trash_bot/domain/model/trashday"
	"trash_bot/domain/repository"

)

type TrashDayUseCase interface {
	GetTrashDay(id string) (result *trashday.TrashDay, err error)
	GetTrashDays() (result []trashday.TrashDay, err error)
	CreateTrashDay(day string, trash string) error
	UpdateTrashDay(id string, day string, trash string) error
	DeleteTrashDay(id string) error
}

type trashDayUseCase struct {
	trashDayRepository repository.TrashDayRepository
}

func NewTrashDayUseCase(tr repository.TrashDayRepository) TrashDayUseCase {
	return &trashDayUseCase{
		trashDayRepository: tr,
	}
}

// １件の取得
func (tu *trashDayUseCase) GetTrashDay(id string) (result *trashday.TrashDay, err error) {
	td, err := tu.trashDayRepository.GetTrashDay(id)
	if err != nil {
		return nil, err
	}
	return td, nil
}

// 一覧の取得
func (tu *trashDayUseCase) GetTrashDays() (result []trashday.TrashDay, err error) {
	tds, err := tu.trashDayRepository.GetTrashDays()
	if err != nil {
		return nil, err
	}
	return tds, nil
}

// 登録
func (tu *trashDayUseCase) CreateTrashDay(day string, trash string) error {
	td, err := trashday.Create(day, trash)
	if err != nil {
		return err
	}
	err = tu.trashDayRepository.InsertTrashDay(td)
	if err != nil {
		return err
	}
	return nil
}

// 更新
func (tu *trashDayUseCase) UpdateTrashDay(id string, day string, trash string) error {
	current_trashDay, err := tu.trashDayRepository.GetTrashDay(id)
	if err != nil {
		return err
	}

	trashDayId := string(current_trashDay.GetTrashDayId())

	update_trashDay, err := trashday.New(trashDayId, day, trash)
	if err != nil {
		return err
	}

	err = tu.trashDayRepository.UpdateTrashDay(update_trashDay)
	if err != nil {
		return err
	}
	return nil
}

// 削除
func (tu *trashDayUseCase) DeleteTrashDay(id string) error {
	td, err := tu.trashDayRepository.GetTrashDay(id)
	if err != nil {
		return err
	}

	err = tu.trashDayRepository.DeleteTrashDay(td)
	if err != nil {
		return err
	}
	return nil
}