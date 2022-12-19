package repository

import (
	"trash_bot/domain/model/comment"
)

type CommentRepository interface {
	GetComment(id string) (result *comment.Comment, err error)
	GetComments() (result []*comment.Comment, err error)
	InsertComment(td *comment.Comment) error
	UpdateComment(td *comment.Comment) error
	DeleteComment(td *comment.Comment) error
}