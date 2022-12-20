package vo

import (
	"fmt"
)

type Day string

func NewDay(value string) (*Day, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:day newDay()")
		return nil, err
	}

	day := Day(value)

	return &day, nil
}
