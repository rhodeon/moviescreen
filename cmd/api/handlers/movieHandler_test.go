package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"net/http"
	"net/http/httptest"
	"path"
	"reflect"
	"strings"
	"testing"
)

func TestMovieHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := newTestApp(t)

	testCases := createMovieTestCases
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/movies/", strings.NewReader(tc.requestBody))
			app.Router(testRouteHandlers).ServeHTTP(rr, req)
			code, body, headers := parseResponse(t, rr.Result())

			// assert status code
			testhelpers.AssertEqual(t, code, tc.wantCode)

			// assert response body
			wantBody, _ := json.Marshal(tc.wantBody)
			testhelpers.AssertEqual(t, body, string(wantBody))

			// assert headers
			assertHeaders(t, headers, tc.wantHeaders)
		})
	}
}

func TestMovieHandler_GetById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := newTestApp(t)

	testCases := getMovieByIdTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, path.Join("/v1/movies", tc.requestId), nil)
			app.Router(testRouteHandlers).ServeHTTP(rr, req)
			code, body, _ := parseResponse(t, rr.Result())

			// assert status code
			testhelpers.AssertEqual(t, code, tc.wantCode)

			// assert body
			wantBody, _ := json.Marshal(tc.wantBody)
			testhelpers.AssertEqual(t, body, string(wantBody))
		})
	}
}

func TestMovieHandler_List(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := newTestApp(t)

	testCases := listMoviesTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/movies/", nil)

			q := req.URL.Query()
			if tc.titleQuery != "" {
				q.Set("title", tc.titleQuery)
			}

			if tc.genresQuery != nil {
				q.Set("genres", strings.Join(tc.genresQuery, ","))
			}

			if !reflect.DeepEqual(tc.filterQueries, request.Filters{}) {
				for k, v := range tc.filterQueries {
					q.Set(k, v)
				}
			}
			req.URL.RawQuery = q.Encode()

			app.Router(testRouteHandlers).ServeHTTP(rr, req)
			code, body, _ := parseResponse(t, rr.Result())

			// assert status code
			testhelpers.AssertEqual(t, code, tc.wantCode)

			// assert body
			wantBody, _ := json.Marshal(tc.wantBody)
			testhelpers.AssertEqual(t, body, string(wantBody))
		})
	}
}

func TestMovieHandler_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := newTestApp(t)

	testCases := updateMovieTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPatch, path.Join("/v1/movies", tc.requestId), strings.NewReader(tc.requestBody))
			app.Router(testRouteHandlers).ServeHTTP(rr, req)
			code, body, _ := parseResponse(t, rr.Result())

			// assert status code
			testhelpers.AssertEqual(t, code, tc.wantCode)

			// assert response body
			wantBody, _ := json.Marshal(tc.wantBody)
			testhelpers.AssertEqual(t, body, string(wantBody))
		})
	}
}

func TestMovieHandler_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := newTestApp(t)
	testCases := deleteMovieTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, path.Join("/v1/movies", tc.requestId), nil)
			app.Router(testRouteHandlers).ServeHTTP(rr, req)
			code, body, _ := parseResponse(t, rr.Result())

			// assert status code
			testhelpers.AssertEqual(t, code, tc.wantCode)

			// assert response body
			wantBody, _ := json.Marshal(tc.wantBody)
			testhelpers.AssertEqual(t, body, string(wantBody))
		})
	}
}
