package repository

import "github.com/rhodeon/moviescreen/domain/models"

type MovieRepository interface {
	Create(movie models.Movie) error
	Get(movie models.Movie) error
	Update(movie models.Movie) error
	Delete(movie models.Movie) error
}
