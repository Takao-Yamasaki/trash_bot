package controller

import (
	"strconv"
	"trash_bot/model"

	"github.com/gin-gonic/gin"
)

func IndexAdmin(c *gin.Context) {
	admins := model.GetAdmins()
	c.HTML(200, "index.html", gin.H{"admins": admins})
}

func DetailsAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	admin := model.GetAdmin(id)
	c.HTML(200, "detail.html", gin.H{"admin": admin})
}

func CreateAdmin(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	admin := model.Admin{Name: name, Email: email, Password: password}
	admin.Create()

	c.Redirect(301, "/admins")
}
