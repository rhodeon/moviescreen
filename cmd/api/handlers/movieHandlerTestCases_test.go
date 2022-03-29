package handlers

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"net/http"
	"time"
)

var createMovieTestCases = map[string]struct {
	request     request.MovieRequest
	wantCode    int
	wantBody    response.BaseResponse
	wantHeaders http.Header
}{
	"valid request": {
		request: request.MovieRequest{
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{"Drama"},
		},
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
		request:  request.MovieRequest{},
		wantCode: 422,
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
		request: request.MovieRequest{
			Title:   "ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{"Drama"},
		},
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"title": "must not have more than 500 characters",
			},
		}),
	},

	"year before 1888": {
		request: request.MovieRequest{
			Title:   "The Shawshank Redemption",
			Year:    1666,
			Runtime: 142,
			Genres:  []string{"Drama"},
		},
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"year": "must not be before 1888",
			},
		}),
	},

	"year in the future": {
		request: request.MovieRequest{
			Title:   "The Shawshank Redemption",
			Year:    time.Now().Year() + 5,
			Runtime: 142,
			Genres:  []string{"Drama"},
		},
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"year": "must not be in the future",
			},
		}),
	},

	"negative runtime": {
		request: request.MovieRequest{
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: -142,
			Genres:  []string{"Drama"},
		},
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"runtime": "must be a positive integer",
			},
		}),
	},

	"empty genres": {
		request: request.MovieRequest{
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{},
		},
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"genres": "must have at least 1 genre",
			},
		}),
	},

	"blank genre": {
		request: request.MovieRequest{
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{"Drama", ""},
		},
		wantCode: 422,
		wantBody: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"genres": "must not have any blank genres",
			},
		}),
	},

	"duplicate genres": {
		request: request.MovieRequest{
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{"Drama", "Drama"},
		},
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
	requestId int
	wantCode  int
	wantBody  response.BaseResponse
}{
	"valid request": {
		requestId: 1,
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

	"non-existing id": {
		requestId: 99,
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
