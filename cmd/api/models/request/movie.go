package request

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"time"
)
import "github.com/go-playground/validator/v10"

type Movie struct {
	Title   string   `json:"title" binding:"required,max=500"`
	Year    int      `json:"year" binding:"required,gte=1888"`
	Runtime int      `json:"runtime" binding:"required,gt=0"`
	Genres  []string `json:"genres" binding:"required,min=1,unique"`
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

func (request *Movie) ValidationErrors(errs validator.ValidationErrors) map[string]string {
	errMessages := map[string]string{}

	for _, ve := range errs {
		switch ve.StructField() {
		case "Title":
			switch ve.Tag() {
			case tagRequired:
				errMessages["title"] = "required field"
			case tagMaximum:
				errMessages["title"] = "must have a maximum of 500 characters"
			}

		case "Year":
			switch ve.Tag() {
			case tagRequired:
				errMessages["year"] = "required field"
			case tagGreaterThanOrEqual:
				errMessages["year"] = "must not be before 1888"
			}

		case "Runtime":
			switch ve.Tag() {
			case tagRequired:
				errMessages["runtime"] = "required field"
			case tagGreaterThan:
				errMessages["runtime"] = "must be greater than 0"
			}

		case "Genres":
			switch ve.Tag() {
			case tagRequired:
				errMessages["genres"] = "required field"
			case tagMinimum:
				errMessages["genres"] = "must contain at least one item"
			case tagUnique:
				errMessages["genres"] = "must have unique items"
			}
		}
	}

	return errMessages
}

func (request *Movie) Validate(errMessages map[string]string) {
	if _, exists := errMessages["year"]; !exists {
		// validate that Year is not in the future
		if request.Year > time.Now().Year() {
			errMessages["year"] = "must not be in the future"
		}
	}
}
