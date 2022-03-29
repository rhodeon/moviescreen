package database

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/rhodeon/moviescreen/domain/models"
)

type MovieController struct {
	Db *sql.DB
}

// Create inserts the existing values of the Movie pointer into the database,
// and updates the values of the pointer's id, creation time and version.
// An error is returned if the operation fails.
func (m MovieController) Create(movie *models.Movie) error {
	stmt := `INSERT INTO movies (title, year, runtime, genres)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, version`

	row := m.Db.QueryRow(stmt, movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres))
	return row.Scan(&movie.Id, &movie.Created, &movie.Version)
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
