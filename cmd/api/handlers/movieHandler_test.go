package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"net/http"
	"net/http/httptest"
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

			code, body, _ := parseResponse(t, rr.Result())

			// assert code
			testhelpers.AssertEqual(t, code, tc.wantResponse.Status)

			// assert body
			wantBody, _ := json.Marshal(tc.wantResponse)
			testhelpers.AssertEqual(t, body, string(wantBody))

		})
	}
}
