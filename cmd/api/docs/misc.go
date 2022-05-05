package docs

// HEATLHCHECK

// swagger:route GET /healthcheck Misc healthcheck
// API status.
//
// Responses:
//	200: healthcheckResponse

// swagger:response healthcheckResponse
type healthcheckResponseWrapper struct {
	// in: body
	Body struct {
		// example: available
		Status string `json:"status"`

		// example: production
		Environment string `json:"environment"`

		// example: v1.0.0
		Version string `json:"version"`
	}
}

// swagger:response emptyResponse
type emptyResponseWrapper struct {
	// in: body
	Body struct{}
}
