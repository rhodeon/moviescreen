package models

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/internal/types"
	"time"
)

type Movie struct {
	Id      int
	Title   string
	Year    int
	Runtime types.Runtime
	Genres  []string
	Version int
	Created time.Time
}

func (movie *Movie) ToResponse() response.MovieResponse {
	return response.MovieResponse{
		Id:      movie.Id,
		Title:   movie.Title,
		Year:    movie.Year,
		Runtime: movie.Runtime,
		Genres:  movie.Genres,
		Version: movie.Version,
	}
}
