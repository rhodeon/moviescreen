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

var getUserByEmailTestCases = map[string]struct {
	email    string
	wantUser models.User
	wantErr  error
}{
	"valid email": {
		email: "rhodeon@dev.mail",
		wantUser: models.User{
			Id:       1,
			Username: "rhodeon",
			Email:    "rhodeon@dev.mail",
			Password: types.Password{
				Hash: []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm"),
			},
			Activated: true,
			Version:   1,
			Created:   time.Time{},
		},
		wantErr: nil,
	},

	"invalid email": {
		email:    "wrong@email.com",
		wantUser: models.User{},
		wantErr:  repository.ErrRecordNotFound,
	},
}

var updateUserTestCases = map[string]struct {
	user        models.User
	updatedUser models.User
	wantErr     error
}{
	"valid data": {
		user: models.User{
			Id:       2,
			Username: "newname",
			Email:    "new@mail.com",
			Password: types.Password{
				Hash: []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.skdjd"),
			},
			Activated: true,
			Version:   1,
			Created:   time.Time{},
		},
		updatedUser: models.User{
			Id:       2,
			Username: "newname",
			Email:    "new@mail.com",
			Password: types.Password{
				Hash: []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.skdjd"),
			},
			Activated: true,
			Version:   2,
			Created:   time.Time{},
		},
		wantErr: nil,
	},

	"duplicate username": {
		user: models.User{
			Id:       2,
			Username: "rhodeon",
			Email:    "person@mail.com",
			Password: types.Password{
				Hash: []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm"),
			},
			Activated: true,
			Version:   1,
			Created:   time.Time{},
		},
		updatedUser: models.User{},
		wantErr:     repository.ErrDuplicateUsername,
	},

	"duplicate email": {
		user: models.User{
			Id:       2,
			Username: "person",
			Email:    "rhodeon@dev.mail",
			Password: types.Password{
				Hash: []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm"),
			},
			Activated: true,
			Version:   1,
			Created:   time.Time{},
		},
		updatedUser: models.User{},
		wantErr:     repository.ErrDuplicateEmail,
	},

	"outdated version": {
		user: models.User{
			Id:       3,
			Username: "person",
			Email:    "rhodeon@dev.mail",
			Password: types.Password{
				Hash: []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm"),
			},
			Activated: true,
			Version:   2,
			Created:   time.Time{},
		},
		updatedUser: models.User{},
		wantErr:     repository.ErrEditConflict,
	},
}
