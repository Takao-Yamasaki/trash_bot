package controller

import (
	"strconv"
	"trash_bot/model"

	"github.com/gin-gonic/gin"
)

type TrashDayController struct{}

// 一覧の取得
func (tc TrashDayController) IndexTrashDay(c *gin.Context) {
	tds := model.GetTrashDays()
	c.HTML(200, "trashday/index.html", gin.H{"tds": tds})
}

// 詳細の取得
func (tc TrashDayController) DetailsTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	td := model.GetTrashDay(id)
	c.HTML(200, "trashday/detail.html", gin.H{"td": td})
}

// 登録
func (tc TrashDayController) CreateTrashDay(c *gin.Context) {
	day := c.PostForm("day")
	trash := c.PostForm("trash")
	td := model.TrashDay{Day: day, Trash: trash}
	td.Create()

	c.Redirect(301, "/trash-day/index")
}

// 更新
func (tc TrashDayController) UpdateTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	td := model.GetTrashDay(id)
	day := c.PostForm("day")
	trash := c.PostForm("trash")
	td.Day = day
	td.Trash = trash
	td.Update()

	c.Redirect(301, "/trash-day/index")
}

// 削除
func (tc TrashDayController) DeleteTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	td := model.GetTrashDay(id)
	td.Delete()

	c.Redirect(301, "/trash-day/index")
}
