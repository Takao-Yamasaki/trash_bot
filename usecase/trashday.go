package usecase

import (
	"trash_bot/domain/model"
	"trash_bot/domain/repository"

)

type TrashDayUseCase interface {
	GetTrashDay(id int) (result *model.TrashDay, err error)
	GetTrashDays() (result []model.TrashDay, err error)
	CreateTrashDay(day string, trash string) error
	UpdateTrashDay(id int, day string, trash string) error
	DeleteTrashDay(id int) error
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
func (tu *trashDayUseCase) GetTrashDay(id int) (result *model.TrashDay, err error) {
	td, err := tu.trashDayRepository.GetTrashDay(id)
	if err != nil {
		return nil, err
	}
	return td, nil
}

// 一覧の取得
func (tu *trashDayUseCase) GetTrashDays() (result []model.TrashDay, err error) {
	tds, err := tu.trashDayRepository.GetTrashDays()
	if err != nil {
		return nil, err
	}
	return tds, nil
}

// 登録
func (tu *trashDayUseCase) CreateTrashDay(day string, trash string) error {
	td := model.TrashDay{Day: day, Trash: trash}
	err := tu.trashDayRepository.Create(td)
	if err != nil {
		return err
	}
	return nil
}

// 更新
func (tu *trashDayUseCase) UpdateTrashDay(id int, day string, trash string) error {
	td, err := tu.trashDayRepository.GetTrashDay(id)
	if err != nil {
		return err
	}
	
	td.Day = day
	td.Trash = trash
	err = tu.trashDayRepository.Update(*td)
	if err != nil {
		return err
	}
	return nil
}

// 削除
func (tu *trashDayUseCase) DeleteTrashDay(id int) error {
	td, err := tu.trashDayRepository.GetTrashDay(id)
	if err != nil {
		return err
	}

	err = tu.trashDayRepository.Delete(*td)
	if err != nil {
		return err
	}
	return nil
}