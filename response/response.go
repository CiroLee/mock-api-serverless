package response

type CommonRes[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type Consumer struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}
