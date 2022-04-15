package handlers

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/infrastructure/mock"
)

var RegisterMovieTestCases = map[string]struct {
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
