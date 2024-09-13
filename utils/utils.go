package utils

import (
	"github.com/CiroLee/go-lorem"
)

func RandomEnum[T any](values []T) T {
	return lorem.Elements(values, 1)[0]
}
