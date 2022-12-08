package controller

import (
	"strconv"
	"trash_bot/model"

	"github.com/gin-gonic/gin"
)

// 一覧の取得
func IndexTrashDay(c *gin.Context) {
	trashdays := model.GetTrashDays()
	c.HTML(200, "trashday/index.html", gin.H{"trashdays": trashdays})
}

// 詳細の取得
func DetailsTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	trashday := model.GetTrashDay(id)
	c.HTML(200, "trashday/detail.html", gin.H{"trashday": trashday})
}

// 登録
func CreateTrashDay(c *gin.Context) {
	day := c.PostForm("day")
	trash := c.PostForm("trash")
	td := model.TrashDay{Day: day, Trash: trash}
	td.Create()

	c.Redirect(301, "/trash-days")
}

// 更新
func UpdateTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	trashday := model.GetTrashDay(id)
	day := c.PostForm("day")
	trash := c.PostForm("trash")
	trashday.Day = day
	trashday.Trash = trash
	trashday.Update()

	c.Redirect(301, "/trash-days")
}

// 削除
func DeleteTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	td := model.GetTrashDay(id)
	td.Delete()

	c.Redirect(301, "/trash-days")
}