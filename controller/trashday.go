package controller

import (
	"strconv"
	"trash_bot/domain/model"
	"trash_bot/domain/repository"

	"github.com/gin-gonic/gin"
)

type trashDayController struct{
	trashDayRepository repository.TrashDayRepository
}

func NewTrashDayController(tr repository.TrashDayRepository) trashDayController {
	return trashDayController{
		trashDayRepository: tr,
	}
}

// 一覧の取得
func (tc *trashDayController) IndexTrashDay(c *gin.Context) {
	tds := tc.trashDayRepository.GetTrashDays()
	c.HTML(200, "trashday/index.html", gin.H{"tds": tds})
}

// 詳細の取得
func (tc *trashDayController) DetailTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	td := tc.trashDayRepository.GetTrashDay(id)
	c.HTML(200, "trashday/detail.html", gin.H{"td": td})
}

// 登録
func (tc *trashDayController) CreateTrashDay(c *gin.Context) {
	day := c.PostForm("day")
	trash := c.PostForm("trash")
	td := model.TrashDay{Day: day, Trash: trash}
	tc.trashDayRepository.Create(td)

	c.Redirect(301, "/trash-day/index")
}

// 更新
func (tc *trashDayController) UpdateTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	td := tc.trashDayRepository.GetTrashDay(id)
	
	day := c.PostForm("day")
	trash := c.PostForm("trash")
	td.Day = day
	td.Trash = trash
	tc.trashDayRepository.Update(td)

	c.Redirect(301, "/trash-day/index")
}

// 削除
func (tc *trashDayController) DeleteTrashDay(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	td := tc.trashDayRepository.GetTrashDay(id)
	tc.trashDayRepository.Delete(td)

	c.Redirect(301, "/trash-day/index")
}
