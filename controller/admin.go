package controller

import (
	"strconv"
	"trash_bot/model"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (ac AdminController) IndexAdmin(c *gin.Context) {
	admins := model.GetAdmins()
	c.HTML(200, "admin/index.html", gin.H{"admins": admins})
}

func (ac AdminController) DetailsAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	admin := model.GetAdmin(id)
	c.HTML(200, "admin/detail.html", gin.H{"admin": admin})
}

func (ac AdminController) CreateAdmin(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	admin := model.Admin{Name: name, Email: email, Password: password}
	admin.Create()

	c.Redirect(301, "/admin/index")
}

func (ac AdminController) UpdateAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	admin := model.GetAdmin(id)
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	admin.Name = name
	admin.Email = email
	admin.Password = password
	admin.Update()

	c.Redirect(301, "/admin/index")
}

func (ac AdminController) DeleteAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	admin := model.GetAdmin(id)
	admin.Delete()

	c.Redirect(301, "/admin/index")
}