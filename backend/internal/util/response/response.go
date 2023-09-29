package response

import "net/http"

var (
	// 2xx
	Success                 = NewMessage().SendResponse(http.StatusOK, "20000", "success", "success")
	SuccessCreateUserWallet = NewMessage().SendResponse(http.StatusCreated, "20100", "success created user wallet", "Success created user wallet.")

	// 4xx
	BadRequestError = NewMessage().SendResponse(http.StatusBadRequest, "40000", "Bad Request", "Oops! Your request couldn't be processed. Please check your input and try again.")

	IncorrectCredentialsError = NewMessage().SendResponse(http.StatusUnauthorized, "40100", "Incorrect credentials", "Authentication failed. Please double-check your credentials and try again.")
	TokenExpiredError         = NewMessage().SendResponse(http.StatusUnauthorized, "40101", "Session expired", "Session expired. Please log in again.")
	TokenSignatureInvalid     = NewMessage().SendResponse(http.StatusUnauthorized, "40102", "Token signature invalid", "Invalid session. Please log in again.")
	RefreshTokenInvalid       = NewMessage().SendResponse(http.StatusUnauthorized, "40103", "Refresh token invalid", "Refresh token invalid")
	InvalidTokenError         = NewMessage().SendResponse(http.StatusUnauthorized, "40104", "Invalid Token", "Invalid Token")

	ForbiddenError = NewMessage().SendResponse(http.StatusForbidden, "40300", "Forbidden", "Forbidden")

	RouteNotFoundError  = NewMessage().SendResponse(http.StatusNotFound, "40400", "Route not found", "Oops! The page you're looking for can't be found.")
	RecordNotFoundError = NewMessage().SendResponse(http.StatusNotFound, "40401", "Record not found", "Record not found.")
	UserNotFoundError   = NewMessage().SendResponse(http.StatusNotFound, "40402", "User not found", "User not registered, please register to continue.")

	ConflictError               = NewMessage().SendResponse(http.StatusConflict, "40900", "Conflict", "Conflict")
	UsernameOrEmailAlreadyExist = NewMessage().SendResponse(http.StatusConflict, "40900", "Username or email already taken", "Username or email is already used. Please choose another.")

	// 5xx
	InternalServerError = NewMessage().SendResponse(http.StatusInternalServerError, "50000", "Internal Server Error", "Internal Server Error. Please try again later.")
)

type response struct {
	Code          string      `json:"code"`
	Message       string      `json:"message"`
	ClientMessage string      `json:"client_message"`
	Data          interface{} `json:"data"`
}

type Message[T any] struct {
	HttpCode int
	Response T
}

func NewMessage() *Message[interface{}] {
	return &Message[interface{}]{}
}

func (m Message[T]) SendResponse(httpCode int, internalCode string, message string, clientMessage string) *Message[response] {
	resp := new(Message[response])
	resp.HttpCode = httpCode
	resp.Response = response{
		Code:          internalCode,
		Message:       message,
		ClientMessage: clientMessage,
	}
	return resp
}
