package controller

import (
	"trash_bot/model"
	
	"github.com/gin-gonic/gin"
)


func IndexAdmin(c *gin.Context) {
	admins := model.GetAdmins()
	c.HTML(200, "index.html", gin.H{"admins": admins})
}

func CreateAdmin(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	
	admin := model.Admin{Name: name, Email: email, Password: password}
	admin.Create()

	c.Redirect(301, "/admins")
}