package controller

import (
	"fmt"
	"strconv"
	"trash_bot/usecase"

	"github.com/gin-gonic/gin"
)

type trashDayController struct {
	trashDayUseCase usecase.TrashDayUseCase
}

func NewTrashDayController(tu usecase.TrashDayUseCase) trashDayController {
	return trashDayController{
		trashDayUseCase: tu,
	}
}

// 一覧の取得
func (tc *trashDayController) IndexTrashDay(c *gin.Context) {
	tds, err := tc.trashDayUseCase.GetTrashDays()
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
	td, err := tc.trashDayUseCase.GetTrashDay(id)
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
		return
	}

	day := form.Day
	trash := form.Trash

	err := tc.trashDayUseCase.CreateTrashDay(day, trash)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

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

	err = tc.trashDayUseCase.UpdateTrashDay(id, day, trash)
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
		ID string `form:"id" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(form.ID)
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	err = tc.trashDayUseCase.DeleteTrashDay(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/trash-day/index")
}