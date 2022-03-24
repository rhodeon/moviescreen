package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"net/http"
	"strconv"
)

type movieHandler struct {
	config common.Config
}

func NewMovieHandler(config common.Config) common.MovieHandler {
	return &movieHandler{config: config}
}

// GetById returns a movie with the specified id.
func (m movieHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	// return a 404 error if the id can't be resolved
	_, err := strconv.Atoi(id)
	if err != nil {
		NewErrorHandler().NotFound(ctx)
		return
	}

	ctx.JSON(
		http.StatusOK,
		[]response.Movie{
			{
				Id:      19489443,
				Title:   "Where they are",
				Year:    2009,
				Runtime: 150,
				Genres:  []string{"action", "romance"},
				Version: 1,
			},
			{
				Id:      19489343,
				Title:   "Run of the mill",
				Year:    2012,
				Runtime: 1,
				Genres:  []string{"comedy", "drama"},
				Version: 1,
			},
		},
	)
}

// List returns a list of movies.
func (m movieHandler) List(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		response.Movie{
			Id:      19489443,
			Title:   "Where they are",
			Year:    2009,
			Runtime: 150,
			Genres:  []string{"action", "romance"},
			Version: 1,
		},
	)
}

// Create adds a new movie to the database, and returns the newly created movie.
func (m movieHandler) Create(ctx *gin.Context) {
	movieRequest := request.Movie{}

	err := ctx.BindJSON(&movieRequest)
	if err != nil {
		// return a BadRequestError on failed binding
		ctx.JSON(
			http.StatusBadRequest,
			response.BadRequestError(err),
		)
		return
	}

	// return the newly created movie
	ctx.JSON(
		http.StatusOK,
		movieRequest.ToResponse(0, 1),
	)
}
