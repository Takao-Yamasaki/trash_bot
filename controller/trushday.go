package controller

import (
	"github.com/gin-gonic/gin"
	"trash_bot/model"
)

func Index(c *gin.Context) {
	trashdays := model.GetTrashdays()
	c.HTML(200, "index.html", gin.H{"trashdays": trashdays})
}

func CreateTrashday(c *gin.Context) {
	dayofweek := c.PostForm("dayofweek") 
	typeoftrush := c.PostForm("typeoftrash")	
	
	trashday := model.Trashday{Dayofweek: dayofweek, Typeoftrash: typeoftrush}
	trashday.Create()
	
	c.Redirect(301, "/")
}