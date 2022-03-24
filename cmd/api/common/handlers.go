package common

import (
	"github.com/gin-gonic/gin"
)

// RouteHandlers hosts the handlers to be passed into the router.
type RouteHandlers struct {
	Misc MiscHandler
}

type MiscHandler interface {
	HealthCheck(ctx *gin.Context)
}
