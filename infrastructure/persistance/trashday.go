package persistance

import (
	"gorm.io/gorm"

	"trash_bot/domain/model"
	"trash_bot/domain/repository"
)

type trashDayPersistance struct {
	Conn *gorm.DB
}

func NewTrashDayPersistance(conn *gorm.DB, c repository.TrashDayRepository) *trashDayPersistance {
	return &trashDayPersistance{Conn: conn}
}

// １件の取得
func (tr *trashDayPersistance) GetTrashDay(id int) (result *model.TrashDay, err error) {
	var td model.TrashDay
	if result := tr.Conn.First(&td, id); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &td, nil
}

// 一覧の取得
func (tr *trashDayPersistance) GetTrashDays() (result []model.TrashDay, err error) {
	var tds []model.TrashDay
	if result := tr.Conn.Find(&tds); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return tds, nil
}

// 登録
func (tr *trashDayPersistance) Create(td model.TrashDay) error {
	if result := tr.Conn.Create(&td); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

// 更新
func (tr *trashDayPersistance) Update(td model.TrashDay) error {
	if result := tr.Conn.Save(&td); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

// 削除
func (tr *trashDayPersistance) Delete(td model.TrashDay) error {
	if result := tr.Conn.Delete(&td); result.Error != nil {
		err := result.Error
		return err
	}
	
	return nil
}