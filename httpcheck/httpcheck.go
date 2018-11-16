package httpcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jose-R-Rodriguez/Golang_Assertions/assertions"
	"github.com/gorilla/mux"
)

// ResponseCode asserts both codes should be equal, otherwise fails the test
func ResponseCode(t *testing.T, expected, actual int) {
	assertions.Assert(t, assertions.FailLater, expected == actual, "Got http code %d, expected code %d\n", actual, expected)
}

// ExecuteRequest makes an http request with the router provided and returns the responseRecorder
func ExecuteRequest(request *http.Request, router *mux.Router) (responseRecorder *httptest.ResponseRecorder) {
	responseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	return responseRecorder
}

// ExecuteAuthorizedRequest makes an http request with the router provided and returns the responseRecorder
// but will append an authorization token beforehand
func ExecuteAuthorizedRequest(request *http.Request, router *mux.Router, token string) (responseRecorder *httptest.ResponseRecorder) {
	request.Header.Add("Authorization", token)
	responseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	return responseRecorder
}
