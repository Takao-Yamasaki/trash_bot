package admin

import (
	"trash_bot/domain/model/vo"

	"github.com/google/uuid"
)

type Admin struct {
	adminId  vo.UuId
	name     vo.Name
	email    vo.Email
	password vo.Password
}

func New(adminId string, name string, email string, password string) (*Admin, error) {
	createdAdminId, err := vo.NewUuid(adminId)
	if err != nil {
		return nil, err
	}

	createdName, err := vo.NewName(name)
	if err != nil {
		return nil, err
	}

	createdEmail, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}

	createdPassword, err := vo.NewPassword(password)
	if err != nil {
		return nil, err
	}

	admin := Admin{
		adminId:  *createdAdminId,
		name:     *createdName,
		email:    *createdEmail,
		password: *createdPassword,
	}

	return &admin, nil
}

func Create(name string, email string, password string) (*Admin, error) {
	adminId := uuid.New().String()
	admin, err := New(adminId, name, email, password)

	if err != nil {
		return nil, err
	}

	return admin, err
}

// Getter
func (ad Admin) GetAdminId() string {
	return string(ad.adminId)
}

func (ad Admin) GetName() string {
	return string(ad.name)
}

func (ad Admin) GetEmail() string {
	return string(ad.email)
}

func (ad Admin) GetPassword() string {
	return string(ad.password)
}
