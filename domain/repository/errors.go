package repository

import "errors"

var (
	ErrRecordNotFound = errors.New("no matching record found")
	ErrEditConflict   = errors.New("edit conflict")
)
