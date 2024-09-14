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
