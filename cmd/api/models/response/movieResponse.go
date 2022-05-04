package response

type MovieResponse struct {
	Id      int      `json:"id,omitempty"`
	Title   string   `json:"title,omitempty"`
	Year    int      `json:"year,omitempty"`
	Runtime int      `json:"runtime,omitempty"`
	Genres  []string `json:"genres,omitempty"`
	Version int      `json:"version,omitempty"`
}
