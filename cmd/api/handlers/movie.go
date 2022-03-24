package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
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
				Runtime: 150,
				Genres:  []string{"action", "romance"},
				Version: 1,
			},
			{
				Id:      19489343,
				Title:   "Run of the mill",
				Runtime: 1,
				Genres:  []string{"comedy", "drama"},
				Version: 1,
			},
		},
	)
}

func (m movieHandler) List(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		response.Movie{
			Id:      19489443,
			Title:   "Where they are",
			Runtime: 150,
			Genres:  []string{"action", "romance"},
			Version: 1,
		},
	)
}
