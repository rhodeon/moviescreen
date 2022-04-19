package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/responseErrors"
	"github.com/rhodeon/moviescreen/domain/repository"
)

// RequirePermission ensures that the authenticated user has the specified
// permission code before proceeding.
func RequirePermission(code string, repositories repository.Repositories) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// retrieve authenticated user from context
		user := common.ContextGetUser(ctx)

		// retrieve user permissions
		permissions, err := repositories.Permissions.GetAllForUser(user)
		if err != nil {
			responseErrors.HandleInternalServerError(ctx, err)
			return
		}

		// check if user permissions include the code
		if !permissions.Includes(code) {
			responseErrors.NewErrorHandler().NotPermitted(ctx)
			return
		}

		ctx.Next()
	}
}
