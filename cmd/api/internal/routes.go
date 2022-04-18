package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/middleware"
	"path"
)

const (
	homeRoute       = "/"
	apiVersionRoute = "/v1"
)

// Router returns a gin Engine which associates the handlers with their routes.
func (app Application) Router(handlers common.RouteHandlers) *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()

	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.NoRoute(handlers.Error.NotFound)
	router.NoMethod(handlers.Error.MethodNotAllowed)
	router.Use(middleware.RateLimit(app.Config), middleware.Authenticate(app.Repositories), middleware.MaxSizeLimit())

	router.GET(withVersion("healthcheck"), handlers.Misc.HealthCheck)

	movies := router.Group(withVersion("movies"))
	{
		movies.GET("/", handlers.Movies.List)
		movies.POST("/", handlers.Movies.Create)
		movies.GET("/:id", handlers.Movies.GetById)
		movies.PATCH("/:id", handlers.Movies.Update)
		movies.DELETE("/:id", handlers.Movies.Delete)
	}

	users := router.Group(withVersion("users"))
	{
		users.POST("/", handlers.Users.Register)
		users.PUT("/activate", handlers.Users.Activate)
		users.POST("/authenticate", handlers.Users.Authenticate)
	}

	return router
}

func withVersion(relativeRoute string) string {
	return path.Join(apiVersionRoute, relativeRoute)
}
