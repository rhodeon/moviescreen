package handlers

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/infrastructure/mock"
	"time"
)

var registerUserTestCases = map[string]struct {
	RequestBody string
	WantCode    int
	WantBody    response.BaseResponse
}{
	"valid request": {
		RequestBody: `{
			"username": "person",
			"email": "person@mail.com",
			"password": "password"
		}`,
		WantCode: 201,
		WantBody: response.SuccessResponse(201, response.UserResponse{
			Id:        4,
			Username:  "person",
			Email:     "person@mail.com",
			Version:   1,
			Activated: false,
			Created:   mock.MockDate,
		}),
	},

	"missing username": {
		RequestBody: `{
			"email": "person@mail.com",
			"password": "password"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"username": "must be provided",
				},
			},
		),
	},

	"missing email": {
		RequestBody: `{
			"username": "person",
			"password": "password"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"email": "must be provided",
				},
			},
		),
	},

	"missing password": {
		RequestBody: `{
			"username": "person",
			"email": "person@mail.com"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"password": "must be provided",
				},
			},
		),
	},

	"duplicate username": {
		RequestBody: `{
			"username": "rhodeon",
			"email": "person@mail.com",
			"password": "password"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"username": "this username is already taken",
				},
			},
		),
	},

	"duplicate email": {
		RequestBody: `{
			"username": "aperson",
			"email": "rhodeon@dev.mail",
			"password": "password"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"email": "a user with this email address already exists",
				},
			},
		),
	},
}

var activateUserTestCases = map[string]struct {
	RequestBody string
	WantCode    int
	WantBody    response.BaseResponse
}{
	"valid token": {
		RequestBody: `{
 		   "token": "2QRJK3S54HAIUNIHNXEF4WSZSI"
		}`,
		WantCode: 200,
		WantBody: response.SuccessResponse(200, response.UserResponse{
			Id:        1,
			Username:  "rhodeon",
			Email:     "rhodeon@dev.mail",
			Activated: true,
			Version:   1,
			Created:   mock.MockDate,
		}),
	},

	"token too short": {
		RequestBody: `{
 		   "token": "2QRJK3S54HAIUWSZS"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"token": "must have exactly 26 characters",
				},
			},
		),
	},

	"token too long": {
		RequestBody: `{
 		   "token": "2QRJK3S54HAIUNIHNXEF4WSZSIUjkefjk"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"token": "must have exactly 26 characters",
				},
			},
		),
	},

	"invalid token": {
		RequestBody: `{
 		   "token": "2QRJK3S54HAKFDSHNXEF4WSZSI"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"token": "invalid or expired activation token",
				},
			},
		),
	},
}

var authenticateUserTestCases = map[string]struct {
	RequestBody string
	WantCode    int
	WantBody    response.BaseResponse
}{
	"valid request": {
		RequestBody: `{
			"email": "rhodeon@dev.mail",
			"password": "password"
		}`,
		WantCode: 201,
		WantBody: response.SuccessResponse(
			201,
			response.TokenResponse{
				PlainText: "token",
				Expires:   mock.AuthenticationBaseDate.Add(1 * 24 * time.Hour),
			},
		),
	},

	"missing email": {
		RequestBody: `{
			"password": "password"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"email": "must be provided",
				},
			},
		),
	},

	"missing password": {
		RequestBody: `{
			"email": "rhodeon@dev.mail"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"password": "must be provided",
				},
			},
		),
	},

	"non-existent email": {
		RequestBody: `{
			"email": "rhod@mail.com",
			"password": "password"
		}`,
		WantCode: 401,
		WantBody: response.ErrorResponse(
			401,
			response.GenericError("invalid user credentials"),
		),
	},

	"wrong password": {
		RequestBody: `{
			"email": "rhodeon@dev.mail",
			"password": "passwords"
		}`,
		WantCode: 401,
		WantBody: response.ErrorResponse(
			401,
			response.GenericError("invalid user credentials"),
		),
	},

	"unactivated user": {
		RequestBody: `{
			"email": "ruona@mail.com",
			"password": "password"
		}`,
		WantCode: 403,
		WantBody: response.ErrorResponse(
			403,
			response.GenericError("your account must be activated to access this resource"),
		),
	},
}

