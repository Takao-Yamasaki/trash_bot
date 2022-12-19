package usecase

import (
	"trash_bot/domain/model/comment"
	"trash_bot/domain/repository"
)

type CommentUseCase interface {
	GetComment(id string ) (result *comment.Comment, err error)
	GetComments() (result []*comment.Comment, err error)
	CreateComment(contents string) error
	UpdateComment(id string, contents string) error
	DeleteComment(id string) error
}

type commentUseCase struct {
	commentRepository repository.CommentRepository
}

func NewCommentUseCase(cr repository.CommentRepository) CommentUseCase {
	return &commentUseCase{
		commentRepository: cr,
	}
}

// 詳細の取得
func (cu *commentUseCase) GetComment(id string) (result *comment.Comment, err error) {
	comment, err := cu.commentRepository.GetComment(id)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// 一覧の取得
func (cu *commentUseCase) GetComments() (result []*comment.Comment, err error){
	comments, err := cu.commentRepository.GetComments()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// 登録
func (cu *commentUseCase) CreateComment(contents string) error {
	comment, err := comment.Create(contents)
	if err != nil {
		return err
	}

	err = cu.commentRepository.InsertComment(comment)
	if err != nil {
		return err
	}

	return nil
}

// 更新
func (cu *commentUseCase) UpdateComment(id string, contents string) error {
	current_comment, err := cu.commentRepository.GetComment(id)
	if err != nil {
		return err
	}

	commentId := current_comment.GetCommentId() 
	
	update_comment, err := comment.New(commentId, contents)
	if err != nil {
		return err
	}

	err = cu.commentRepository.UpdateComment(update_comment)
	if err != nil {
		return err
	}

	return nil
}

// 削除
func (cu *commentUseCase) DeleteComment(id string) error {
	comment, err := cu.commentRepository.GetComment(id)
	if err != nil {
		return err
	}

	err = cu.commentRepository.DeleteComment(comment)
	if err != nil {
		return err
	}

	return nil
}