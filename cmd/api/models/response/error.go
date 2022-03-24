package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	ErrMessage404 = "resource not found"
	ErrMessage405 = "method not allowed"
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

// BadRequestError returns a 400 error with the specific reason.
// Usually due to an error in the JSON to response model conversion.
func BadRequestError(err error) *Error {
	response := NewErrorResponse(400, "")

	// possible JSON errors
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	var invalidUnmarshalError *json.InvalidUnmarshalError
	unknownFieldPrefix := "json: unknown field "
	bodyTooLargeError := "http: request body too large"

	switch {
	case errors.Is(err, io.EOF):
		// due to an empty JSON request body
		response.Message = "empty JSON request"

	case errors.Is(err, io.ErrUnexpectedEOF):
		// due to a malformed JSON request
		response.Message = "body contains malformed JSON"

	case errors.As(err, &syntaxError):
		// due to a malformed JSON request
		response.Message = fmt.Sprintf("body contains malformed JSON (at character %d)", syntaxError.Offset)

	case errors.As(err, &unmarshalTypeError):
		// due to incompatible JSON-to-Go mapping
		if unmarshalTypeError.Field != "" {
			response.Message = fmt.Sprintf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
		}
		response.Message = fmt.Sprintf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

	case errors.As(err, &invalidUnmarshalError):
		// due to invalid target.
		// should never happen.
		panic(err)

	case err.Error() == bodyTooLargeError:
		// due to JSON size larger than max limit
		response.Message = "JSON request body cannot be greater than 1MB"

	case strings.HasPrefix(err.Error(), unknownFieldPrefix):
		// due to unknown field in JSON request
		field := strings.TrimPrefix(err.Error(), unknownFieldPrefix)
		response.Message = fmt.Sprintf("unknown field: %s", field)

	default:
		response.Message = fmt.Sprintf("bad request")
	}

	return response
}
