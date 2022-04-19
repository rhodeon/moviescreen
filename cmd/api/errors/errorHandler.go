package errors

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
	ctx.AbortWithStatusJSON(
		http.StatusNotFound,
		response.ErrorResponse(
			http.StatusNotFound,
			response.GenericError(response.ErrMessage404),
		),
	)
}

func (e errorHandler) MethodNotAllowed(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(
		http.StatusMethodNotAllowed,
		response.ErrorResponse(
			http.StatusMethodNotAllowed,
			response.GenericError(response.ErrMessage405),
		),
	)
}

func (e errorHandler) InternalServer(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(
		http.StatusInternalServerError,
		response.ErrorResponse(
			http.StatusInternalServerError,
			response.GenericError(response.ErrMessage500),
		),
	)
}

func (e errorHandler) EditConflict(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(
		http.StatusConflict,
		response.ErrorResponse(
			http.StatusConflict,
			response.GenericError(response.ErrMessageEditConflict),
		),
	)
}

func (e errorHandler) InvalidCredentials(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(
		http.StatusUnauthorized,
		response.ErrorResponse(
			http.StatusUnauthorized,
			response.GenericError(response.ErrInvalidCredentials),
		),
	)
}

func (e errorHandler) InvalidAuthenticationToken(ctx *gin.Context) {
	ctx.Header("WWW-Authenticate", "Bearer")
	ctx.AbortWithStatusJSON(
		http.StatusUnauthorized,
		response.ErrorResponse(
			http.StatusUnauthorized,
			response.GenericError(response.ErrInvalidAuthToken),
		),
	)
}

func (e errorHandler) UnauthenticatedUser(ctx *gin.Context) {
	SetStatusAndBody(
		ctx,
		http.StatusUnauthorized,
		response.GenericError(response.ErrUnauthenticatedAccess),
	)
}

func (e errorHandler) UnactivatedUser(ctx *gin.Context) {
	SetStatusAndBody(
		ctx,
		http.StatusForbidden,
		response.GenericError(response.ErrUnactivatedAccess),
	)
}
