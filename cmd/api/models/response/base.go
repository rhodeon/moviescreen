package response

import "net/http"

// BaseResponse represents the whole JSON response sent to the client.
// It can either be a SuccessResponse or an ErrorResponse.
type BaseResponse struct {
	// Success is true on a successful request, and false on an error
	Success bool `json:"success"`

	// Status is the status code of the response
	Status int `json:"status,omitempty"`

	// Data is the data of a success response.
	// It is mutually exclusive to the Error.
	Data Data `json:"data,omitempty"`

	// Error is the data of an error response.
	// It is mutually exclusive to the Data.
	Error Error `json:"error,omitempty"`
}

// ErrorResponse is a constructor for an error response.
func ErrorResponse(code int, error Error) BaseResponse {
	return BaseResponse{
		Success: false,
		Status:  code,
		Data:    nil,
		Error:   error,
	}
}

// SuccessResponse is a constructor for a success response.
func SuccessResponse() BaseResponse {
	return BaseResponse{
		Success: true,
		Status:  http.StatusOK,
		Data:    nil,
		Error:   Error{},
	}
}

type Data interface{}
