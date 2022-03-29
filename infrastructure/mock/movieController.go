package mock

import (
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
	"time"
)

var movies = []models.Movie{
	{
		Id:      1,
		Title:   "Bullet Train",
		Year:    2022,
		Runtime: 108,
		Genres:  []string{"Action", "Comedy"},
		Version: 1,
		Created: time.Now(),
	},

	{
		Id:      2,
		Title:   "Hamilton",
		Year:    2020,
		Runtime: 140,
		Genres:  []string{"Musical", "Drama"},
		Version: 1,
		Created: time.Now(),
	},
}

type MovieController struct{}

func (m MovieController) Create(movie *models.Movie) error {
	movie.Id = 3
	movie.Version = 1
	movie.Created = time.Now()
	return nil
}

func (m MovieController) Get(id int) (models.Movie, error) {
	for _, movie := range movies {
		if movie.Id == id {
			return movie, nil
		}
	}

	return models.Movie{}, repository.ErrRecordNotFound
}

func (m MovieController) Update(id int, movie *models.Movie) error {
	for _, mov := range movies {
		if mov.Id == id {
			movie.Version = mov.Version + 1
			movie.Id = id
			return nil
		}
	}
	return repository.ErrRecordNotFound
}

func (m MovieController) Delete(id int) error {
	for _, movie := range movies {
		if movie.Id == id {
			// delete nothing as mock data is not persistent
			return nil
		}
	}
	return repository.ErrRecordNotFound
}
