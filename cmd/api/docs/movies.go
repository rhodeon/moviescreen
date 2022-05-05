package docs

// ROUTES

// swagger:route POST /movies/ Movies createMovie
// Create movie.
// Creates a movie with the given data in the request body.
// All fields in the request body are required.
// Requires a user with the "movies:write" permission.
//
// Responses:
//	201: movieResponse
//	400: badRequestError
//	401: unauthenticatedError
//	403: permissionError
//  422: validationError

// swagger:route GET /movies/ Movies listMovies
// List movies.
// Returns a list of movies satisfying the query parameters.
//
// Responses:
//	200: moviesResponse
//	401: unauthenticatedError
//  422: validationError

// swagger:route GET /movies/{id} Movies getMovie
// Get movie.
// Returns the details of the movie with the given id.
//
// Responses:
//	200: movieResponse
//	401: unauthenticatedError
//	404: notFoundError

// swagger:route PATCH /movies/{id} Movies updateMovie
// Update movie.
// Updates the details of the movie with the given id with those in the request body.
// Fields in the request body are optional.
// Requires a user with the "movies:write" permission.
//
// Responses:
//	200: movieResponse
//	400: badRequestError
//	401: unauthenticatedError
//	403: permissionError
//	404: notFoundError
//	409: editConflictError
//  422: validationError

// swagger:route DELETE /movies/{id} Movies deleteMovie
// Delete movie.
// Deletes the movie with the given id.
// Requires a user with the "movies:write" permission.
//
// Responses:
//	200: emptyResponse
//	401: unauthenticatedError
//	403: permissionError
//	404: notFoundError

// PARAMETERS

// swagger:parameters getMovie deleteMovie
type movieIdPath struct {
	// Movie ID.
	// in:path
	Id int `json:"id"`
}

// swagger:parameters createMovie
type movieRequestBody struct {
	// in:body
	Body struct {
		// example: For a Few Dollars More
		Title *string `json:"title"`

		// example: 1968
		Year *int `json:"year"`

		// Movie runtime in minutes
		// example: 200
		Runtime *int `json:"runtime"`

		// example: ["action", "western"]
		Genres []string `json:"genres"`
	}
}

// swagger:parameters listMovies
type listMovieQueries struct {
	// Movie title (partial or complete).
	// in: query
	Title string `json:"title"`

	// Comma-separated list of movie genres.
	// Example: genres=action,comedy
	// in: query
	Genres []string `json:"genres"`

	// Page number.
	// minimum: 1
	// maximum: 10_000_000
	// in: query
	Page int `json:"page"`

	// Number of movies per page.
	// minimum: 1
	// maximum: 100
	// in: query
	Limit int `json:"limit"`

	// Possible values: id | title | year | runtime
	// Sort values can be prefixed with a "-" to denote descending order.
	// in: query
	Sort string `json:"sort"`
}

// swagger:parameters updateMovie
type updateMovieParams struct {
	movieIdPath
	movieRequestBody
}

// RESPONSES

// swagger:response movieResponse
type movieResponseWrapper struct {
	// in: body
	Body struct {
		movieResponse
	}
}

// swagger:response moviesResponse
type moviesResponseWrapper struct {
	// in: body
	Body []movieResponse
}
