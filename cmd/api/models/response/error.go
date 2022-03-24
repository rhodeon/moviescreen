package response

const (
	ErrMessage404 = "resource not found"
	ErrMessage405 = "page not found"
	ErrMessage500 = "internal server error"
)

type Error struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code int, message string) *Error {
	return &Error{Status: "error", Code: code, Message: message}
}
