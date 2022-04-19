package responseErrors

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/prettylog"
	"net/http"
)

// SetStatusAndBody sets the status code and the body of the context response
// to the given values.
func SetStatusAndBody(ctx *gin.Context, statusCode int, body response.Error) {
	ctx.AbortWithStatusJSON(
		statusCode,
		response.ErrorResponse(statusCode, body),
	)
}

// HandleInternalServerError logs the error and sends
// a generic 500 error response to the client.
func HandleInternalServerError(ctx *gin.Context, err error) {
	prettylog.ErrorF("internal server error: %s", err.Error())
	SetStatusAndBody(
		ctx,
		http.StatusInternalServerError,
		response.GenericError(ErrMessageInternalServer),
	)
}
