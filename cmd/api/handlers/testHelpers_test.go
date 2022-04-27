package handlers

import (
	"fmt"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/internal"
	"github.com/rhodeon/moviescreen/cmd/api/responseErrors"
	"github.com/rhodeon/moviescreen/domain/repository"
	"github.com/rhodeon/moviescreen/infrastructure/mock"
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"io"
	"net/http"
	"reflect"
	"sync"
	"testing"
)

const mockRequestToken = "2QRJK3S54HAIUNIHNXEF4WSZSI"

func newTestApp(t *testing.T) internal.Application {
	t.Helper()

	return internal.Application{
		Config:       testConfig,
		Repositories: testRepos,
	}
}

var testConfig = common.Config{
	Env:     "testing",
	Version: "1.0.0",
	Port:    4000,
}

var testRepos = repository.Repositories{
	Tokens:      mock.NewTokenController(),
	Movies:      mock.NewMovieController(),
	Users:       mock.NewUserController(),
	Permissions: mock.NewPermissionController(),
}

var testWaitGroup = sync.WaitGroup{}

var testRouteHandlers = common.RouteHandlers{
	Error:  responseErrors.NewErrorHandler(),
	Misc:   NewMiscHandler(testConfig),
	Movies: NewMovieHandler(testConfig, testRepos),
	Users:  NewUserHandler(testConfig, testRepos, &testWaitGroup),
}

// parseResponse parses a http response and returns the code, body and header.
func parseResponse(t *testing.T, result *http.Response) (int, string, http.Header) {
	t.Helper()

	body, err := io.ReadAll(result.Body)
	testhelpers.AssertFatalError(t, err)
	defer result.Body.Close()
	return result.StatusCode, string(body), result.Header
}

func assertHeaders(t *testing.T, gotHeaders http.Header, wantHeaders http.Header) {
	t.Helper()

	for key, wantValue := range wantHeaders {
		if gotValue, exists := gotHeaders[key]; exists {
			if !reflect.DeepEqual(gotValue, wantValue) {
				t.Errorf("\nGot:\t%v\nWant:\t%v", gotValue, wantValue)
			}
		} else {
			t.Errorf("header with key %q not found", key)
		}
	}
}

func setBearerToken(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mockRequestToken))
}
