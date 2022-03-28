package database

import (
	"database/sql"
	"github.com/rhodeon/moviescreen/domain/models"
)

type MovieController struct {
	Db *sql.DB
}

func (m MovieController) Create(movie models.Movie) error {
	//TODO implement me
	panic("implement me")
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
