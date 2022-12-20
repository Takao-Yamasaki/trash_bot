package comment

import (
	"trash_bot/domain/model/vo"

	"github.com/google/uuid"
)

type Comment struct {
	commentId vo.UuId
	contents vo.Contents
}

func New(commentId string, contents string) (*Comment, error) {
	createdCommentId, err := vo.NewUuid(commentId)
	if err != nil {
		return nil, err
	}

	createdContents, err := vo.NewContents(contents)
	if err != nil {
		return nil, err
	}

	comment := Comment{
		commentId: *createdCommentId,
		contents: *createdContents,
	}

	return &comment, nil
}

func Create(contents string) (*Comment, error) {
	commentId := uuid.New().String()
	comment, err := New(commentId, contents)

	if err != nil {
		return nil, err
	}

	return comment, err
}

// Getter
func (c Comment) GetCommentId() string {
	return string(c.commentId)
}

func (c Comment) GetContents() string {
	return string(c.contents)
}