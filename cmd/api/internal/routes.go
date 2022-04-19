package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/middleware"
	"github.com/rhodeon/moviescreen/domain/models"
	"path"
)

const (
	apiVersionRoute = "/v1"
)

// Router returns a gin Engine which associates the handlers with their routes.
func (app Application) Router(handlers common.RouteHandlers) *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()

	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.NoRoute(handlers.Error.NotFound)
	router.NoMethod(handlers.Error.MethodNotAllowed)

	// set general middleware
	router.Use(middleware.Metrics())
	router.Use(middleware.RateLimit(app.Config))
	router.Use(middleware.MaxSizeLimit())
	router.Use(middleware.Authenticate(app.Repositories))

	router.GET(withVersion("healthcheck"), handlers.Misc.HealthCheck)

	// metrics endpoint with "metrics:view" permission requirement
	router.GET(
		"/metrics",
		middleware.RequirePermission(models.PermissionMetricsView, app.Repositories),
		handlers.Misc.Metrics,
	)

	movies := router.Group(withVersion("movies"))
	{
		// set middleware for activation and permission requirements
		movies.Use(middleware.RequireActivatedUser())
		requireRead := middleware.RequirePermission(models.PermissionMoviesRead, app.Repositories)
		requireWrite := middleware.RequirePermission(models.PermissionMoviesWrite, app.Repositories)

		movies.GET("/", requireRead, handlers.Movies.List)
		movies.POST("/", requireWrite, handlers.Movies.Create)
		movies.GET("/:id", requireRead, handlers.Movies.GetById)
		movies.PATCH("/:id", requireWrite, handlers.Movies.Update)
		movies.DELETE("/:id", requireWrite, handlers.Movies.Delete)
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
