package docs

// swagger:response notFoundError
type notFoundError struct {
	// in: body
	Body struct {
		genericType

		// Required: true
		// Example: {"message": "resource not found"}
		Data map[string]string `json:"data"`
	}
}

// A BadRequestError is returned when the request body contains unparsable data.
// swagger:response badRequestError
type badRequestError struct {
	// in: body
	Body struct {
		genericType

		// Required: true
		// Example: {"message": "body contains malformed JSON"}
		Data map[string]string `json:"data"`
	}
}

// swagger:response editConflictError
type editConflictError struct {
	// in: body
	Body struct {
		genericType

		// Required: true
		// Example: {"message": "unable to update the record due to an edit conflict, please try again"}
		Data map[string]string `json:"data"`
	}
}

// An UnauthenticatedError is returned when the request is made by an anonymous user to an endpoint that requires a valid user.
// swagger:response unauthenticatedError
type unauthenticatedError struct {
	// in: body
	Body struct {
		genericType

		// Required: true
		// Example: {"message": "you must be authenticated to access this resource"}
		Data map[string]string `json:"data"`
	}
}

// An PermissionError is returned when the user has insufficient permissions to carry out the request.
// swagger:response permissionError
type permissionError struct {
	// in: body
	Body struct {
		genericType

		// Required: true
		// Example: {"message": "your account doesn't have the necessary permissions to access this resource"}
		Data map[string]string `json:"data"`
	}
}

// A ValidationError is returned when the required input fails validation.
// swagger:response validationError
type validationError struct {
	// in: body
	Body struct {
		// Request type with validation errors.
		// Required: true
		// Example: movie
		Type string `json:"type"`

		// Mapping of invalidated fields to their error messages.
		// Required: true
		// Example: {"title": "must not be blank", "runtime": "must be provided"}
		Data map[string]string `json:"data"`
	}
}

type genericType struct {
	// Required: true
	// Example: generic
	Type string `json:"type"`
}
