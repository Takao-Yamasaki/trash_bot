package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"trash_bot/usecase"
	"strconv"
)

type adminController struct {
	adminUseCase usecase.AdminUseCase
}

func NewAdminController(au usecase.AdminUseCase) adminController {
	return adminController{
		adminUseCase: au,
	}
}

// 一覧の取得
func (ac *adminController) IndexAdmin(c *gin.Context) {
	admins, err := ac.adminUseCase.GetAdmins()
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

	admin, err := ac.adminUseCase.GetAdmin(id)
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


	err := ac.adminUseCase.CreateAdmin(name, email, password)
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
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	name := form.Name
	email := form.Email
	password := form.Password

	err = ac.adminUseCase.UpdateAdmin(id, name, email, password)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H {"error": err.Error()})
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

	id, err := strconv.Atoi(form.ID)
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	err = ac.adminUseCase.DeleteAdmin(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/admin/index")
}