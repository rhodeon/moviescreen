package repository

import "errors"

var (
	ErrRecordNotFound    = errors.New("no matching record found")
	ErrEditConflict      = errors.New("edit conflict")
	ErrDuplicateUsername = errors.New("username already exists")
	ErrDuplicateEmail    = errors.New("email already exists")
)
