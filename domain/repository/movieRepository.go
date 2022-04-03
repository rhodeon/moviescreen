package repository

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/domain/models"
)

type MovieRepository interface {
	Create(movie *models.Movie) error
	Get(id int) (models.Movie, error)
	List(title string, genres []string, filters request.Filters) (models.Movies, response.Metadata, error)
	Update(movie *models.Movie) error
	Delete(id int) error
}
