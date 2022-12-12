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
func (tr *trashDayPersistance) GetTrashDay(id int) model.TrashDay {
	var td model.TrashDay
	tr.Conn.First(&td, id)

	return td
}

// 一覧の取得
func (tr *trashDayPersistance) GetTrashDays() []model.TrashDay {
	var tds []model.TrashDay
	tr.Conn.Find(&tds)
	return tds
}

// 登録
func (tr *trashDayPersistance) Create(td model.TrashDay) {
	tr.Conn.Create(&td)
}

// 更新
func (tr *trashDayPersistance) Update(td model.TrashDay) {
	tr.Conn.Save(&td)
}

// 削除
func (tr *trashDayPersistance) Delete(td model.TrashDay) {
	tr.Conn.Delete(&td)
}