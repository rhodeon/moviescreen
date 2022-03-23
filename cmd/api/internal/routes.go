package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/handlers"
	"path"
)

const (
	homeRoute       = "/"
	apiVersionRoute = "/v1"
)

func (app Application) Router() *gin.Engine {
	router := gin.Default()
	miscHandler := handlers.NewMiscHandler(app.Config)

	router.GET(withVersion("healthcheck"), miscHandler.HealthCheck)
	return router
}

func withVersion(relativeRoute string) string {
	return path.Join(apiVersionRoute, relativeRoute)
}
