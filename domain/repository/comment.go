package repository

import (
	"trash_bot/domain/model"
)

type CommentRepository interface {
	GetComment(id int) model.Comment
	GetComments() []model.Comment
	Create(td model.Comment)
	Update(td model.Comment)
	Delete(td model.Comment)
}