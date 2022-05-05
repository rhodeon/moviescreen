package docs

import "github.com/rhodeon/moviescreen/cmd/api/models/response"

// HEATLHCHECK

// swagger:route GET /healthcheck Misc healthcheck
// API status.
//
// Responses:
//	200: healthcheckResponse

// swagger:response healthcheckResponse
type healthcheckResponseWrapper struct {
	// in:body
	Body response.HealthCheckResponse
}

// swagger:response emptyResponse
type emptyResponseWrapper struct {
	// in:body
	Body struct{}
}
