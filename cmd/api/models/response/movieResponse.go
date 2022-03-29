package response

import (
	"github.com/rhodeon/moviescreen/internal/types"
)

type MovieResponse struct {
	Id      int           `json:"id,omitempty"`
	Title   string        `json:"title,omitempty"`
	Year    int           `json:"year,omitempty"`
	Runtime types.Runtime `json:"runtime,omitempty"`
	Genres  []string      `json:"genres,omitempty"`
	Version int           `json:"version,omitempty"`
}
