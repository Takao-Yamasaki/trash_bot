package controller

import (
	"fmt"
	"trash_bot/usecase"

	"github.com/gin-gonic/gin"
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

	type ResultDataField struct {
		CommentId string
		Contents  string
	}

	var data []ResultDataField
	for _, comment := range comments {
		commentId := comment.GetCommentId()
		contents := comment.GetContents()
		data = append(data, ResultDataField{CommentId: commentId, Contents: contents})
	}

	c.HTML(200, "comment/index.html", gin.H{"comments": data})
}

// 詳細の取得
func (cc *commentController) DetailComment(c *gin.Context) {
	id := c.Param("id")
	comment, err := cc.commentUseCase.GetComment(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	type ResultDataField struct {
		CommentId string
		Contents  string
	}

	data := ResultDataField{
		CommentId: comment.GetCommentId(),
		Contents:  comment.GetContents(),
	}
	c.HTML(200, "comment/detail.html", gin.H{"comment": data})
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
		ID       string `form:"id" binding:"required"`
		Contents string `form:"contents" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	id := form.ID

	contents := form.Contents

	err := cc.commentUseCase.UpdateComment(id, contents)
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

	id := form.ID
	err := cc.commentUseCase.DeleteComment(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/comment/index")
}
