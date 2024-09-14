package utils

type StatusCode struct {
	InvalidParams int
}

func Code() StatusCode {
	return StatusCode{
		InvalidParams: 1001,
	}
}
