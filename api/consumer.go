package api

import (
	"encoding/json"
	"mock-api-serverless/mock"
	"mock-api-serverless/response"
	"mock-api-serverless/utils"
	"net/http"
	"strconv"
)

type ConsumerRes struct {
	Total int                 `json:"total"`
	List  []response.Consumer `json:"list"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_num := 10 // 默认10条数据
	query := r.URL.Query()
	numStr := query.Get("num")
	if numStr != "" {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			res := response.CommonRes[string]{
				Code: utils.Code().InvalidParams,
				Msg:  "invalid num",
			}
			json.NewEncoder(w).Encode(res)
			return
		} else {
			_num = num
		}
	}

	res := response.CommonRes[ConsumerRes]{
		Code: 0,
		Msg:  "success",
		Data: ConsumerRes{
			Total: _num,
			List:  mock.Consumer(_num),
		},
	}
	json.NewEncoder(w).Encode(res)
}
