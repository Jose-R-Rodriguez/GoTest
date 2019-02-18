package httpcheck

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jose-R-Rodriguez/Golang_Assertions/assertions"
)

type httpServer interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// EndpointTester will test an endpoint
type EndpointTester struct {
	Token    *string
	Response *httptest.ResponseRecorder
	method   string
	url      string
	router   httpServer
}

// NewEndpointTester creates a new endpoint tester object receiving a method and a url, this object's purpose is to just
// provide a constant wrap around checking response codes and executing requests to the same url with the same mmethod
func NewEndpointTester(method string, url string, router httpServer) (response *EndpointTester) {
	return &EndpointTester{method: method, url: url, router: router}
}

// NewAuthEndpointTester is the same as NewEndpointTester but sets an authentication token
func NewAuthEndpointTester(method string, url string, router httpServer, token string) *EndpointTester {
	return &EndpointTester{method: method, url: url, router: router, Token: &token}
}

// ChangeEndpoint will change the method and url up for testing
func (et *EndpointTester) ChangeEndpoint(method string, url string) {
	et.method, et.url = method, url
}

// Test will execute a request on the url with the given method and return a test function to be used in a subtest
// to describe it's functionality
func (et *EndpointTester) Test(exp int, payload []byte) func(*testing.T) {
	req, _ := http.NewRequest(et.method, et.url, bytes.NewBuffer(payload))
	if et.Token != nil {
		et.Response = ExecuteAuthorizedRequest(req, et.router, *et.Token)
	} else {
		et.Response = ExecuteRequest(req, et.router)
	}
	return func(t *testing.T) {
		ResponseCode(t, exp, et.Response.Code)
	}
}

// ResponseCode asserts both codes should be equal, otherwise fails the test
func ResponseCode(t *testing.T, expected, actual int) {
	assertions.Assert(t, assertions.FailLater, expected == actual, "Got http code %d, expected code %d\n", actual, expected)
}

// ExecuteRequest makes an http request with the router provided and returns the response
func ExecuteRequest(request *http.Request, router httpServer) (responseRecorder *httptest.ResponseRecorder) {
	responseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	return responseRecorder
}

// ExecuteAuthorizedRequest makes an http request with the router provided and returns the responseRecorder
// but will append an authorization token beforehand
func ExecuteAuthorizedRequest(request *http.Request, router httpServer, token string) (responseRecorder *httptest.ResponseRecorder) {
	request.Header.Add("Authorization", token)
	responseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	return responseRecorder
}
