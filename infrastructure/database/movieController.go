package database

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
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

// Get returns the movie with the given ID from the database.
// A "record not found" error is returned if the ID doesn't belong to any movie.
func (m MovieController) Get(id int) (models.Movie, error) {
	stmt := `SELECT id, title, year, runtime, genres, created_at, version 
	FROM movies
	WHERE id = $1`

	row := m.Db.QueryRow(stmt, id)
	movie := models.Movie{}
	err := row.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Runtime,
		pq.Array(&movie.Genres), &movie.Created, &movie.Version)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Movie{}, repository.ErrRecordNotFound
		} else {
			return models.Movie{}, err
		}
	}

	return movie, nil
}

// Update replaces the data of the movie in the database with those in the passed-in movie.
func (m MovieController) Update(id int, movie *models.Movie) error {
	stmt := `UPDATE movies 
	SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1
	WHERE id = $5
	RETURNING id, version`

	row := m.Db.QueryRow(stmt, movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres), id)
	err := row.Scan(&movie.Id, &movie.Version)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrRecordNotFound
		} else {
			return err
		}
	}
	return nil
}

// Delete removes the movie with the given id from the database.
// An error is returned if no movie with the id is found.
func (m MovieController) Delete(id int) error {
	stmt := `DELETE FROM movies 
	WHERE id = $1`

	result, err := m.Db.Exec(stmt, id)
	if err != nil {
		return err
	}

	// check if no row was deleted and return a "record not found" error
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return repository.ErrRecordNotFound
	}

	return nil
}
