package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
	"time"
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

	// create context for database operation with a 3-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.Db.QueryRowContext(ctx, stmt, movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres))
	return row.Scan(&movie.Id, &movie.Created, &movie.Version)
}

// Get returns the movie with the given ID from the database.
// A "record not found" error is returned if the ID doesn't belong to any movie.
func (m MovieController) Get(id int) (models.Movie, error) {
	stmt := `SELECT id, title, year, runtime, genres, created_at, version 
	FROM movies
	WHERE id = $1`

	// create context for database operation with a 3-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.Db.QueryRowContext(ctx, stmt, id)
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

// List fetches a list of movies from the database.
// The movies are fetched based on the query and filter parameters.
//
// title supports partial searching.
// The metadata for the query is also returned.
func (m MovieController) List(title string, genres []string, filters request.Filters) (models.Movies, response.Metadata, error) {
	// interpolate the sort column and direction into the SQL query
	// as keywords cannot be parameterized
	stmt := fmt.Sprintf(
		`SELECT count(*) OVER(), id, title, year, runtime, genres, created_at, version
	FROM movies
	WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
	AND (genres @> $2 OR $2 = '{}')
	ORDER BY %s %s, id ASC
	LIMIT $3 OFFSET $4`, filters.SortColumn(request.MovieFilterSortId), filters.SortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.Db.QueryContext(ctx, stmt, title, pq.Array(genres), filters.Limit, filters.Offset())
	if err != nil {
		return nil, response.Metadata{}, err
	}
	defer rows.Close()

	movies := models.Movies{}
	var totalRecords int

	for rows.Next() {
		movie := &models.Movie{}
		_ = rows.Scan(&totalRecords, &movie.Id, &movie.Title, &movie.Year, &movie.Runtime,
			pq.Array(&movie.Genres), &movie.Created, &movie.Version)

		movies = append(movies, *movie)
	}
	if err = rows.Err(); err != nil {
		return nil, response.Metadata{}, err
	}

	metadata := response.CalculateMetadata(filters.Page, filters.Limit, totalRecords)
	return movies, metadata, nil
}

// Update replaces the data of the movie in the database with those in the passed-in movie.
// An "edit conflict" error is returned if the version of the movie in the database does not
// match that in the parameter. This is done to prevent data races.
func (m MovieController) Update(movie *models.Movie) error {
	stmt := `UPDATE movies 
	SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1
	WHERE id = $5 AND version = $6
	RETURNING version`

	// create context for database operation with a 3-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.Db.QueryRowContext(ctx, stmt, movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres), movie.Id, movie.Version)
	err := row.Scan(&movie.Version)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrEditConflict
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

	// create context for database operation with a 3-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.Db.ExecContext(ctx, stmt, id)
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
