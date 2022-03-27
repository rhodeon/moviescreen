package response

import "github.com/rhodeon/moviescreen/cmd/api/models/common/types"

type Movie struct {
	Id      int           `json:"id,omitempty"`
	Title   string        `json:"title,omitempty"`
	Year    int           `json:"year,omitempty"`
	Runtime types.Runtime `json:"runtime,omitempty"`
	Genres  []string      `json:"genres,omitempty"`
}
