package comment

import (
	"fmt"
	"github.com/google/uuid"
)

type Comment struct {
	commentId commentId
	contents contents
}

type commentId string
type contents string

func New(commentId string, contents string) (*Comment, error) {
	createdCommentId, err := newCommentId(commentId)
	if err != nil {
		return nil, err
	}

	createdContents, err := newContents(contents)
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

// value constructors
func newCommentId(value string) (*commentId, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg: newCommentId()")
		return nil, err
	}
	
	commentId := commentId(value)

	return  &commentId, nil
}

func newContents(value string) (*contents, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg: newComments()")
		return nil, err
	}

	contents := contents(value)

	return &contents, nil
}