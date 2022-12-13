package controller

import (
	"fmt"
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
	admins, err := ac.adminRepository.GetAdmins()
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(200, "admin/index.html", gin.H{"admins": admins})
}

// 詳細の取得
func (ac *adminController) DetailAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	admin, err := ac.adminRepository.GetAdmin(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(200, "admin/detail.html", gin.H{"admin": admin})
}

// 登録
func (ac *adminController) CreateAdmin(c *gin.Context) {
	type RequestDataField struct {
		Name string `form:"name" binding:"required"`
		Email string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	
	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	name := form.Name
	email := form.Email
	password := form.Password

	admin := model.Admin{Name: name, Email: email, Password: password}
	err := ac.adminRepository.Create(admin)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/admin/index")
}

// 更新
func (ac *adminController) UpdateAdmin(c *gin.Context) {
	type RequestDataField struct {
		ID string `form:"id" binding:"required"`
		Name string `form:"name" binding:"required"`
		Email string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	
	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(form.ID)
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	name := form.Name
	email := form.Email
	password := form.Password

	admin, err := ac.adminRepository.GetAdmin(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H {"error": err.Error()})
		return
	}
	
	admin.Name = name
	admin.Email = email
	admin.Password = password
	err = ac.adminRepository.Update(*admin)
	if err != nil {
		fmt.Println(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/admin/index")
}

// 削除
func (ac *adminController) DeleteAdmin(c *gin.Context) {
	type RequestDataField struct {
		ID string `form:"id" binding:"required"`
	}
	
	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}
	
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}
	
	admin, err := ac.adminRepository.GetAdmin(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	err = ac.adminRepository.Delete(*admin)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/admin/index")
}