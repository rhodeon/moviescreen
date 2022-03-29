package handlers

import (
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/internal"
	"github.com/rhodeon/moviescreen/domain/repository"
	"github.com/rhodeon/moviescreen/infrastructure/mock"
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"io"
	"net/http"
	"testing"
)

func newTestApp(t *testing.T) internal.Application {
	t.Helper()

	return internal.Application{
		Config: testConfig,
	}
}

var testConfig = common.Config{
	Env:     "testing",
	Version: "1.0.0",
	Port:    4000,
}

var testRepos = repository.Repositories{
	Movies: mock.MovieController{},
}

var testRouteHandlers = common.RouteHandlers{
	Error:  NewErrorHandler(),
	Misc:   NewMiscHandler(testConfig),
	Movies: NewMovieHandler(testConfig, testRepos),
}

// parseResponse parses a http response and returns the code, body and header.
func parseResponse(t *testing.T, result *http.Response) (int, string, http.Header) {
	t.Helper()

	body, err := io.ReadAll(result.Body)
	testhelpers.AssertFatalError(t, err)
	defer result.Body.Close()
	return result.StatusCode, string(body), result.Header
}
