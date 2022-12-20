package vo

import (
	"fmt"
)

type Name string

func NewName(value string) (*Name, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:name newName()")
		return nil, err
	}

	name := Name(value)

	return &name, nil
}
