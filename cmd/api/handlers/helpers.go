package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/internal/validator"
	"github.com/rhodeon/prettylog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

// parseIdParam attempts to convert the "id" parameter of a request
// and find its integer value.
// A 404 response is returned if a failure occurs.
func parseIdParam(ctx *gin.Context) (int, error) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		NewErrorHandler().NotFound(ctx)
		return 0, err
	}
	return id, nil
}

// parseQueryString converts a url query parameter to a string.
func parseQueryString(queries url.Values, key string, defaultValue string) string {
	query := queries.Get(key)

	if query == "" {
		return defaultValue
	}
	return query
}

// parseQueryString converts a url query parameter to an integer.
func parseQueryInt(queries url.Values, key string, defaultValue int) int {
	query := queries.Get(key)

	if query == "" {
		return defaultValue
	}

	queryValue, err := strconv.Atoi(query)
	if err != nil {
		return defaultValue
	}
	return queryValue
}

// parseQueryString converts a url query parameter with multiple items to a list.
func parseQueryCsv(query url.Values, key string, defaultValue []string) []string {
	csv := query.Get(key)

	if csv == "" {
		return defaultValue
	}
	return strings.Split(csv, ",")
}
