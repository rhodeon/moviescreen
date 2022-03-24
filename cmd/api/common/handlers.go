package common

import (
	"github.com/gin-gonic/gin"
)

// RouteHandlers hosts the handlers to be passed into the router.
type RouteHandlers struct {
	Error  ErrorHandler
	Misc   MiscHandler
	Movies MovieHandler
}

type ErrorHandler interface {
	NotFound(ctx *gin.Context)
	MethodNotAllowed(ctx *gin.Context)
	InternalServer(ctx *gin.Context)
}

type MiscHandler interface {
	HealthCheck(ctx *gin.Context)
}

type MovieHandler interface {
	GetById(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
}
