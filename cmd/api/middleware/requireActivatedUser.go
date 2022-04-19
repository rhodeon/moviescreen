package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/errors"
)

// RequireActivatedUser allows only authenticated and activated users to proceed,
// returning a 401 error for unauthenticated users, and a 403 error for unactivated users.
func RequireActivatedUser() func(ctx *gin.Context) {
	errorHandler := errors.NewErrorHandler()

	return func(ctx *gin.Context) {
		user := common.ContextGetUser(ctx)

		if user.IsAnonymous() {
			errorHandler.UnauthenticatedUser(ctx)
			return
		}

		if !user.Activated {
			errorHandler.UnactivatedUser(ctx)
			return
		}

		// proceed with activated user
		ctx.Next()
	}
}
