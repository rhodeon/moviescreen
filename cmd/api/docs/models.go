package docs

import (
	"time"
)

// ApiResponse
//
// This is the root object returned from the API.
// It wraps the success data or error along with their status codes and metadata (if applicable).
// All response examples are in actuality nested in this object, with successes being in `data` and errors in `errors`.
//
// swagger:model ApiResponse
type baseModel struct {
	// in: body
	Body struct {
		// True on a successful request, and false on an error
		// required: true
		Success bool `json:"success"`

		// Status code of the response
		// required: true
		Status int `json:"status,omitempty"`

		// Metadata of response.
		// It is only present with paginated results.
		Metadata struct {
			CurrentPage  int `json:"current_page"`
			PageLimit    int `json:"page_limit"`
			LastPage     int `json:"last_page"`
			TotalRecords int `json:"total_records"`
		} `json:"metadata"`

		// Data of a success response.
		// It is mutually exclusive to the `error`.
		Data interface{} `json:"data"`

		// Data of an error response.
		// It is mutually exclusive to the `data`.
		Error struct {
			// required: true
			Type string `json:"type"`

			// required: true
			Data map[string]string `json:"data"`
		} `json:"error"`
	}
}

// swagger:model Movie
type movieResponse struct {
	// example: 1
	Id int `json:"id"`

	// example: Harry Potter and the Philosopher's Stone
	Title string `json:"title"`

	// example: 2001
	Year int `json:"year"`

	// example: 124
	// Movie runtime in minutes.
	Runtime int `json:"runtime"`

	// example: ["adventure", "fantasy"]
	Genres []string `json:"genres"`

	// example: 1
	Version int `json:"version"`
}

// swagger:model User
type userResponse struct {
	// example: 1
	Id int `json:"id"`

	// example: johndoe
	Username string `json:"username"`

	// example: johndoe@mail.com
	Email string `json:"email"`

	// example: 1
	Version int `json:"version"`

	// example: true
	Activated bool `json:"activated"`

	Created time.Time `json:"created"`
}

// swagger:model Token
type tokenResponse struct {
	// example: OTBJEQX2EUIMKZHAMEMIPHE6TQ
	PlainText string `json:"token"`

	Expires time.Time `json:"expires"`
}
