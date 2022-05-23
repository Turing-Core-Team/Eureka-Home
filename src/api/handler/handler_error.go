package handler

const (
	StatusCodeOk  ErrorCode = 200
	BadRequest    ErrorCode = 400
	InternalError ErrorCode = 500
)

type ErrorCode int

func (ec ErrorCode) Value() int {
	return int(ec)
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
