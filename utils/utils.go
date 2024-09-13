package utils

import (
	"github.com/CiroLee/go-lorem"
)

func RandomEnum[T any](values []T) T {
	l := len(values)
	index, _ := lorem.Int(0, l-1)
	return values[index]
}
