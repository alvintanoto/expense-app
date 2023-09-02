package response

import "net/http"

var (
	// 2xx
	Success = NewMessage().SendResponse(http.StatusOK, "20000", "success")

	// 4xx
	BadRequestError = NewMessage().SendResponse(http.StatusBadRequest, "40000", "Bad Request")

	IncorrectCredentialsError = NewMessage().SendResponse(http.StatusUnauthorized, "40100", "Incorrect credentials")
	UnauthorizedError         = NewMessage().SendResponse(http.StatusUnauthorized, "40101", "Unauthorized")
	TokenExpiredError         = NewMessage().SendResponse(http.StatusUnauthorized, "40102", "Session expired")
	TokenSignatureInvalid     = NewMessage().SendResponse(http.StatusUnauthorized, "40103", "Token signature invalid")

	ForbiddenError = NewMessage().SendResponse(http.StatusForbidden, "40300", "Forbidden")

	RouteNotFoundError  = NewMessage().SendResponse(http.StatusNotFound, "40400", "Route not found")
	RecordNotFoundError = NewMessage().SendResponse(http.StatusNotFound, "40401", "Record not found")

	ConflictError = NewMessage().SendResponse(http.StatusConflict, "40900", "Conflict")

	UnprocessableError = NewMessage().SendResponse(http.StatusUnprocessableEntity, "42200", "Unprocessable Entity")

	// 5xx
	InternalServerError = NewMessage().SendResponse(http.StatusInternalServerError, "50000", "Internal Server Error")
)

type response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Message[T any] struct {
	HttpCode int
	Response T
}

func NewMessage() *Message[interface{}] {
	return &Message[interface{}]{}
}

func (m Message[T]) SendResponse(httpCode int, internalCode string, message string) *Message[response] {
	resp := new(Message[response])
	resp.HttpCode = httpCode
	resp.Response = response{
		Code:    internalCode,
		Message: message,
	}
	return resp
}
