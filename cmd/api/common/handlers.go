package common

import (
	"github.com/gin-gonic/gin"
)

// RouteHandlers hosts the handlers to be passed into the router.
type RouteHandlers struct {
	Error  ErrorHandler
	Misc   MiscHandler
	Movies MovieHandler
	Users  UserHandler
}

type ErrorHandler interface {
	NotFound(ctx *gin.Context)
	MethodNotAllowed(ctx *gin.Context)
	InternalServer(ctx *gin.Context)
	EditConflict(ctx *gin.Context)
	InvalidCredentials(ctx *gin.Context)
	InvalidAuthenticationToken(ctx *gin.Context)
	UnauthenticatedUser(ctx *gin.Context)
	UnactivatedUser(ctx *gin.Context)
	NotPermitted(ctx *gin.Context)
}

type MiscHandler interface {
	HealthCheck(ctx *gin.Context)
	Metrics(ctx *gin.Context)
}

type MovieHandler interface {
	GetById(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type UserHandler interface {
	Register(ctx *gin.Context)
	Activate(ctx *gin.Context)
	Authenticate(ctx *gin.Context)
	GetByEmail(ctx *gin.Context)
	Update(ctx *gin.Context)
	CreatePasswordResetToken(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
}
