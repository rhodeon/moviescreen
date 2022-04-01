package repository

import "github.com/rhodeon/moviescreen/domain/models"

type MovieRepository interface {
	Create(movie *models.Movie) error
	Get(id int) (models.Movie, error)
	List() ([]models.Movie, error)
	Update(movie *models.Movie) error
	Delete(id int) error
}
