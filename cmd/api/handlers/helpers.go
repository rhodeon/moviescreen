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

// parseJsonRequest ensures a JSON request body is properly formed, and populates
// the request struct if so.
// Otherwise, a 400 error is returned.
func parseJsonRequest(ctx *gin.Context, request request.ClientRequest) error {
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		// respond with a BadRequestError
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.BadRequestError(err),
		)
		return err
	}
	return nil
}

// validateJsonRequest returns a 422 error if the request doesn't pass all validation rules.
func validateJsonRequest(ctx *gin.Context, request request.ClientRequest, required []string) error {
	// validate the response fields with custom checks
	if v := request.Validate(required); !v.Valid() {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityError(v),
		)
		return validator.NewError()
	}
	return nil
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

// parseParamId attempts to convert the "id" parameter of a request
// and find its integer value.
// A 404 response is returned if a failure occurs.
func parseParamId(ctx *gin.Context) (int, error) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		NewErrorHandler().NotFound(ctx)
		return 0, err
	}
	return id, nil
}
