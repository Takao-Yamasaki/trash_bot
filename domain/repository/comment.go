package repository

import (
	"trash_bot/domain/model"
)

type CommentRepository interface {
	GetComment(id int) (result *model.Comment, err error)
	GetComments() (result []model.Comment, err error)
	Create(td model.Comment) error
	Update(td model.Comment) error
	Delete(td model.Comment) error
}