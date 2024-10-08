package utils

import (
	"strconv"

	"github.com/CiroLee/go-lorem"
)

func RandomEnum[T any](values []T) T {
	return lorem.Elements(values, 1)[0]
}

func TransferNumberQuery(query string, defaultVal int) (int, error) {
	if query != "" {
		num, err := strconv.Atoi(query)
		if err != nil {
			return 0, err
		} else {
			return num, nil
		}
	}
	return defaultVal, nil
}
