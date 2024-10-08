package api

import (
	"encoding/json"
	"mock-api-serverless/middleware"
	"mock-api-serverless/mock"
	"mock-api-serverless/response"
	"mock-api-serverless/utils"
	"net/http"
)

type ListRes struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

func ListApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	middleware.Cors(w, r)
	query := r.URL.Query()
	numStr := query.Get("num")
	minStr := query.Get("min")
	maxStr := query.Get("max")
	decimalStr := query.Get("decimal")
	num, errNum := utils.TransferNumberQuery(numStr, 12)
	min, errMin := utils.TransferNumberQuery(minStr, 0)
	max, errMax := utils.TransferNumberQuery(maxStr, 2000)
	decimal, decimalErr := utils.TransferNumberQuery(decimalStr, 0)
	if errNum != nil {
		res := response.CommonRes[string]{
			Code: utils.Code().InvalidParams,
			Msg:  "list: invalid num",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	if errMin != nil {
		res := response.CommonRes[string]{
			Code: utils.Code().InvalidParams,
			Msg:  "list: invalid min",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	if errMax != nil {
		res := response.CommonRes[string]{
			Code: utils.Code().InvalidParams,
			Msg:  "list: invalid max",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if decimalErr != nil {
		res := response.CommonRes[string]{
			Code: utils.Code().InvalidParams,
			Msg:  "list: invalid decimal",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	var list any
	if decimal != 0 {
		list = mock.List(num, min, max, mock.GenerateFloat(min, max, decimal))
	} else {
		list = mock.List(num, max, min, mock.GenerateInt(min, max))
	}

	res := response.CommonRes[ListRes]{
		Code: 0,
		Msg:  "success",
		Data: ListRes{
			Total: num,
			List:  list,
		},
	}

	json.NewEncoder(w).Encode(res)

}
