package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rhodeon/moviescreen/internal/validator"
	"io"
	"net/http"
	"strings"
)

const (
	ErrMessage404          = "resource not found"
	ErrMessage405          = "method not allowed"
	ErrMessage500          = "internal server error"
	ErrMessageEditConflict = "unable to update the record due to an edit conflict, please try again"
)

// Error represents the data of an error in a response.
type Error struct {
	// Type of the request which incurs an error.
	// It is "generic" if the error is independent of the request fields,
	// as in the cases of 404 and 405 errors, amongst others.
	Type string `json:"type"`

	// Data of the error content.
	// For a "generic" Type, the Data consists only of a single "message" key.
	Data map[string]string `json:"data"`
}

// NewError is a constructor for a new error.
func NewError(errorType string, errors map[string]string) Error {
	return Error{
		Type: errorType,
		Data: errors,
	}
}

// GenericError returns an Error which a "generic" Type.
func GenericError(message string) Error {
	return Error{
		Type: "generic",
		Data: map[string]string{
			"message": message,
		},
	}
}

// BadRequestError returns a 400 error with the specific reason.
// Usually due to an error in the JSON to response model conversion.
func BadRequestError(err error) BaseResponse {
	var errorMessage string

	// possible JSON errors
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	var invalidUnmarshalError *json.InvalidUnmarshalError
	unknownFieldPrefix := "json: unknown field "
	bodyTooLargeError := "http: request body too large"

	switch {
	case errors.Is(err, io.EOF):
		// due to an empty JSON request body
		errorMessage = "empty JSON request"

	case errors.Is(err, io.ErrUnexpectedEOF):
		// due to a malformed JSON request
		errorMessage = "body contains malformed JSON"

	case errors.As(err, &syntaxError):
		// due to a malformed JSON request
		errorMessage = fmt.Sprintf("body contains malformed JSON (at character %d)", syntaxError.Offset)

	case errors.As(err, &unmarshalTypeError):
		// due to incompatible JSON-to-Go mapping
		if unmarshalTypeError.Field != "" {
			errorMessage = fmt.Sprintf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
		}
		errorMessage = fmt.Sprintf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

	case errors.As(err, &invalidUnmarshalError):
		// due to invalid target.
		// should never happen.
		panic(err)

	case err.Error() == bodyTooLargeError:
		// due to JSON size larger than max limit
		errorMessage = "JSON request body cannot be greater than 1MB"

	case strings.HasPrefix(err.Error(), unknownFieldPrefix):
		// due to unknown field in JSON request
		field := strings.TrimPrefix(err.Error(), unknownFieldPrefix)
		errorMessage = fmt.Sprintf("unknown field: %s", field)

	default:
		errorMessage = fmt.Sprintf(err.Error())
	}

	return ErrorResponse(http.StatusBadRequest, GenericError(errorMessage))
}

// UnprocessableEntityError returns a 422 error.
// It occurs due to a request which failed validation.
func UnprocessableEntityError(v *validator.Validator) BaseResponse {
	// extract the first error of each validation field
	// to display in the error response
	firstErrors := map[string]string{}
	for field, errs := range v.Errors {
		firstErrors[field] = errs[0]
	}

	return ErrorResponse(
		http.StatusUnprocessableEntity,
		NewError(v.Type, firstErrors),
	)
}
