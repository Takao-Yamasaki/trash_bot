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
func (am *adminPersistance) GetAdmin(id int) (result *model.Admin, err error) {
	var ad model.Admin
	if result := am.Conn.First(&ad, id); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &ad, nil
}

// 一覧の取得
func (am *adminPersistance) GetAdmins() (result []model.Admin, err error) {
	
	var ads []model.Admin
	if result := am.Conn.Find(&ads); result.Error != nil {
		err := result.Error
		return nil, err
	}
	return ads, nil
}

// 登録
func (am *adminPersistance) Create(ad model.Admin) error {
	if result := am.Conn.Create(&ad); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

// 更新
func (am *adminPersistance) Update(ad model.Admin) error {
	if result := am.Conn.Save(&ad); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

// 削除
func (am *adminPersistance) Delete(ad model.Admin) error {
	if result := am.Conn.Delete(&ad); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}