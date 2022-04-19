package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	respErrors "github.com/rhodeon/moviescreen/cmd/api/responseErrors"
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
	"strings"
	"unicode/utf8"
)

// Authenticate determines the next course of action based on the existence and validity
// of an authentication bearer token in the request.
//
// If a valid authentication token is found, the associated user is stored in the request context before proceeding.
//
// If an invalid or malformed token is found, a 401 error is returned to the client.
//
// If no token is found, an anonymous user is stored in the request context.
func Authenticate(repositories repository.Repositories) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Vary", "Authorization")
		authorizationHeader := ctx.GetHeader("Authorization")

		// proceed as an anonymous user if no authorization token is found
		if authorizationHeader == "" {
			common.ContextSetUser(ctx, models.AnonymousUser)
			ctx.Next()
			return
		}

		// return an error for a malformed token format
		errorHandler := respErrors.NewErrorHandler()
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			errorHandler.InvalidAuthenticationToken(ctx)
			return
		}

		// validate token
		token := headerParts[1]
		if utf8.RuneCountInString(token) != 26 {
			errorHandler.InvalidAuthenticationToken(ctx)
			return
		}

		// retrieve associated user via token
		user, err := repositories.Users.GetByToken(token, models.ScopeAuthentication)
		if err != nil {
			switch {
			case errors.Is(err, repository.ErrRecordNotFound):
				errorHandler.InvalidAuthenticationToken(ctx)

			default:
				respErrors.HandleInternalServerError(ctx, err)
			}

			return
		}

		// set valid user in context
		common.ContextSetUser(ctx, user)
		ctx.Next()
	}
}
