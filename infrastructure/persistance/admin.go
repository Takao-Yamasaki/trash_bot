package persistance

import (
	"gorm.io/gorm"

	"trash_bot/domain/model/admin"
	"trash_bot/domain/repository"
	"trash_bot/infrastructure/dto"
)

type adminPersistance struct {
	Conn *gorm.DB
}

func NewAdminPersistance(conn *gorm.DB) repository.AdminRepository {
	return &adminPersistance{Conn: conn}
}

func (ap *adminPersistance) GetAdminByEmail(email string) (result *admin.Admin, err error) {
	var admin dto.Admin
	if result := ap.Conn.Where("email = ?", email).First(&admin); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_admin, err := dto.AdaptAdmin(&admin)
	if err != nil {
		return nil, err
	}

	return result_admin, nil
}

// 1件の取得
func (ap *adminPersistance) GetAdmin(id string) (result *admin.Admin, err error) {
	var ad dto.Admin
	if result := ap.Conn.Where("admin_id = ?", id).First(&ad); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_admin, err := dto.AdaptAdmin(&ad)
	if err != nil {
		return nil, err
	}
	return result_admin, nil
}

// 一覧の取得
func (ap *adminPersistance) GetAdmins() (result []admin.Admin, err error) {
	
var ads []*dto.Admin
	if result := ap.Conn.Find(&ads); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_admins, err := dto.AdaptAdmins(ads)
	if err != nil {
		return nil, err
	}

	return result_admins, nil
}

// 登録
func (ap *adminPersistance) InsertAdmin(ad *admin.Admin) error {
	converted_admin := dto.ConvertAdmin(ad)

	if result := ap.Conn.Create(converted_admin); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

// 更新
func (ap *adminPersistance) UpdateAdmin(ad *admin.Admin) error {
	converted_admin := dto.ConvertAdmin(ad)

	if result := ap.Conn.Where("admin_id = ?", converted_admin.AdminId).Updates(converted_admin); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

// 削除
func (ap *adminPersistance) DeleteAdmin(ad *admin.Admin) error {
	converted_admin := dto.ConvertAdmin(ad)

	if result := ap.Conn.Where("admin_id = ?", converted_admin.AdminId).Delete(&ad); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}