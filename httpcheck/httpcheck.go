package httpcheck

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jose-R-Rodriguez/Golang_Assertions/assertions"
)

type canServeHTTP interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// EndpointTester will test an endpoint
type EndpointTester struct {
	Token    *string
	Response *httptest.ResponseRecorder
}

// CreateEndPointTesterFunc is a curried function who's goal is to provide tests partially to an endpoint
// with an EndpointTester struct you can receive the response and enable authorization for the test
func (et *EndpointTester) CreateEndPointTesterFunc(method string, URL string, router canServeHTTP) func(exp int, payload []byte) func(*testing.T) {
	return func(exp int, payload []byte) func(*testing.T) {
		req, _ := http.NewRequest(method, URL, bytes.NewBuffer(payload))
		if et.Token != nil {
			et.Response = ExecuteAuthorizedRequest(req, router, *et.Token)
		} else {
			et.Response = ExecuteRequest(req, router)
		}
		return func(t *testing.T) {
			ResponseCode(t, exp, et.Response.Code)
		}
	}
}

// CreateEndPointTesterFunc is a curried function who's goal is to provide tests partially to an endpoint
// if you need use of a JWT authorization token then use the variation from the EndpointTester version
func CreateEndPointTesterFunc(method string, URL string, router canServeHTTP) func(exp int, payload []byte) func(*testing.T) {
	return func(exp int, payload []byte) func(*testing.T) {
		req, _ := http.NewRequest(method, URL, bytes.NewBuffer(payload))
		response := ExecuteRequest(req, router)
		return func(t *testing.T) {
			ResponseCode(t, exp, response.Code)
		}
	}
}

// ResponseCode asserts both codes should be equal, otherwise fails the test
func ResponseCode(t *testing.T, expected, actual int) {
	assertions.Assert(t, assertions.FailLater, expected == actual, "Got http code %d, expected code %d\n", actual, expected)
}

// ExecuteRequest makes an http request with the router provided and returns the responseRecorder
func ExecuteRequest(request *http.Request, router canServeHTTP) (responseRecorder *httptest.ResponseRecorder) {
	responseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	return responseRecorder
}

// ExecuteAuthorizedRequest makes an http request with the router provided and returns the responseRecorder
// but will append an authorization token beforehand
func ExecuteAuthorizedRequest(request *http.Request, router canServeHTTP, token string) (responseRecorder *httptest.ResponseRecorder) {
	request.Header.Add("Authorization", token)
	responseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	return responseRecorder
}
