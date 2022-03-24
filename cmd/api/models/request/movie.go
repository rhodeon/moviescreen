package request

import "github.com/rhodeon/moviescreen/cmd/api/models/response"

type Movie struct {
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Runtime int      `json:"runtime"`
	Genres  []string `json:"genres"`
}

func (request Movie) ToResponse(id int, version int) response.Movie {
	return response.Movie{
		Id:      id,
		Title:   request.Title,
		Year:    request.Year,
		Runtime: response.Runtime(request.Runtime),
		Genres:  request.Genres,
		Version: version,
	}
}
