package persistance

import (
	"gorm.io/gorm"

	"trash_bot/domain/model"
	"trash_bot/domain/repository"
)

type commentPersistance struct {
	Conn *gorm.DB
}

func NewCommentPersistance(conn *gorm.DB, c repository.CommentRepository) *commentPersistance {
	return &commentPersistance{Conn: conn}
}

// 1件の取得
func (co *commentPersistance) GetComment(id int) model.Comment {
	var comment model.Comment
	co.Conn.First(&comment)

	return comment
}

// 一覧の取得
func (co *commentPersistance) GetComments() []model.Comment {
	var cms []model.Comment
	co.Conn.Find(&cms)
	return cms
}

// 登録
func (co *commentPersistance) Create(cm model.Comment) {
	co.Conn.Create(&cm)
}

func (co *commentPersistance) Update(cm model.Comment) {
	co.Conn.Save(&cm)
}

func (co *commentPersistance) Delete(cm model.Comment) {
	co.Conn.Delete(&cm)
}