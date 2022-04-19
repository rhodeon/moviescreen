package responseErrors

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
	SetStatusAndBody(
		ctx,
		http.StatusNotFound,
		response.GenericError(ErrMessageNotFound),
	)
}

func (e errorHandler) MethodNotAllowed(ctx *gin.Context) {
	SetStatusAndBody(
		ctx,
		http.StatusMethodNotAllowed,
		response.GenericError(ErrMessageNotAllowed),
	)
}

func (e errorHandler) InternalServer(ctx *gin.Context) {
	SetStatusAndBody(
		ctx,
		http.StatusInternalServerError,
		response.GenericError(ErrMessageInternalServer),
	)
}

func (e errorHandler) EditConflict(ctx *gin.Context) {
	SetStatusAndBody(
		ctx,
		http.StatusConflict,
		response.GenericError(ErrMessageEditConflict),
	)
}

func (e errorHandler) InvalidCredentials(ctx *gin.Context) {
	SetStatusAndBody(
		ctx,
		http.StatusUnauthorized,
		response.GenericError(ErrMessageInvalidCredentials),
	)
}

func (e errorHandler) InvalidAuthenticationToken(ctx *gin.Context) {
	ctx.Header("WWW-Authenticate", "Bearer")
	SetStatusAndBody(
		ctx,
		http.StatusUnauthorized,
		response.GenericError(ErrMessageInvalidAuthToken),
	)
}

func (e errorHandler) UnauthenticatedUser(ctx *gin.Context) {
	SetStatusAndBody(
		ctx,
		http.StatusUnauthorized,
		response.GenericError(ErrMessageUnauthenticatedAccess),
	)
}

func (e errorHandler) UnactivatedUser(ctx *gin.Context) {
	SetStatusAndBody(
		ctx,
		http.StatusForbidden,
		response.GenericError(ErrMessageUnactivatedAccess),
	)
}
