package mock

import (
	"github.com/rhodeon/moviescreen/domain/models"
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

type MovieController struct {
}

func (m MovieController) Create(movie *models.Movie) error {
	movie.Id = 3
	movie.Version = 1
	movie.Created = time.Now()
	return nil
}

func (m MovieController) Get(movie models.Movie) error {
	//TODO implement me
	panic("implement me")
}

func (m MovieController) Update(movie models.Movie) error {
	//TODO implement me
	panic("implement me")
}

func (m MovieController) Delete(movie models.Movie) error {
	//TODO implement me
	panic("implement me")
}
