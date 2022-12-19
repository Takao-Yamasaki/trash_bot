package dto

import (
	"time"
	"trash_bot/domain/model/comment"
)

type Comment struct {
	ID        int
	CommentId string
	Contents  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func ConvertComment(comment *comment.Comment) *Comment {
	return &Comment{
		CommentId: comment.GetCommentId(),
		Contents:  comment.GetContents(),
	}
}

func AdaptComment(converted_comment *Comment) (*comment.Comment, error) {
	comment, err := comment.New(
		converted_comment.CommentId,
		converted_comment.Contents,
	)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func AdaptComments(converted_comments []*Comment) ([]*comment.Comment, error) {
	var comments []*comment.Comment

	for _, converted_comment := range converted_comments {
		comment, err := comment.New(
			converted_comment.CommentId,
			converted_comment.Contents,
		)

		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
