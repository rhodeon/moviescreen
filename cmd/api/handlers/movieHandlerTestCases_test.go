package handlers

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/cmd/api/responseErrors"
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
			"Runtime": 142,
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
			"Runtime": 142,
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
			"Runtime": 142,
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
			"Runtime": 142,
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
			"Runtime": -142,
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
			"Runtime": 142,
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
			"Runtime": 142,
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
			"Runtime": 142,
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
					"message": responseErrors.ErrMessageNotFound,
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
					"message": responseErrors.ErrMessageNotFound,
				},
			},
		),
	},
}

var listMoviesTestCases = map[string]struct {
	titleQuery    string
	genresQuery   []string
	filterQueries map[string]string
	wantCode      int
	wantBody      response.BaseResponse
}{
	"valid request (without queries)": {
		wantCode: 200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  1,
				PageLimit:    20,
				LastPage:     1,
				TotalRecords: 2,
			},
			Data: []response.MovieResponse{
				{
					Id:      1,
					Title:   "Bullet Train",
					Year:    2022,
					Runtime: 108,
					Genres:  []string{"Action", "Comedy"},
					Version: 1,
				},
				{
					Id:      2,
					Title:   "Hamilton",
					Year:    2020,
					Runtime: 140,
					Genres:  []string{"Musical", "Drama"},
					Version: 1,
				},
			},
		},
	},

	"valid request (with only title query)": {
		titleQuery: "Bullet",
		wantCode:   200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  1,
				PageLimit:    20,
				LastPage:     1,
				TotalRecords: 1,
			},
			Data: []response.MovieResponse{
				{
					Id:      1,
					Title:   "Bullet Train",
					Year:    2022,
					Runtime: 108,
					Genres:  []string{"Action", "Comedy"},
					Version: 1,
				},
			},
		},
	},

	"valid request (with only genres query)": {
		genresQuery: []string{"action", "comedy"},
		wantCode:    200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  1,
				PageLimit:    20,
				LastPage:     1,
				TotalRecords: 1,
			},
			Data: []response.MovieResponse{
				{
					Id:      1,
					Title:   "Bullet Train",
					Year:    2022,
					Runtime: 108,
					Genres:  []string{"Action", "Comedy"},
					Version: 1,
				},
			},
		},
	},

	"valid request (with page and limit)": {
		filterQueries: map[string]string{
			"page":  "2",
			"limit": "1",
		},
		wantCode: 200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  2,
				PageLimit:    1,
				LastPage:     2,
				TotalRecords: 2,
			},
			Data: []response.MovieResponse{
				{
					Id:      2,
					Title:   "Hamilton",
					Year:    2020,
					Runtime: 140,
					Genres:  []string{"Musical", "Drama"},
					Version: 1,
				},
			},
		},
	},

	"valid request (with sort by id - ascending)": {
		filterQueries: map[string]string{
			"sort": "id",
		},
		wantCode: 200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  1,
				PageLimit:    20,
				LastPage:     1,
				TotalRecords: 2,
			},
			Data: []response.MovieResponse{
				{
					Id:      1,
					Title:   "Bullet Train",
					Year:    2022,
					Runtime: 108,
					Genres:  []string{"Action", "Comedy"},
					Version: 1,
				},
				{
					Id:      2,
					Title:   "Hamilton",
					Year:    2020,
					Runtime: 140,
					Genres:  []string{"Musical", "Drama"},
					Version: 1,
				},
			},
		},
	},

	"valid request (with sort by id - descending)": {
		filterQueries: map[string]string{
			"sort": "-id",
		},
		wantCode: 200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  1,
				PageLimit:    20,
				LastPage:     1,
				TotalRecords: 2,
			},
			Data: []response.MovieResponse{
				{
					Id:      2,
					Title:   "Hamilton",
					Year:    2020,
					Runtime: 140,
					Genres:  []string{"Musical", "Drama"},
					Version: 1,
				},
				{
					Id:      1,
					Title:   "Bullet Train",
					Year:    2022,
					Runtime: 108,
					Genres:  []string{"Action", "Comedy"},
					Version: 1,
				},
			},
		},
	},

	"valid request (with sort by title - ascending)": {
		filterQueries: map[string]string{
			"sort": "title",
		},
		wantCode: 200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  1,
				PageLimit:    20,
				LastPage:     1,
				TotalRecords: 2,
			},
			Data: []response.MovieResponse{
				{
					Id:      1,
					Title:   "Bullet Train",
					Year:    2022,
					Runtime: 108,
					Genres:  []string{"Action", "Comedy"},
					Version: 1,
				},
				{
					Id:      2,
					Title:   "Hamilton",
					Year:    2020,
					Runtime: 140,
					Genres:  []string{"Musical", "Drama"},
					Version: 1,
				},
			},
		},
	},

	"valid request (with sort by title - descending)": {
		filterQueries: map[string]string{
			"sort": "-title",
		},
		wantCode: 200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  1,
				PageLimit:    20,
				LastPage:     1,
				TotalRecords: 2,
			},
			Data: []response.MovieResponse{
				{
					Id:      2,
					Title:   "Hamilton",
					Year:    2020,
					Runtime: 140,
					Genres:  []string{"Musical", "Drama"},
					Version: 1,
				},
				{
					Id:      1,
					Title:   "Bullet Train",
					Year:    2022,
					Runtime: 108,
					Genres:  []string{"Action", "Comedy"},
					Version: 1,
				},
			},
		},
	},

	"valid request (with sort by year - ascending)": {
		filterQueries: map[string]string{
			"sort": "year",
		},
		wantCode: 200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  1,
				PageLimit:    20,
				LastPage:     1,
				TotalRecords: 2,
			},
			Data: []response.MovieResponse{
				{
					Id:      2,
					Title:   "Hamilton",
					Year:    2020,
					Runtime: 140,
					Genres:  []string{"Musical", "Drama"},
					Version: 1,
				},
				{
					Id:      1,
					Title:   "Bullet Train",
					Year:    2022,
					Runtime: 108,
					Genres:  []string{"Action", "Comedy"},
					Version: 1,
				},
			},
		},
	},

	"valid request (with sort by year - descending)": {
		filterQueries: map[string]string{
			"sort": "-year",
		},
		wantCode: 200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  1,
				PageLimit:    20,
				LastPage:     1,
				TotalRecords: 2,
			},
			Data: []response.MovieResponse{
				{
					Id:      1,
					Title:   "Bullet Train",
					Year:    2022,
					Runtime: 108,
					Genres:  []string{"Action", "Comedy"},
					Version: 1,
				},
				{
					Id:      2,
					Title:   "Hamilton",
					Year:    2020,
					Runtime: 140,
					Genres:  []string{"Musical", "Drama"},
					Version: 1,
				},
			},
		},
	},

	"valid request (with sort by runtime - ascending)": {
		filterQueries: map[string]string{
			"sort": "runtime",
		},
		wantCode: 200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  1,
				PageLimit:    20,
				LastPage:     1,
				TotalRecords: 2,
			},
			Data: []response.MovieResponse{
				{
					Id:      1,
					Title:   "Bullet Train",
					Year:    2022,
					Runtime: 108,
					Genres:  []string{"Action", "Comedy"},
					Version: 1,
				},
				{
					Id:      2,
					Title:   "Hamilton",
					Year:    2020,
					Runtime: 140,
					Genres:  []string{"Musical", "Drama"},
					Version: 1,
				},
			},
		},
	},

	"valid request (with sort by runtime - descending)": {
		filterQueries: map[string]string{
			"sort": "-runtime",
		},
		wantCode: 200,
		wantBody: response.BaseResponse{
			Success: true,
			Status:  200,
			Metadata: &response.Metadata{
				CurrentPage:  1,
				PageLimit:    20,
				LastPage:     1,
				TotalRecords: 2,
			},
			Data: []response.MovieResponse{
				{
					Id:      2,
					Title:   "Hamilton",
					Year:    2020,
					Runtime: 140,
					Genres:  []string{"Musical", "Drama"},
					Version: 1,
				},
				{
					Id:      1,
					Title:   "Bullet Train",
					Year:    2022,
					Runtime: 108,
					Genres:  []string{"Action", "Comedy"},
					Version: 1,
				},
			},
		},
	},

	"zero page": {
		filterQueries: map[string]string{
			"page": "0",
		},
		wantCode: 422,
		wantBody: response.BaseResponse{
			Success: false,
			Status:  422,
			Error: &response.Error{
				Type: "filter",
				Data: map[string]string{
					"page": "must be greater than zero",
				},
			},
		},
	},

	"page exceeds 10 million": {
		filterQueries: map[string]string{
			"page": "20000000",
		},
		wantCode: 422,
		wantBody: response.BaseResponse{
			Success: false,
			Status:  422,
			Error: &response.Error{
				Type: "filter",
				Data: map[string]string{
					"page": "must be a maximum of 10 million",
				},
			},
		},
	},

	"zero limit": {
		filterQueries: map[string]string{
			"limit": "0",
		},
		wantCode: 422,
		wantBody: response.BaseResponse{
			Success: false,
			Status:  422,
			Error: &response.Error{
				Type: "filter",
				Data: map[string]string{
					"limit": "must be greater than zero",
				},
			},
		},
	},

	"invalid sort": {
		filterQueries: map[string]string{
			"sort": "actor",
		},
		wantCode: 422,
		wantBody: response.BaseResponse{
			Success: false,
			Status:  422,
			Error: &response.Error{
				Type: "filter",
				Data: map[string]string{
					"sort": "invalid sort value",
				},
			},
		},
	},

	"limit exceeds 100": {
		filterQueries: map[string]string{
			"limit": "200",
		},
		wantCode: 422,
		wantBody: response.BaseResponse{
			Success: false,
			Status:  422,
			Error: &response.Error{
				Type: "filter",
				Data: map[string]string{
					"limit": "must be a maximum of 100",
				},
			},
		},
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
			"Runtime": 110,
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

	"partial update (only title)": {
		requestId: "1",
		requestBody: `{
			"Title":   "In The Heights"
		}`,
		wantCode: 200,
		wantBody: response.SuccessResponse(
			200,
			response.MovieResponse{
				Id:      1,
				Title:   "In The Heights",
				Year:    2022,
				Runtime: 108,
				Genres:  []string{"Action", "Comedy"},
				Version: 2,
			},
		),
	},

	"partial update (only year)": {
		requestId: "1",
		requestBody: `{
			"Year":    1988
		}`,
		wantCode: 200,
		wantBody: response.SuccessResponse(
			200,
			response.MovieResponse{
				Id:      1,
				Title:   "Bullet Train",
				Year:    1988,
				Runtime: 108,
				Genres:  []string{"Action", "Comedy"},
				Version: 2,
			},
		),
	},

	"partial update (only runtime)": {
		requestId: "1",
		requestBody: `{
			"Runtime": 200
		}`,
		wantCode: 200,
		wantBody: response.SuccessResponse(
			200,
			response.MovieResponse{
				Id:      1,
				Title:   "Bullet Train",
				Year:    2022,
				Runtime: 200,
				Genres:  []string{"Action", "Comedy"},
				Version: 2,
			},
		),
	},

	"partial update (only genres)": {
		requestId: "1",
		requestBody: `{
			"Genres":  ["musical", "comedy"]
		}`,
		wantCode: 200,
		wantBody: response.SuccessResponse(
			200,
			response.MovieResponse{
				Id:      1,
				Title:   "Bullet Train",
				Year:    2022,
				Runtime: 108,
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
					"message": responseErrors.ErrMessageNotFound,
				},
			},
		),
	},

	"non-existent id": {
		requestId: "99",
		requestBody: `{
			"Title":   "In The Heights",
			"Year":    2021,
			"Runtime": 110,
			"Genres":  ["musical", "comedy"]
		}`,
		wantCode: 404,
		wantBody: response.ErrorResponse(
			404,
			response.Error{
				Type: "generic",
				Data: map[string]string{
					"message": responseErrors.ErrMessageNotFound,
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
					"message": responseErrors.ErrMessageNotFound,
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
					"message": responseErrors.ErrMessageNotFound,
				},
			},
		),
	},
}
