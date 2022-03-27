package handlers

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"time"
)

var createMovieTestCases = map[string]struct {
	request      request.Movie
	wantResponse response.BaseResponse
}{
	"valid request": {
		request: request.Movie{
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{"Drama"},
		},
		wantResponse: response.SuccessResponse(200, response.Movie{
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{"Drama"},
		}),
	},

	"missing required fields": {
		request: request.Movie{},
		wantResponse: response.ErrorResponse(422, response.Error{
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
		request: request.Movie{
			Title:   "ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{"Drama"},
		},
		wantResponse: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"title": "must not have more than 500 characters",
			},
		}),
	},

	"year before 1888": {
		request: request.Movie{
			Title:   "The Shawshank Redemption",
			Year:    1666,
			Runtime: 142,
			Genres:  []string{"Drama"},
		},
		wantResponse: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"year": "must not be before 1888",
			},
		}),
	},

	"year in the future": {
		request: request.Movie{
			Title:   "The Shawshank Redemption",
			Year:    time.Now().Year() + 5,
			Runtime: 142,
			Genres:  []string{"Drama"},
		},
		wantResponse: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"year": "must not be in the future",
			},
		}),
	},

	"negative runtime": {
		request: request.Movie{
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: -142,
			Genres:  []string{"Drama"},
		},
		wantResponse: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"runtime": "must be a positive integer",
			},
		}),
	},

	"empty genres": {
		request: request.Movie{
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{},
		},
		wantResponse: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"genres": "must have at least 1 genre",
			},
		}),
	},

	"blank genre": {
		request: request.Movie{
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{"Drama", ""},
		},
		wantResponse: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"genres": "must not have any blank genres",
			},
		}),
	},

	"duplicate genres": {
		request: request.Movie{
			Title:   "The Shawshank Redemption",
			Year:    1994,
			Runtime: 142,
			Genres:  []string{"Drama", "Drama"},
		},
		wantResponse: response.ErrorResponse(422, response.Error{
			Type: "movie",
			Data: map[string]string{
				"genres": "must have unique genres",
			},
		}),
	},
}
