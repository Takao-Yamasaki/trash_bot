package controller

import (
	"strconv"
	"trash_bot/model"

	"github.com/gin-gonic/gin"
)

// 一覧の取得
func Index(c *gin.Context) {
	tds := model.GetTrashDays()
	c.HTML(200, "index.html", gin.H{"tds": tds})
}

// 詳細の取得
func DetailsTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	td := model.GetTrashDay(id)
	c.HTML(200, "detail.html", gin.H{"td": td})
}

// 登録
func CreateTrashDay(c *gin.Context) {
	week := c.PostForm("week") 
	trash := c.PostForm("trash")	
	td := model.TrashDay{Week: week, Trash: trash}
	td.Create()

	c.Redirect(301, "/")
}

// 更新
func UpdateTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	td := model.GetTrashDay(id)
	week := c.PostForm("week")
	trash := c.PostForm("trash")
	td.Week = week
	td.Trash = trash
	td.Update()

	c.Redirect(301, "/")
}

// 削除
func DeleteTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	td := model.GetTrashDay(id)
	td.Delete()

	c.Redirect(301, "/")
}