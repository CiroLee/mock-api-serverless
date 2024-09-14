package api

import (
	"encoding/json"
	"mock-api-serverless/response"
	"mock-api-serverless/utils"
	"net/http"

	"github.com/CiroLee/go-lorem"
)

func generateData() []response.Consumer {
	data := []response.Consumer{}
	for i := 0; i < 10; i++ {
		age, _ := lorem.Int(10, 60)
		temp := response.Consumer{
			Id:     lorem.NanoId(12),
			Name:   lorem.Name("en", true),
			Age:    age,
			Gender: utils.RandomEnum[int]([]int{0, 1}), // 0: male, 1: female
			Avatar: lorem.Picsum(lorem.PicsumOption{Width: 200}),
		}
		data = append(data, temp)
	}

	return data
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := response.CommonRes[[]response.Consumer]{
		Code: 0,
		Msg:  "ok",
		Data: generateData(),
	}
	json.NewEncoder(w).Encode(res)
}
