package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"trash_bot/usecase"
	"strconv"

)

type commentController struct {
	commentUseCase usecase.CommentUseCase
}

func NewCommentController(cu usecase.CommentUseCase) commentController {
	return commentController{
		commentUseCase: cu,
	}
}

// 一覧の取得
func (cc *commentController) IndexComment(c *gin.Context) {
	comments, err := cc.commentUseCase.GetComments()
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(200, "comment/index.html", gin.H{"comments": comments})
}

// 詳細の取得
func (cc *commentController) DetailComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	comment, err := cc.commentUseCase.GetComment(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(200, "comment/detail.html", gin.H{"comment": comment})
}

// 登録
func (cc *commentController) CreateComment(c *gin.Context) {
	type RequestDataField struct {
		Contents string `form:"contents" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}
	
	contents := form.Contents
	err := cc.commentUseCase.CreateComment(contents)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/comment/index")
}

// 更新
func (cc *commentController) UpdateComment(c *gin.Context) {
	
	type RequestDataField struct {
		ID string `form:"id" binding:"required"`
		Contents string `form:"contents" binding:"required"`
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

	contents := form.Contents

	err = cc.commentUseCase.UpdateComment(id,contents)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/comment/index")
}

// 削除
func (cc *commentController) DeleteComment(c *gin.Context) {
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

	err = cc.commentUseCase.DeleteComment(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/comment/index")
}