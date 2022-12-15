package usecase

import (
	"trash_bot/domain/model"
	"trash_bot/domain/repository"
)

type AdminUseCase interface {
	GetAdmin(id int) (result *model.Admin, err error)
	GetAdmins() (result []model.Admin, err error)
	CreateAdmin(name string, email string, password string) error
	UpdateAdmin(id int, name string, email string, password string) error
	DeleteAdmin(id int) error
}

type adminUseCase struct {
	adminRepository repository.AdminRepository
}

func NewAdminUseCase(ar repository.AdminRepository) AdminUseCase {
	return &adminUseCase{
		adminRepository: ar,
	}
}

func (au *adminUseCase) GetAdmin(id int) (result *model.Admin, err error) {
	admin, err := au.adminRepository.GetAdmin(id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (au *adminUseCase) GetAdmins() (result []model.Admin, err error) {
	admins, err := au.adminRepository.GetAdmins()
	if err != nil {
		return nil, err
	}

	return admins, nil
}

func (au *adminUseCase) CreateAdmin(name string, email string, password string) error {
	admin := model.Admin{Name: name, Email: email, Password: password}
	err := au.adminRepository.Create(admin)
	if err != nil {
		return err
	}

	return nil
}

func (au *adminUseCase) UpdateAdmin(id int, name string, email string, password string) error {
	admin, err := au.adminRepository.GetAdmin(id)
	if err != nil {
		return err
	}

	admin.Name = name
	admin.Email = email
	admin.Password = password
	err = au.adminRepository.Update(*admin)
	if err != nil {
		return err
	}

	return nil
}

func (au *adminUseCase) DeleteAdmin(id int) error {
	admin, err := au.adminRepository.GetAdmin(id)
	if err != nil {
		return err
	}

	err = au.adminRepository.Delete(*admin)
	if err != nil {
		return err
	}

	return nil
}