package database

import (
	"github.com/rhodeon/moviescreen/domain/repository"
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"reflect"
	"testing"
	"time"
)

func TestMovieController_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}
	testCases := createMovieTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			db, teardown := newTestDb(t)
			movieController := MovieController{Db: db}
			defer teardown()

			err := movieController.Create(&tc.movie)

			testhelpers.AssertError(t, err, tc.wantErr)

			// reset new movie creation time as it can't be tested with the current implementation
			tc.movie.Created = time.Time{}
			if !reflect.DeepEqual(tc.movie, tc.wantNewMovie) {
				t.Errorf("\nGot:\t%+v\nWant:\t%+v", tc.movie, tc.wantNewMovie)
			}
		})
	}
}

func TestMovieController_Get(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}
	testCases := getMovieTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			db, teardown := newTestDb(t)
			movieController := MovieController{Db: db}
			defer teardown()

			movie, err := movieController.Get(tc.id)

			testhelpers.AssertError(t, err, tc.wantErr)

			movie.Created = time.Time{}
			if !reflect.DeepEqual(movie, tc.wantMovie) {
				t.Errorf("\nGot:\t%+v\nWant:\t%+v", movie, tc.wantMovie)
			}
		})
	}
}

func TestMovieController_Update(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}
	testCases := updateMovieTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			db, teardown := newTestDb(t)
			movieController := MovieController{Db: db}
			defer teardown()

			err := movieController.Update(&tc.movie)
			testhelpers.AssertError(t, err, tc.wantErr)

			tc.movie.Created = time.Time{}
			if !reflect.DeepEqual(tc.movie, tc.wantUpdatedMovie) {
				t.Errorf("\nGot:\t%+v\nWant:\t%+v", tc.movie, tc.wantUpdatedMovie)
			}
		})
	}
}

func TestMovieController_Delete(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}
	testCases := deleteMovieTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			db, teardown := newTestDb(t)
			movieController := MovieController{Db: db}
			defer teardown()

			err := movieController.Delete(tc.id)
			testhelpers.AssertError(t, err, tc.wantErr)

			// check database to ensure the movie was deleted
			_, err = movieController.Get(tc.id)
			if err != repository.ErrRecordNotFound {
				t.Errorf("movie with id %d still exists in the database", tc.id)
			}
		})
	}
}
