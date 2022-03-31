package handlers

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"net/http"
)

var createMovieTestCases = map[string]struct {
	requestBody string
	wantCode    int
	wantBody    response.BaseResponse
	wantHeaders http.Header
}{
	"valid request": {
		requestBody: `{
			"Title":   "The Shawshank Redemption",
			"Year":    1994,
			"Runtime": "142 mins",
			"Genres":  ["Drama"]
		}`,
		wantCode: 201,
		wantBody: response.SuccessResponse(201, response.MovieResponse{
			Id:      3,
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{"Drama"},
			Version: 1,
		}),
		wantHeaders: map[string][]string{
			"Location": {"/v1/movies/3"},
		},
	},

	"missing required fields": {
		requestBody: `{}`,
		wantCode:    422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"genres":  "must be provided",
				"runtime": "must be provided",
				"title":   "must be provided",
				"year":    "must be provided",
			},
		}),
	},

	"title with over 500 characters": {
		requestBody: `{
			"Title":   "ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss",
			"Year":    1994,
			"Runtime": "142 mins",
			"Genres":  ["Drama"]
		}`,
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"title": "must not have more than 500 characters",
			},
		}),
	},

	"year before 1888": {
		requestBody: `{
			"Title":   "The Shawshank Redemption",
			"Year":    1666,
			"Runtime": "142 mins",
			"Genres":  ["Drama"]
		}`,
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"year": "must not be before 1888",
			},
		}),
	},

	"year in the future": {
		requestBody: `{
			"Title":   "The Shawshank Redemption",
			"Year":    3000,
			"Runtime": "142 mins",
			"Genres":  ["Drama"]
		}`,
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"year": "must not be in the future",
			},
		}),
	},

	"negative runtime": {
		requestBody: `{
			"Title":   "The Shawshank Redemption",
			"Year":    1994,
			"Runtime": "-142 mins",
			"Genres":  ["Drama"]
		}`,
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"runtime": "must be a positive integer",
			},
		}),
	},

	"empty genres": {
		requestBody: `{
			"Title":   "The Shawshank Redemption",
			"Year":    1994,
			"Runtime": "142 mins",
			"Genres":  []
		}`,
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"genres": "must have at least 1 genre",
			},
		}),
	},

	"blank genre": {
		requestBody: `{
			"Title":   "The Shawshank Redemption",
			"Year":    1994,
			"Runtime": "142 mins",
			"Genres":  ["Drama", ""]
		}`,
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"genres": "must not have any blank genres",
			},
		}),
	},

	"duplicate genres": {
		requestBody: `{
			"Title":   "The Shawshank Redemption",
			"Year":    1994,
			"Runtime": "142 mins",
			"Genres":  ["Drama", "Drama"]
		}`,
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"genres": "must have unique genres",
			},
		}),
	},
}

var getMovieByIdTestCases = map[string]struct {
	requestId string
	wantCode  int
	wantBody  response.BaseResponse
}{
	"valid request": {
		requestId: "1",
		wantCode:  200,
		wantBody: response.SuccessResponse(
			200,
			response.MovieResponse{
				Id:      1,
				Title:   "Bullet Train",
				Year:    2022,
				Runtime: 108,
				Genres:  []string{"Action", "Comedy"},
				Version: 1,
			},
		),
	},

	"non-integer id": {
		requestId: "one",
		wantCode:  404,
		wantBody: response.ErrorResponse(
			404,
			response.Error{
				Type: "generic",
				Data: map[string]string{
					"message": response.ErrMessage404,
				},
			},
		),
	},

	"non-existent id": {
		requestId: "99",
		wantCode:  404,
		wantBody: response.ErrorResponse(
			404,
			response.Error{
				Type: "generic",
				Data: map[string]string{
					"message": response.ErrMessage404,
				},
			},
		),
	},
}

var updateMovieTestCases = map[string]struct {
	requestId   string
	requestBody string
	wantCode    int
	wantBody    response.BaseResponse
}{
	"valid request": {
		requestId: "1",
		requestBody: `{
			"Title":   "In The Heights",
			"Year":    2021,
			"Runtime": "110 mins",
			"Genres":  ["musical", "comedy"]
		}`,
		wantCode: 200,
		wantBody: response.SuccessResponse(
			200,
			response.MovieResponse{
				Id:      1,
				Title:   "In The Heights",
				Year:    2021,
				Runtime: 110,
				Genres:  []string{"musical", "comedy"},
				Version: 2,
			},
		),
	},

	"non-integer id": {
		requestId: "one",
		requestBody: `{
			"Title":   "In The Heights",
			"Year":    2021,
			"Runtime": "110 mins",
			"Genres":  ["musical", "comedy"]
		}`,
		wantCode: 404,
		wantBody: response.ErrorResponse(
			404,
			response.Error{
				Type: "generic",
				Data: map[string]string{
					"message": response.ErrMessage404,
				},
			},
		),
	},

	"non-existent id": {
		requestId: "99",
		requestBody: `{
			"Title":   "In The Heights",
			"Year":    2021,
			"Runtime": "110 min",
			"Genres":  ["musical", "comedy"]
		}`,
		wantCode: 404,
		wantBody: response.ErrorResponse(
			404,
			response.Error{
				Type: "generic",
				Data: map[string]string{
					"message": response.ErrMessage404,
				},
			},
		),
	},
}

var deleteMovieTestCases = map[string]struct {
	requestId string
	wantCode  int
	wantBody  response.BaseResponse
}{
	"valid request": {
		requestId: "1",
		wantCode:  200,
		wantBody: response.SuccessResponse(
			200,
			struct{}{},
		),
	},

	"non-integer id": {
		requestId: "one",
		wantCode:  404,
		wantBody: response.ErrorResponse(
			404,
			response.Error{
				Type: "generic",
				Data: map[string]string{
					"message": response.ErrMessage404,
				},
			},
		),
	},

	"non-existent id": {
		requestId: "99",
		wantCode:  404,
		wantBody: response.ErrorResponse(
			404,
			response.Error{
				Type: "generic",
				Data: map[string]string{
					"message": response.ErrMessage404,
				},
			},
		),
	},
}
