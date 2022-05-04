package models

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"time"
)

type Movie struct {
	Id      int
	Title   string
	Year    int
	Runtime int
	Genres  []string
	Version int
	Created time.Time
}

func (movie *Movie) ToResponse() response.MovieResponse {
	return response.MovieResponse{
		Id:    movie.Id,
		Title: movie.Title,
		Year:  movie.Year,

		// runtime is in minutes
		Runtime: movie.Runtime,

		Genres:  movie.Genres,
		Version: movie.Version,
	}
}

type Movies []Movie

func (movies Movies) ToResponse() []response.MovieResponse {
	moviesResponse := []response.MovieResponse{}
	for _, movie := range movies {
		moviesResponse = append(moviesResponse, movie.ToResponse())
	}
	return moviesResponse
}
