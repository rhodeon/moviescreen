package handlers

import (
	"expvar"
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"net/http"
)

type miscHandler struct {
	config common.Config
}

func NewMiscHandler(config common.Config) common.MiscHandler {
	return &miscHandler{config: config}
}

// HealthCheck returns the API availability status and metadata.
func (h *miscHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		response.SuccessResponse(
			http.StatusOK,
			response.HealthCheckResponse{
				Status:      "available",
				Environment: h.config.Env,
				Version:     h.config.Version,
			},
		),
	)
}

// Metrics returns a response with the server current metrics.
func (h *miscHandler) Metrics(ctx *gin.Context) {
	expvar.Handler().ServeHTTP(ctx.Writer, ctx.Request)
}
