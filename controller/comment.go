package controller

import (
	"trash_bot/model"

	"github.com/gin-gonic/gin"
)

func IndexComment(c *gin.Context) {
	comments := model.GetTrashDays()
	c.HTML(200, "index.html", gin.H{"comments": comments})
}

func CreateComment(c *gin.Context) {
	contents := c.PostForm("contents")
	comment := model.Comment{Contents: contents}
	comment.Create()

	c.Redirect(301, "/comments")
}
