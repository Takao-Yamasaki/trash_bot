package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"trash_bot/usecase"
)

type adminController struct {
	adminUseCase usecase.AdminUseCase
}

func NewAdminController(au usecase.AdminUseCase) adminController {
	return adminController{
		adminUseCase: au,
	}

}

func (ad *adminController) Login(c *gin.Context) {
	c.HTML(200, "login_logout/login.html", gin.H{})
}

func (ad *adminController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.HTML(200, "login_logout/logout.html", gin.H{})
}

func (ad *adminController) AuthLogin(c *gin.Context) {
	type RequestDataField struct {
		Email	string	`form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	email := form.Email
	password := form.Password

	admin, err := ad.adminUseCase.GetAdminForAuth(email)
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.GetPassword()), []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		c.HTML(401, "login_logout/login.html", gin.H{})
		return
	}

	session := sessions.Default(c)
	session.Set("AdminId", admin.GetAdminId())
	session.Save()
	c.Redirect(301, "/admin/index")
}

// 一覧の取得
func (ac *adminController) IndexAdmin(c *gin.Context) {
	admins, err := ac.adminUseCase.GetAdmins()
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	type ResultDataField struct {
		AdminId string
		Name string
		Email string
		Password string
	}
	
		var data []ResultDataField
		for _, admin := range admins {
			adminId := admin.GetAdminId()
			name := admin.GetName()
			email := admin.GetEmail()
			password := admin.GetPassword()
			data = append(data, ResultDataField{AdminId: adminId, Name: name, Email: email, Password: password})
		}

	c.HTML(200, "admin/index.html", gin.H{"admins": data})
}

// 詳細の取得
func (ac *adminController) DetailAdmin(c *gin.Context) {
	id := c.Param("id")
	admin, err := ac.adminUseCase.GetAdmin(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	type ResultDataField struct {
		AdminId string
		Name string
		Email string
		Password string
	}

	data := ResultDataField{
		AdminId: admin.GetAdminId(),
		Name: admin.GetName(),
		Email: admin.GetEmail(),
		Password: admin.GetPassword(),
	}

	c.HTML(200, "admin/detail.html", gin.H{"admin": data})
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

	hash, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 12)
	name := form.Name
	email := form.Email
	password := string(hash)


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

	hash, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 12)
	id := form.ID
	name := form.Name
	email := form.Email
	password := string(hash)


	err := ac.adminUseCase.UpdateAdmin(id, name, email, password)
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

	id := form.ID

	err := ac.adminUseCase.DeleteAdmin(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/admin/index")
}