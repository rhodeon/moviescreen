package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"net/http"
	"net/http/httptest"
	"path"
	"strconv"
	"testing"
)

func Test_movieHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := newTestApp(t)

	testCases := createMovieTestCases
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			requestJson, _ := json.Marshal(tc.request)
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/movies/", bytes.NewReader(requestJson))
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

func Test_movieHandler_GetById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := newTestApp(t)

	testCases := getMovieByIdTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, path.Join("/v1/movies", strconv.Itoa(tc.requestId)), nil)
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

func Test_movieHandler_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := newTestApp(t)

	testCases := updateMovieTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			jsonRequest, _ := json.Marshal(tc.requestBody)
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, path.Join("/v1/movies", strconv.Itoa(tc.requestId)), bytes.NewReader(jsonRequest))
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
