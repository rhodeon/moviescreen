package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"net/http"
)

// handleJsonRequest ensures a JSON request is properly formed and validated.
func handleJsonRequest(ctx *gin.Context, request request.ClientRequest) error {
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		// check if the error is due to validation and respond
		// with a UnprocessableEntityError if so
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errs := request.ValidationErrors(ve)
			request.Validate(errs)
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				response.UnprocessableEntityError(errs),
			)
			return err
		}

		// otherwise, respond with a BadRequestError
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.BadRequestError(err),
		)
		return err
	}

	return err
}
