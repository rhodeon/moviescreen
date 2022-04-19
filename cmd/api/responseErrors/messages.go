package responseErrors

const (
	ErrMessageNotFound              = "resource not found"
	ErrMessageNotAllowed            = "method not allowed"
	ErrMessageInternalServer        = "internal server error"
	ErrMessageEditConflict          = "unable to update the record due to an edit conflict, please try again"
	ErrMessageRateLimitExceeded     = "rate limit exceeded"
	ErrMessageInvalidCredentials    = "invalid user credentials"
	ErrMessageInvalidAuthToken      = "invalid or missing authentication token"
	ErrMessageUnauthenticatedAccess = "you must be authenticated to access this resource"
	ErrMessageUnactivatedAccess     = "your account must be activated to access this resource"
	ErrMessageNotPermitted          = "your account doesn't have the necessary permissions to access this resource"
)
