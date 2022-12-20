package usecase

import (
	"trash_bot/domain/model/admin"
	"trash_bot/domain/repository"
)

type AdminUseCase interface {
	GetAdminForAuth(email string) (result *admin.Admin, err error)
	GetAdmin(id string) (result *admin.Admin, err error)
	GetAdmins() (result []admin.Admin, err error)
	CreateAdmin(name string, email string, password string) error
	UpdateAdmin(id string, name string, email string, password string) error
	DeleteAdmin(id string) error
}

type adminUseCase struct {
	adminRepository repository.AdminRepository
}

func NewAdminUseCase(ar repository.AdminRepository) AdminUseCase {
	return &adminUseCase{
		adminRepository: ar,
	}
}

func (au *adminUseCase) GetAdmin(id string) (result *admin.Admin, err error) {
	admin, err := au.adminRepository.GetAdmin(id)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (au *adminUseCase) GetAdminForAuth(email string) (result *admin.Admin, err error) {
	current_admin, err := au.adminRepository.GetAdminByEmail(email)
	if err != nil {
		return nil, err
	}

	return  current_admin, nil
}

func (au *adminUseCase) GetAdmins() (result []admin.Admin, err error) {
	admins, err := au.adminRepository.GetAdmins()
	if err != nil {
		return nil, err
	}

	return admins, nil
}

func (au *adminUseCase) CreateAdmin(name string, email string, password string) error {
	admin, err := admin.Create(name, email, password)
	if err != nil {
		return err
	}
	err = au.adminRepository.InsertAdmin(admin)
	if err != nil {
		return err
	}
	return nil
}

func (au *adminUseCase) UpdateAdmin(id string, name string, email string, password string) error {
	current_admin, err := au.adminRepository.GetAdmin(id)
	if err != nil {
		return err
	}

	adminId := current_admin.GetAdminId()

	update_admin, err := admin.New(adminId, name, email, password)
	if err != nil {
		return err
	}
	err = au.adminRepository.UpdateAdmin(update_admin)
	if err != nil {
		return err
	}
	return nil
}

func (au *adminUseCase) DeleteAdmin(id string) error {
	admin, err := au.adminRepository.GetAdmin(id)
	if err != nil {
		return err
	}

	err = au.adminRepository.DeleteAdmin(admin)
	if err != nil {
		return err
	}
	return nil
}