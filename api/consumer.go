package api

import (
	"encoding/json"
	"mock-api-serverless/middleware"
	"mock-api-serverless/mock"
	"mock-api-serverless/response"
	"mock-api-serverless/utils"
	"net/http"
)

type ConsumerRes struct {
	Total int                 `json:"total"`
	List  []response.Consumer `json:"list"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	middleware.Cors(w, r)

	query := r.URL.Query()
	numStr := query.Get("num")
	num, err := utils.TransferNumberQuery(numStr, 10)

	if err != nil {
		res := response.CommonRes[string]{
			Code: utils.Code().InvalidParams,
			Msg:  "consumer:invalid num",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := response.CommonRes[ConsumerRes]{
		Code: 0,
		Msg:  "success",
		Data: ConsumerRes{
			Total: num,
			List:  mock.Consumer(num),
		},
	}
	json.NewEncoder(w).Encode(res)
}
