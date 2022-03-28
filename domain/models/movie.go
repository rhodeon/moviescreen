package models

import (
	"github.com/rhodeon/moviescreen/internal/types"
	"time"
)

type Movie struct {
	Id      int
	Title   string
	Year    int
	Runtime types.Runtime
	Genres  []string
	Version int
	Created time.Time
}
