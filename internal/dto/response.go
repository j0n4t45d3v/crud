package dto

type ResponseUser[T any] struct {
	Timestamp string `json:"timestamp"`
	Status    int16  `json:"status"`
	Data      T      `json:"data"`
}

type ResponseSucess struct {
	Timestamp string `json:"timestamp"`
	Status    int16  `json:"status"`
	Message   string `json:"message"`
}

type ResponseError struct {
	Timestamp string `json:"timestamp"`
	Status    int16  `json:"status"`
	Error     string `json:"error"`
}
