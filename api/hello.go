package api

import (
	"encoding/json"
	"mock-api-serverless/middleware"
	"net/http"
)

func ConsumerApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	middleware.Cors(w, r)
	// if r.Method != http.MethodGet {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	fmt.Fprintf(w, "Method not allowed")
	// 	return
	// }

	res := map[string]any{
		"code": 0,
		"msg":  "ok",
		"data": "hello world",
	}
	json.NewEncoder(w).Encode(res)
}
