package controller

import (
	"fmt"
	"strconv"
	"trash_bot/domain/model"
	"trash_bot/domain/repository"

	"github.com/gin-gonic/gin"
)

type trashDayController struct {
	trashDayRepository repository.TrashDayRepository
}

func NewTrashDayController(tr repository.TrashDayRepository) trashDayController {
	return trashDayController{
		trashDayRepository: tr,
	}
}

// 一覧の取得
func (tc *trashDayController) IndexTrashDay(c *gin.Context) {
	tds, err := tc.trashDayRepository.GetTrashDays()
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(200, "trashday/index.html", gin.H{"tds": tds})
}

// 詳細の取得
func (tc *trashDayController) DetailTrashDay(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}
	td, err := tc.trashDayRepository.GetTrashDay(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(200, "trashday/detail.html", gin.H{"td": td})
}

// 登録
func (tc *trashDayController) CreateTrashDay(c *gin.Context) {
	type RequestDataField struct {
		Day string `form:"day" binding:"required"`
		Trash string `form:"trash" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
	}

	day := form.Day
	trash := form.Trash

	td := model.TrashDay{Day: day, Trash: trash}
	tc.trashDayRepository.Create(td)

	c.Redirect(301, "/trash-day/index")
}

// 更新
func (tc *trashDayController) UpdateTrashDay(c *gin.Context) {
	
	type RequestDataField struct {
		ID string `form:"id" binding:"required"`
		Day string	`form:"day" binding:"required"`
		Trash string `form:"trash" binding:"required"`
	}
	
	var form RequestDataField
	
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(form.ID)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	day := form.Day
	trash := form.Trash

	td, err := tc.trashDayRepository.GetTrashDay(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}
	
	td.Day = day
	td.Trash = trash
	err = tc.trashDayRepository.Update(*td)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/trash-day/index")
}

// 削除
func (tc *trashDayController) DeleteTrashDay(c *gin.Context) {
	type RequestDataField struct {
		ID string `form:"id" binging:"required"`
	}

	var form RequestDataField
	
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(form.ID)
	if err != nil {
		fmt.Println()
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	td, err := tc.trashDayRepository.GetTrashDay(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}
	
	err = tc.trashDayRepository.Delete(*td)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/trash-day/index")
}
