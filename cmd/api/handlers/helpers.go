package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/internal/validator"
	"github.com/rhodeon/prettylog"
	"net/http"
	"strconv"
)

// handleJsonRequest ensures a JSON request is properly formed and validated.
func handleJsonRequest(ctx *gin.Context, request request.ClientRequest) error {
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		// respond with a BadRequestError
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.BadRequestError(err),
		)
		return err
	}

	// validate the response fields with custom checks
	if v := request.Validate(); !v.Valid() {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityError(v),
		)
		return validator.NewError()
	}

	return err
}

// handleInternalServerError logs the error and sends
// a generic 500 error response to the client.
func handleInternalServerError(ctx *gin.Context, err error) {
	prettylog.ErrorF("internal server error: %s", err.Error())
	ctx.AbortWithStatusJSON(
		http.StatusInternalServerError,
		response.ErrorResponse(
			http.StatusInternalServerError,
			response.GenericError(response.ErrMessage500),
		),
	)
}

// parseQueryId attempts to convert the "id" query of a request
// and find its integer value.
// A 404 response is returned if a failure occurs.
func parseQueryId(ctx *gin.Context) (int, error) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		NewErrorHandler().NotFound(ctx)
		return 0, err
	}
	return id, nil
}
