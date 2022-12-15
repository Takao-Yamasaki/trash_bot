package usecase

import (
	"trash_bot/domain/model"
	"trash_bot/domain/repository"
)

type CommentUseCase interface {
	GetComment(id int ) (result *model.Comment, err error)
	GetComments() (result []model.Comment, err error)
	CreateComment(contents string) error
	UpdateComment(id int, contents string) error
	DeleteComment(id int) error
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
func (cu *commentUseCase) GetComment(id int) (result *model.Comment, err error) {
	comment, err := cu.commentRepository.GetComment(id)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// 一覧の取得
func (cu *commentUseCase) GetComments() (result []model.Comment, err error){
	comments, err := cu.commentRepository.GetComments()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// 登録
func (cu *commentUseCase) CreateComment(contents string) error {
	comment := model.Comment{Contents: contents}
	err := cu.commentRepository.Create(comment)
	if err != nil {
		return err
	}

	return nil
}

// 更新
func (cu *commentUseCase) UpdateComment(id int, contents string) error {
	comment, err := cu.commentRepository.GetComment(id)
	if err != nil {
		return err
	}

	comment.Contents = contents
	err = cu.commentRepository.Update(*comment)
	if err != nil {
		return err
	}

	return nil
}

// 削除
func (cu *commentUseCase) DeleteComment(id int) error {
	comment, err := cu.commentRepository.GetComment(id)
	if err != nil {
		return err
	}

	err = cu.commentRepository.Delete(*comment)
	if err != nil {
		return err
	}

	return nil
}