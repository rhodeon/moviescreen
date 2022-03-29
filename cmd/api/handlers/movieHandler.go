package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
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

// Create adds a new movie to the database, and returns the newly created movie.
func (m movieHandler) Create(ctx *gin.Context) {
	movieRequest := &request.MovieRequest{}
	err := handleJsonRequest(ctx, movieRequest)
	if err != nil {
		return
	}

	// attempt to create a new movie in the repository from the request
	newMovie := movieRequest.ToModel()
	err = m.repositories.Movies.Create(&newMovie)
	if err != nil {
		handleInternalServerError(ctx, err)
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
	id := ctx.Param("id")

	// return a 404 error if the id can't be resolved
	idValue, err := strconv.Atoi(id)
	if err != nil {
		NewErrorHandler().NotFound(ctx)
		return
	}

	// attempt to fetch movie from the repository
	movie, err := m.repositories.Movies.Get(idValue)
	if err != nil {
		// return a 404 error if the movie id doesn't exist in the repository
		if errors.Is(err, repository.ErrRecordNotFound) {
			NewErrorHandler().NotFound(ctx)
		} else {
			handleInternalServerError(ctx, err)
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
	ctx.JSON(
		http.StatusOK,
		response.SuccessResponse(
			http.StatusOK,
			[]response.MovieResponse{
				{
					Id:      19489443,
					Title:   "Where they are",
					Year:    2009,
					Runtime: 150,
					Genres:  []string{"action", "romance"},
				},
				{
					Id:      19489343,
					Title:   "Run of the mill",
					Year:    2012,
					Runtime: 1,
					Genres:  []string{"comedy", "drama"},
				},
			},
		),
	)
}
