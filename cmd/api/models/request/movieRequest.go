package request

import (
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/internal/types"
	"github.com/rhodeon/moviescreen/internal/validator"
	"github.com/rhodeon/moviescreen/internal/validator/rules"
	"strings"
	"time"
	"unicode/utf8"
)

type MovieRequest struct {
	Title   *string        `json:"title"`
	Year    *int           `json:"year"`
	Runtime *types.Runtime `json:"runtime"`
	Genres  []string       `json:"genres"`
}

const (
	MovieFieldTitle   = "title"
	MovieFieldYear    = "year"
	MovieFieldRuntime = "runtime"
	MovieFieldGenres  = "genres"
)

// ToModel creates a movie model from a request with all fields being non-nil.
// An error occurs if a nil field is encountered.
// This should only be used when all fields are required in the validation.
func (request *MovieRequest) ToModel() models.Movie {
	return models.Movie{
		Title:   *request.Title,
		Year:    *request.Year,
		Runtime: *request.Runtime,
		Genres:  request.Genres,
	}
}

// UpdateModel maps the request to an already existing movie model,
// replacing with the non-nil request values.
func (request *MovieRequest) UpdateModel(model *models.Movie) {
	if request.Title != nil {
		model.Title = *request.Title
	}
	if request.Year != nil {
		model.Year = *request.Year
	}
	if request.Runtime != nil {
		model.Runtime = *request.Runtime
	}
	if request.Genres != nil {
		model.Genres = request.Genres
	}
}

func (request *MovieRequest) Validate(required []string) *validator.Validator {
	v := validator.New("movie")

	for _, field := range required {
		switch field {
		case MovieFieldTitle:
			v.Check(request.Title != nil, MovieFieldTitle, "must be provided")

		case MovieFieldYear:
			v.Check(request.Year != nil, MovieFieldYear, "must be provided")

		case MovieFieldRuntime:
			v.Check(request.Runtime != nil, MovieFieldRuntime, "must be provided")

		case MovieFieldGenres:
			v.Check(request.Genres != nil, MovieFieldGenres, "must be provided")
		}
	}

	if request.Title != nil {
		v.Check(strings.TrimSpace(*request.Title) != "", MovieFieldTitle, "must not be blank")
		v.Check(utf8.RuneCountInString(*request.Title) <= 500, MovieFieldTitle, "must not have more than 500 characters")
	}

	if request.Year != nil {
		v.Check(*request.Year >= 1888, MovieFieldYear, "must not be before 1888")
		v.Check(*request.Year <= time.Now().Year(), MovieFieldYear, "must not be in the future")
	}

	if request.Runtime != nil {
		v.Check(*request.Runtime > 0, MovieFieldRuntime, "must be a positive integer")
	}

	if request.Genres != nil {
		v.Check(len(request.Genres) >= 1, MovieFieldGenres, "must have at least 1 genre")
		v.Check(rules.NotBlank(request.Genres), MovieFieldGenres, "must not have any blank genres")
		v.Check(rules.Unique(request.Genres), MovieFieldGenres, "must have unique genres")
	}

	return v
}
