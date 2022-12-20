package vo

import (
	"fmt"
)

type Trash string

func NewTrash(value string) (*Trash, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:trash newTrash()")
		return nil, err
	}

	trash := Trash(value)

	return &trash, nil
}
