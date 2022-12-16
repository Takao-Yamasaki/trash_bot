package admin

import (
	"fmt"
	"github.com/google/uuid"
)

type Admin struct {
	adminId  adminId
	name	 name
	email    email
	password password
}

type adminId string
type name string
type email string
type password string

func New(adminId string, name string, email string, password string) (*Admin, error) {
	createdAdminId, err := newAdminId(adminId)
	if err != nil {
		return nil, err
	}

	createdName, err := newName(name)
	if err != nil {
		return nil, err
	}

	createdEmail, err := newEmail(email)
	if err != nil {
		return nil, err
	}

	createdPassword, err := newPassword(password)
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

// value constructors
func newAdminId(value string) (*adminId, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:adminId NewAdminId()")
		return nil, err
	}

	adminId := adminId(value)

	return &adminId, nil
}

func newName(value string) (*name, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:name newName()")
		return nil, err
	}

	name := name(value)

	return &name, nil
}

func newEmail(value string) (*email, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:email newEmail()")
		return nil, err
	}

	email := email(value)

	return &email, nil
}

func newPassword(value string) (*password, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:password newPassword()")
		return nil, err
	}

	password := password(value)
	
	return &password, nil
}