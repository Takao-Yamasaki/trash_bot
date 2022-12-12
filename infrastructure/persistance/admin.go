package persistance

import (
	"gorm.io/gorm"

	"trash_bot/domain/model"
	"trash_bot/domain/repository"
)

type adminPersistance struct {
	Conn *gorm.DB
}

func NewAdminPersistance(conn *gorm.DB, c repository.AdminRepository) *adminPersistance {
	return &adminPersistance{Conn: conn}
}

// 1件の取得
func (am *adminPersistance) GetAdmin(id int) model.Admin {
	var ad model.Admin
	am.Conn.First(&ad, id)

	return ad
}

// 一覧の取得
func (am *adminPersistance) GetAdmins() []model.Admin {
	var ads []model.Admin
	am.Conn.Find(&ads)
	return ads
}

// 登録
func (am *adminPersistance) Create(ad model.Admin) {
	am.Conn.Create(&ad)
}

// 更新
func (am *adminPersistance) Update(ad model.Admin) {
	am.Conn.Save(&ad)
}

// 削除
func (am *adminPersistance) Delete(ad model.Admin) {
	am.Conn.Delete(&ad)
}