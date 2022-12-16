package dto

import (
	"time"
	"trash_bot/domain/model/admin"
)

type Admin struct {
	ID  int
	AdminId string
	Name string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func ConvertAdmin(admin *admin.Admin) *Admin {
	return &Admin{
		AdminId: admin.GetAdminId(),
		Name: admin.GetName(),
		Email: admin.GetEmail(),
		Password: admin.GetPassword(),
	}
}

func AdaptAdmin(converted_admin *Admin) (*admin.Admin, error) {
	admin, err := admin.New(
		converted_admin.AdminId,
		converted_admin.Name,
		converted_admin.Email,
		converted_admin.Password,
	)

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func AdaptAdmins(converted_admins []*Admin) ([]admin.Admin, error) {
	var admins []admin.Admin

	for _, converted_admin := range converted_admins {
		admin, err := admin.New(
			converted_admin.AdminId, 
			converted_admin.Name, 
			converted_admin.Email,
			converted_admin.Password,
		)

		if err != nil {
			return nil, err
		}
		admins = append(admins, *admin)
	}
	
	return admins, nil
}