package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/cmd/api/responseErrors"
	"github.com/rhodeon/moviescreen/domain/repository"
	"net/http"
	"path"
	"strconv"
)

type movieHandler struct {
	config       common.Config
	repositories repository.Repositories
}

func NewMovieHandler(config common.Config, repositories repository.Repositories) common.MovieHandler {
	return &movieHandler{
		config:       config,
		repositories: repositories,
	}
}

// Create adds a new movie to the repository, and returns the newly created movie.
func (m movieHandler) Create(ctx *gin.Context) {
	// parse JSON request body
	movieRequest := &request.MovieRequest{}
	err := parseJsonRequest(ctx, movieRequest)
	if err != nil {
		return
	}

	// validate the request with all fields being mandatory for creation
	err = validateJsonRequest(ctx, movieRequest, []string{
		request.MovieFieldTitle,
		request.MovieFieldYear,
		request.MovieFieldRuntime,
		request.MovieFieldGenres,
	})
	if err != nil {
		return
	}

	// attempt to create a new movie in the repository from the request
	newMovie := movieRequest.ToModel()
	err = m.repositories.Movies.Create(&newMovie)
	if err != nil {
		responseErrors.HandleInternalServerError(ctx, err)
		return
	}

	// return the newly created movie response
	resp := newMovie.ToResponse()
	ctx.Header("Location", path.Join("/v1/movies", strconv.Itoa(resp.Id)))
	ctx.JSON(
		http.StatusCreated,
		response.SuccessResponse(
			http.StatusCreated,
			resp,
		),
	)
}

// GetById returns a movie with the specified id.
func (m movieHandler) GetById(ctx *gin.Context) {
	// validate id
	id, err := parseIdParam(ctx)
	if err != nil {
		return
	}

	// attempt to fetch movie from the repository
	movie, err := m.repositories.Movies.Get(id)
	if err != nil {
		// return a 404 error if the movie id doesn't exist in the repository
		if errors.Is(err, repository.ErrRecordNotFound) {
			responseErrors.NewErrorHandler().NotFound(ctx)
		} else {
			responseErrors.HandleInternalServerError(ctx, err)
		}
		return
	}

	// send response
	ctx.JSON(
		http.StatusOK,
		response.SuccessResponse(
			http.StatusOK,
			movie.ToResponse(),
		),
	)
}

// List returns a list of movies.
func (m movieHandler) List(ctx *gin.Context) {
	// set the queries
	queries := ctx.Request.URL.Query()
	titleQuery := parseQueryString(queries, "title", "")
	genreQuery := parseQueryCsv(queries, "genres", []string{})

	// set and validate the filters
	filers := request.Filters{
		Page:  parseQueryInt(queries, "page", 1),
		Limit: parseQueryInt(queries, "limit", 20),
		Sort:  parseQueryString(queries, "sort", "id"),
		ValidSorts: []string{
			request.MovieFilterSortId,
			request.MovieFilterSortTitle,
			request.MovieFilterSortYear,
			request.MovieFilterSortRuntime,
		},
	}

	validator := filers.Validate()
	if !validator.Valid() {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityError(validator),
		)
		return
	}

	// attempt to retrieve movies
	movies, metadata, err := m.repositories.Movies.List(titleQuery, genreQuery, filers)
	if err != nil {
		responseErrors.HandleInternalServerError(ctx, err)
		return
	}

	// return movie list and metadata response
	ctx.JSON(
		http.StatusOK,
		response.BaseResponse{
			Success:  true,
			Status:   http.StatusOK,
			Data:     movies.ToResponse(),
			Metadata: &metadata,
		},
	)
}

// Update replaces the data of the movie with the given ID query in the repository.
func (m movieHandler) Update(ctx *gin.Context) {
	// validate id
	id, err := parseIdParam(ctx)
	if err != nil {
		return
	}

	// parse JSON request body
	movieRequest := &request.MovieRequest{}
	err = parseJsonRequest(ctx, movieRequest)
	if err != nil {
		return
	}

	// validate the request with all fields being optional for update
	err = validateJsonRequest(ctx, movieRequest, []string{})
	if err != nil {
		return
	}

	// fetch and update movie from repository
	movie, err := m.repositories.Movies.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			responseErrors.NewErrorHandler().NotFound(ctx)
		} else {
			responseErrors.HandleInternalServerError(ctx, err)
		}
		return
	}
	movieRequest.UpdateModel(&movie)

	// reinsert updated movie into the repository
	err = m.repositories.Movies.Update(&movie)
	if err != nil {
		if errors.Is(err, repository.ErrEditConflict) {
			responseErrors.NewErrorHandler().EditConflict(ctx)
		} else {
			responseErrors.HandleInternalServerError(ctx, err)
		}
		return
	}

	// return updated movie
	ctx.JSON(
		http.StatusOK,
		response.SuccessResponse(
			http.StatusOK,
			movie.ToResponse(),
		),
	)
}

// Delete deletes the movie with the given id parameter from the repository.
func (m movieHandler) Delete(ctx *gin.Context) {
	// validate id
	id, err := parseIdParam(ctx)
	if err != nil {
		return
	}

	// attempt to delete movie from the repository
	err = m.repositories.Movies.Delete(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			responseErrors.NewErrorHandler().NotFound(ctx)
		} else {
			responseErrors.HandleInternalServerError(ctx, err)
		}
		return
	}

	// return a success message with an empty "data" field
	// on a successful delete
	ctx.JSON(
		http.StatusOK,
		response.SuccessResponse(
			http.StatusOK,
			map[string]string{"message": "movie deleted successfully"},
		),
	)
}
