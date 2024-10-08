package mock

import (
	"mock-api-serverless/response"
	"mock-api-serverless/utils"
	"time"

	"github.com/CiroLee/go-lorem"
)

func Consumer(num int) []response.Consumer {
	data := []response.Consumer{}
	from := time.Date(2023, time.Month(1), 1, 0, 0, 0, 0, time.Now().Location())
	to := time.Date(2024, time.Month(1), 23, 59, 59, 0, 0, time.Now().Location())
	for i := 0; i < num; i++ {
		age, _ := lorem.Int(10, 60)
		temp := response.Consumer{
			Id:       lorem.NanoId(12),
			Name:     lorem.Name("en", true),
			Age:      age,
			Gender:   utils.RandomEnum([]int{0, 1}), // 0: male, 1: female
			Avatar:   lorem.Picsum(lorem.PicsumOption{Width: 200}),
			CreateAt: lorem.Timestamp(from, to, "sec"),
			UpdateAt: lorem.Timestamp(from, to, "sec"),
		}
		data = append(data, temp)
	}

	return data
}

func List[T any](num, min, max int, generator func() T) []T {
	var data []T
	for i := 0; i < num; i++ {
		value := generator()
		data = append(data, value)
	}
	return data
}

func GenerateInt(min, max int) func() int {
	return func() int {
		val, _ := lorem.Int(min, max)
		return val
	}
}

func GenerateFloat(min, max int, decimal int) func() float64 {
	return func() float64 {
		val, _ := lorem.Float64(float64(min), float64(max), uint(decimal))
		return val
	}
}
