package database

import (
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
)

var createMovieTestCases = map[string]struct {
	movie        models.Movie
	wantNewMovie models.Movie
	wantErr      error
}{
	"valid movie": {
		movie: models.Movie{
			Title:   "The Godfather",
			Year:    1972,
			Runtime: 175,
			Genres:  []string{"Crime", "Drama"},
		},
		wantNewMovie: models.Movie{
			Id:      4,
			Title:   "The Godfather",
			Year:    1972,
			Runtime: 175,
			Genres:  []string{"Crime", "Drama"},
			Version: 1,
		},
		wantErr: nil,
	},
}

var getMovieTestCases = map[string]struct {
	id        int
	wantMovie models.Movie
	wantErr   error
}{
	"valid id": {
		id: 2,
		wantMovie: models.Movie{
			Id:      2,
			Title:   "Hamilton",
			Year:    2020,
			Runtime: 140,
			Genres:  []string{"Musical", "Drama"},
			Version: 1,
		},
		wantErr: nil,
	},

	"non-existent id": {
		id:        99,
		wantMovie: models.Movie{},
		wantErr:   repository.ErrRecordNotFound,
	},
}

var updateMovieTestCases = map[string]struct {
	id               int
	movie            models.Movie
	wantUpdatedMovie models.Movie
	wantErr          error
}{
	"valid data": {
		id: 3,
		movie: models.Movie{
			Id:      3,
			Title:   "Luca",
			Year:    2021,
			Runtime: 100,
			Genres:  []string{"Adventure", "Family"},
			Version: 1,
		},
		wantUpdatedMovie: models.Movie{
			Id:      3,
			Title:   "Luca",
			Year:    2021,
			Runtime: 100,
			Genres:  []string{"Adventure", "Family"},
			Version: 2,
		},
		wantErr: nil,
	},

	"wrong version number (data race)": {
		id: 99,
		movie: models.Movie{
			Id:      0,
			Title:   "Luca",
			Year:    2021,
			Runtime: 100,
			Genres:  []string{"Adventure", "Family"},
			Version: 0,
		},
		wantUpdatedMovie: models.Movie{
			Id:      0,
			Title:   "Luca",
			Year:    2021,
			Runtime: 100,
			Genres:  []string{"Adventure", "Family"},
			Version: 0,
		},
		wantErr: repository.ErrEditConflict,
	},
}

var deleteMovieTestCases = map[string]struct {
	id      int
	wantErr error
}{
	"valid id": {
		id:      1,
		wantErr: nil,
	},

	"non-existent": {
		id:      99,
		wantErr: repository.ErrRecordNotFound,
	},
}
