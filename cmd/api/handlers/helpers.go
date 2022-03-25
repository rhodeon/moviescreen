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
		// check if the error is due to tag validations and respond
		// with a UnprocessableEntityError if so
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errs := request.ValidationErrors(ve)
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

	// validate the response fields with custom checks
	if errs, valid := request.Validate(); !valid {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityError(errs),
		)
		return errors.New("failed custom validation")
	}

	return err
}
