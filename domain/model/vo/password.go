package vo

import (
	"fmt"
)

type Password string

func NewPassword(value string) (*Password, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:password newPassword()")
		return nil, err
	}

	password := Password(value)

	return &password, nil
}
