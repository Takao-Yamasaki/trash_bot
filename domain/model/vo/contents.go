package vo

import (
	"fmt"
)

type Contents string

func NewContents(value string) (*Contents, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:contents newContents()")
		return nil, err
	}

	contents := Contents(value)

	return &contents, nil
}