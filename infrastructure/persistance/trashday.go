package persistance

import (
	"gorm.io/gorm"

	"trash_bot/domain/model/trashday"
	"trash_bot/domain/repository"
	"trash_bot/infrastructure/dto"
)

type trashDayPersistance struct {
	Conn *gorm.DB
}

func NewTrashDayPersistance(conn *gorm.DB) repository.TrashDayRepository {
	return &trashDayPersistance{Conn: conn}
}

// １件の取得
func (tp *trashDayPersistance) GetTrashDay(id string) (result *trashday.TrashDay, err error) {
	var td dto.TrashDay
	if result := tp.Conn.Where("trash_day_id = ?", id).First(&td); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_trashDay, err := dto.AdaptTrashDay(&td)
	if err != nil {
		return nil, err
	}

	return result_trashDay, nil
}

// 一覧の取得
func (tp *trashDayPersistance) GetTrashDays() (result []trashday.TrashDay, err error) {

	var tds []*dto.TrashDay
	if result := tp.Conn.Find(&tds); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_trashDays, err := dto.AdaptTrashDays(tds)
	if err != nil {
		return nil, err
	}

	return result_trashDays, nil
}

func (tp *trashDayPersistance) InsertTrashDay(td *trashday.TrashDay) error {
	converted_trashDay := dto.ConvertTrashDay(td)

	if result := tp.Conn.Create(converted_trashDay); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (tp *trashDayPersistance) UpdateTrashDay(td *trashday.TrashDay) error {
	converted_trashDay := dto.ConvertTrashDay(td)

	if result := tp.Conn.Where("trash_day_id = ?", converted_trashDay.TrashDayId).
		Updates(converted_trashDay); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

// 削除
func (tp *trashDayPersistance) DeleteTrashDay(td *trashday.TrashDay) error {
	converted_trashDay := dto.ConvertTrashDay(td)

	if result := tp.Conn.Where("trash_day_id = ?", converted_trashDay.TrashDayId).Delete(converted_trashDay); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}
