package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"path"
)

const (
	homeRoute       = "/"
	apiVersionRoute = "/v1"
)

// Router returns a gin Engine which associates the handlers with their routes.
func (app Application) Router(handlers common.RouteHandlers) *gin.Engine {
	router := gin.Default()

	router.GET(withVersion("healthcheck"), handlers.Misc.HealthCheck)
	return router
}

func withVersion(relativeRoute string) string {
	return path.Join(apiVersionRoute, relativeRoute)
}
