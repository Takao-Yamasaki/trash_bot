package controller

import (
	"strconv"
	"trash_bot/domain/model"
	"trash_bot/domain/repository"

	"github.com/gin-gonic/gin"
)

type adminController struct {
	adminRepository repository.AdminRepository
}

func NewAdminController(ar repository.AdminRepository) adminController {
	return adminController{
		adminRepository: ar,
	}
}

// 一覧の取得
func (ac *adminController) IndexAdmin(c *gin.Context) {
	admins := ac.adminRepository.GetAdmins()
	c.HTML(200, "admin/index.html", gin.H{"admins": admins})
}

// 詳細の取得
func (ac *adminController) DetailAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	admin := ac.adminRepository.GetAdmin(id)
	c.HTML(200, "admin/detail.html", gin.H{"admin": admin})
}

// 登録
func (ac *adminController) CreateAdmin(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	admin := model.Admin{Name: name, Email: email, Password: password}
	ac.adminRepository.Create(admin)

	c.Redirect(301, "/admin/index")
}

// 更新
func (ac *adminController) UpdateAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	admin := ac.adminRepository.GetAdmin(id)
	
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	admin.Name = name
	admin.Email = email
	admin.Password = password
	ac.adminRepository.Update(admin)

	c.Redirect(301, "/admin/index")
}

// 削除
func (ac *adminController) DeleteAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	admin := ac.adminRepository.GetAdmin(id)
	ac.adminRepository.Delete(admin)

	c.Redirect(301, "/admin/index")
}