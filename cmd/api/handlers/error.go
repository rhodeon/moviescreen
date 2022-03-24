package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"net/http"
)

type errorHandler struct{}

func NewErrorHandler() common.ErrorHandler {
	return &errorHandler{}
}

func (e errorHandler) NotFound(ctx *gin.Context) {
	ctx.JSON(
		http.StatusNotFound,
		response.NewErrorResponse(http.StatusNotFound, response.ErrMessage404),
	)
}

func (e errorHandler) MethodNotAllowed(ctx *gin.Context) {
	ctx.JSON(
		http.StatusMethodNotAllowed,
		response.NewErrorResponse(http.StatusMethodNotAllowed, response.ErrMessage405),
	)
}

func (e errorHandler) InternalServer(ctx *gin.Context) {
	ctx.JSON(
		http.StatusMethodNotAllowed,
		response.NewErrorResponse(http.StatusInternalServerError, response.ErrMessage500),
	)
}
