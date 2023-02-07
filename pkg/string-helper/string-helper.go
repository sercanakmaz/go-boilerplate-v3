package string_helper

import "errors"

var ErrIsNullOrEmpty = errors.New("ErrIsNullOrEmpty")

func IsNullOrEmpty(value string) bool {
	return len(value) <= 0
}
