package database

import (
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
	"github.com/rhodeon/moviescreen/internal/types"
	"time"
)

var registerUserTestCases = map[string]struct {
	user           models.User
	registeredUser models.User
	wantErr        error
}{
	"valid user": {
		user: models.User{
			Username: "person",
			Email:    "person@mail.com",
			Password: types.Password{
				Hash: []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm"),
			},
		},
		registeredUser: models.User{
			Id:       4,
			Username: "person",
			Email:    "person@mail.com",
			Password: types.Password{
				Hash: []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm"),
			},
			Activated: false,
			Version:   1,
			Created:   time.Time{},
		},
		wantErr: nil,
	},

	"duplicate username": {
		user: models.User{
			Username: "rhodeon",
			Email:    "person@mail.com",
			Password: types.Password{
				Hash: []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm"),
			},
		},
		registeredUser: models.User{},
		wantErr:        repository.ErrDuplicateUsername,
	},

	"duplicate email": {
		user: models.User{
			Username: "person",
			Email:    "rhodeon@dev.mail",
			Password: types.Password{
				Hash: []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm"),
			},
		},
		registeredUser: models.User{},
		wantErr:        repository.ErrDuplicateEmail,
	},
}
