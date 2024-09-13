package api

import (
	"encoding/json"
	"net/http"
)

func ConsumerApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := map[string]any{
		"code": 0,
		"msg":  "ok",
		"data": "hello world",
	}
	json.NewEncoder(w).Encode(res)
}