var createActivationTokenTestCases = map[string]struct {
	RequestBody string
	WantCode    int
	WantBody    response.BaseResponse
}{
	"valid request": {
		RequestBody: `{
			"email": "ruona@mail.com",
			"password": "password"
		}`,
		WantCode: 202,
		WantBody: response.SuccessResponse(
			202,
			map[string]string{"message": "an email will be sent to you containing activation instructions"},
		),
	},

	"missing email": {
		RequestBody: `{
			"password": "password"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"email": "must be provided",
				},
			},
		),
	},

	"missing password": {
		RequestBody: `{
			"email": "ruona@mail.com"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"password": "must be provided",
				},
			},
		),
	},

	"non-existent email": {
		RequestBody: `{
			"email": "rhod@mail.com",
			"password": "password"
		}`,
		WantCode: 401,
		WantBody: response.ErrorResponse(
			401,
			response.GenericError("invalid user credentials"),
		),
	},

	"wrong password": {
		RequestBody: `{
			"email": "rhodeon@dev.mail",
			"password": "passwords"
		}`,
		WantCode: 401,
		WantBody: response.ErrorResponse(
			401,
			response.GenericError("invalid user credentials"),
		),
	},

	"already activated account": {
		RequestBody: `{
			"email": "rhodeon@dev.mail",
			"password": "password"
		}`,
		WantCode: 403,
		WantBody: response.ErrorResponse(
			403,
			response.GenericError("user has already been activated"),
		),
	},
}

var createPasswordResetTokenTestCases = map[string]struct {
	RequestBody string
	WantCode    int
	WantBody    response.BaseResponse
}{
	"valid email": {
		RequestBody: `{
			"email": "rhodeon@dev.mail"
		}`,
		WantCode: 202,
		WantBody: response.SuccessResponse(
			202,
			map[string]string{"message": "an email will be sent to you containing password reset instructions"},
		),
	},

	"missing email": {
		RequestBody: `{}`,
		WantCode:    422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"email": "must be provided",
				},
			},
		),
	},

	"unactivated account": {
		RequestBody: `{
			"email": "ruona@mail.com"
		}`,
		WantCode: 403,
		WantBody: response.ErrorResponse(
			403,
			response.GenericError("your account must be activated to access this resource"),
		),
	},

	"non-existent email": {
		RequestBody: `{
			"email": "ruona@dev.com"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"email": "no matching email address found",
				},
			},
		),
	},
}

var updatePasswordTestCases = map[string]struct {
	RequestBody string
	WantCode    int
	WantBody    response.BaseResponse
}{
	"valid request ": {
		RequestBody: `{
			"password": "newpassword",
			"token": "2QRJK3S54HAIUNIHNXEF4WSZSI"
		}`,
		WantCode: 200,
		WantBody: response.SuccessResponse(
			200,
			map[string]string{"message": "your password was successfully reset"},
		),
	},

	"missing password": {
		RequestBody: `{
			"token": "2QRJK3S54HAIUNIHNXEF4WSZSI"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"password": "must be provided",
				},
			},
		),
	},

	"missing token": {
		RequestBody: `{
			"password": "newpassword"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"token": "must be provided",
				},
			},
		),
	},

	"invalid token": {
		RequestBody: `{
			"password": "newpassword",
			"token": "JW65OCFMDV2SLQXSHY4AXIDLUY"
		}`,
		WantCode: 422,
		WantBody: response.ErrorResponse(
			422,
			response.Error{
				Type: "user",
				Data: map[string]string{
					"token": "invalid or expired reset token",
				},
			},
		),
	},
}
