package docs

// ROUTES

// swagger:route POST /users/ Users registerUser
// Register user.
// Registers a new user with the "movies:read" permission granted by default.
// A mail is also sent to the user containing an account activation token with a lifetime of 48 hours.
// All fields in the request body are required.
//
// Responses:
// 	201: registerUserResponse
//  422: validationError

// swagger:route PUT /users/activate Users activateUser
// Activate user.
// Activates a user, allowing them to be authenticated and make requests to the movies endpoints.
// All fields in the request body are required.
//
// Responses:
//	200: activateUserResponse
//	409: editConflictError
//  422: validationError

// swagger:route POST /users/refresh-activation-token Users refreshActivationToken
// Password reset token.
// Sends a mail to the user containing an updated activation token.
// Allowed only for unactivated users.
// All fields in the request body are required.
//
// Responses:
//	202: refreshActivationTokenResponse
//	401: invalidCredentialsError
//	403: alreadyActivateUserError
//  422: validationError

// swagger:route POST /users/authenticate Users authenticateUser
// Authenticate user.
// Returns a user-associated bearer token with a lifetime of 24 hours.
// All fields in the request body are required.
//
// Responses:
//	201: authenticateUserResponse
//	401: invalidCredentialsError
//	403: unactivatedUserError
//  422: validationError

// swagger:route POST /users/password-reset-token Users passwordResetToken
// Password reset token.
// Sends a mail to the user containing a password reset token with a lifetime of 15 minutes.
// All fields in the request body are required.
//
// Responses:
//	202: passwordResetTokenResponse
//	403: unactivatedUserError
//  422: validationError

// swagger:route PUT /users/update-password Users updatePassword
// Update password.
// Updates the user password.
// All fields in the request body are required.
//
// Responses:
//	200: updatePasswordResponse
//	409: editConflictError
//  422: validationError

// PARAMETERS
// swagger:parameters registerUser
type userRequest struct {
	// in: body
	Body struct {
		// required: true
		// example: johndoe
		Username string `json:"username"`

		// required: true
		// example: johndoe@mail.com
		Email string `json:"email"`

		// required: true
		// example: password
		Password string `json:"password"`
	}
}

// swagger:parameters activateUser
type activateUserRequest struct {
	// in: body
	Body struct {
		// required: true
		// example: OTBJEQX2EUIMKZHAMEMIPHE6TQ
		Token string `json:"token"`
	}
}

// swagger:parameters authenticateUser refreshActivationToken
type authenticateUserRequest struct {
	// in: body
	Body struct {
		// required: true
		// example: johndoe@mail.com
		Email string `json:"email"`

		// required: true
		// example: password
		Password string `json:"password"`
	}
}

// swagger:parameters updatePassword
type updatePasswordRequest struct {
	// in: body
	Body struct {
		// required: true
		// example: password
		Password string `json:"password"`

		// required: true
		// example: OTBJEQX2EUIMKZHAMEMIPHE6TQ
		Token string `json:"token"`
	}
}

// swagger:parameters passwordResetToken
type passwordResetTokenRequest struct {
	// in: body
	Body struct {
		// required: true
		// example: johndoe@mail.com
		Email string `json:"email"`
	}
}

// RESPONSES

// swagger:response registerUserResponse
type registerUserResponse struct {
	// in: body
	Body struct {
		userResponse

		// example: false
		Activated bool `json:"activated"`
	}
}

// swagger:response activateUserResponse
type activateUserResponse struct {
	// in: body
	Body struct {
		userResponse

		// example: 2
		Version int `json:"version"`
	}
}

// swagger:response authenticateUserResponse
type authenticateUserResponse struct {
	// in: body
	Body struct {
		tokenResponse
	}
}

// swagger:response passwordResetTokenResponse
type passwordResetTokenResponse struct {
	// in: body
	Body struct {
		// example: an email will be sent to you containing password reset instructions
		Message string `json:"message"`
	}
}

// swagger:response updatePasswordResponse
type updatePasswordResponse struct {
	// in: body
	Body struct {
		// example: your password was successfully reset
		Message string `json:"message"`
	}
}

// swagger:response refreshActivationTokenResponse
type refreshActivationTokenResponse struct {
	// in: body
	Body struct {
		// example: an email will be sent to you containing activation instructions
		Message string `json:"message"`
	}
}
