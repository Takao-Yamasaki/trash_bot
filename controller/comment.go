package controller

import (
	"strconv"
	"trash_bot/domain/model"
	"trash_bot/domain/repository"

	"github.com/gin-gonic/gin"
)

type commentController struct{
	commentRepository repository.CommentRepository
}

func NewCommentController(cr repository.CommentRepository) commentController {
	return commentController{
		commentRepository: cr,
	}
}

// 一覧の取得
func (cc *commentController) IndexComment(c *gin.Context) {
	comments := cc.commentRepository.GetComments()
	c.HTML(200, "comment/index.html", gin.H{"comments": comments})
}

// 詳細の取得
func (cc *commentController) DetailsComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	comment := cc.commentRepository.GetComment(id)
	c.HTML(200, "comment/detail.html", gin.H{"comment": comment})
}

// 登録
func (cc *commentController) CreateComment(c *gin.Context) {
	contents := c.PostForm("contents")
	comment := model.Comment{Contents: contents}
	cc.commentRepository.Create(comment)

	c.Redirect(301, "/comment/index")
}

// 更新
func (cc *commentController) UpdateComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	comment := cc.commentRepository.GetComment(id)
	
	contents := c.PostForm("contents")
	comment.Contents = contents
	cc.commentRepository.Update(comment)

	c.Redirect(301, "/comment/index")
}

// 削除
func (cc *commentController) DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	comment := cc.commentRepository.GetComment(id)
	cc.commentRepository.Delete(comment)

	c.Redirect(301, "/comment/index")
}