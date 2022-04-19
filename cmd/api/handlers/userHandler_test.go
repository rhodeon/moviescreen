package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_userHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := newTestApp(t)
	testCases := RegisterUserTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/users/", strings.NewReader(tc.RequestBody))
			app.Router(testRouteHandlers).ServeHTTP(rr, req)

			code, body, _ := parseResponse(t, rr.Result())

			// assert status code
			testhelpers.AssertEqual(t, code, tc.WantCode)

			// assert response body
			wantBody, _ := json.Marshal(tc.WantBody)
			testhelpers.AssertEqual(t, body, string(wantBody))
		})
	}
}

func Test_userHandler_Activate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := newTestApp(t)
	testCases := ActivateUserTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/users/activate", strings.NewReader(tc.RequestBody))
			app.Router(testRouteHandlers).ServeHTTP(rr, req)

			code, body, _ := parseResponse(t, rr.Result())

			// assert status code
			testhelpers.AssertEqual(t, code, tc.WantCode)

			// assert response body
			wantBody, _ := json.Marshal(tc.WantBody)
			testhelpers.AssertEqual(t, body, string(wantBody))
		})
	}
}

func Test_userHandler_Authenticate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := newTestApp(t)
	testCases := AuthenticateUserTestCases

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/users/authenticate", strings.NewReader(tc.RequestBody))
			app.Router(testRouteHandlers).ServeHTTP(rr, req)

			code, body, _ := parseResponse(t, rr.Result())

			// assert status code
			testhelpers.AssertEqual(t, code, tc.WantCode)

			// assert response body
			wantBody, _ := json.Marshal(tc.WantBody)
			testhelpers.AssertEqual(t, body, string(wantBody))
		})
	}
}
