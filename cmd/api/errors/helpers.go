package errors

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/prettylog"
	"net/http"
)

// HandleInternalServerError logs the error and sends
// a generic 500 error response to the client.
func HandleInternalServerError(ctx *gin.Context, err error) {
	prettylog.ErrorF("internal server error: %s", err.Error())
	ctx.AbortWithStatusJSON(
		http.StatusInternalServerError,
		response.ErrorResponse(
			http.StatusInternalServerError,
			response.GenericError(response.ErrMessage500),
		),
	)
}
