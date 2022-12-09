package controller

import (
	"strconv"
	"trash_bot/model"

	"github.com/gin-gonic/gin"
)

func IndexComment(c *gin.Context) {
	comments := model.GetComments()
	c.HTML(200, "index.html", gin.H{"comments": comments})
}

func DetailsComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	comment := model.GetComment(id)
	c.HTML(200, "detail.html", gin.H{"comment": comment})
}

func CreateComment(c *gin.Context) {
	contents := c.PostForm("contents")
	comment := model.Comment{Contents: contents}
	comment.Create()

	c.Redirect(301, "/comments")
}

func UpdateComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	comment := model.GetComment(id)
	contents := c.PostForm("contents")
	comment.Contents = contents
	comment.Update()

	c.Redirect(301, "/comments")
}

func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	comment := model.GetComment(id)
	comment.Delete()

	c.Redirect(301, "/comments")
}