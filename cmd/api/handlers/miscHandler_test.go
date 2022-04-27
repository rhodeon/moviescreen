package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_miscHandler_HealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/healthcheck", nil)

	app := newTestApp(t)
	app.Router(testRouteHandlers).ServeHTTP(rr, req)

	rs := rr.Result()
	code, body, _ := parseResponse(t, rs)

	testhelpers.AssertEqual(t, code, http.StatusOK)

	wantBody, _ := json.Marshal(
		response.SuccessResponse(
			http.StatusOK,
			response.HealthCheckResponse{
				Status:      "available",
				Environment: "testing",
				Version:     "v1.0.0",
			},
		),
	)
	testhelpers.AssertEqual(t, body, string(wantBody))
}
