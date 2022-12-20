package vo

import (
	"fmt"
)

type UuId string

func NewUuid(value string) (*UuId, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:uuid NewUuId()")
		return nil, err
	}

	uuid := UuId(value)

	return &uuid, nil
}
