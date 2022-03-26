package request

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/internal/validator"
	"github.com/rhodeon/moviescreen/internal/validator/rules"
	"time"
	"unicode/utf8"
)

type Movie struct {
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Runtime int      `json:"runtime"`
	Genres  []string `json:"genres"`
}

func (request *Movie) ToResponse(id int, version int) response.Movie {
	return response.Movie{
		Id:      id,
		Title:   request.Title,
		Year:    request.Year,
		Runtime: response.Runtime(request.Runtime),
		Genres:  request.Genres,
		Version: version,
	}
}

func (request *Movie) Validate() *validator.Validator {
	v := validator.New()
	const (
		fieldTitle   = "title"
		fieldYear    = "year"
		fieldRuntime = "runtime"
		fieldGenres  = "genres"
	)

	v.Check(request.Title != "", fieldTitle, "must be provided")
	v.Check(utf8.RuneCountInString(request.Title) <= 500, fieldTitle, "must not have more than 500 characters")

	v.Check(request.Year != 0, fieldYear, "must be provided")
	v.Check(request.Year >= 1888, fieldYear, "must not be before 1888")
	v.Check(request.Year <= time.Now().Year(), fieldYear, "must not be in the future")

	v.Check(request.Runtime != 0, fieldRuntime, "must be provided")
	v.Check(request.Runtime > 0, fieldRuntime, "must be a positive integer")

	v.Check(request.Genres != nil, fieldGenres, "must be provided")
	v.Check(len(request.Genres) >= 1, fieldGenres, "must have at least 1 genre")
	v.Check(rules.NotBlank(request.Genres), fieldGenres, "must not have any blank genres")
	v.Check(rules.Unique(request.Genres), fieldGenres, "must have unique genres")

	return v
}
