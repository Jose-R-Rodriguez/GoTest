package httpcheck_test

import (
	"net/http"
	"testing"

	"github.com/Jose-R-Rodriguez/Golang_Assertions/httpcheck"
)

type fakeRouter struct{}

func (fr fakeRouter) ServeHTTP(http.ResponseWriter, *http.Request) {}
func ExampleEndpointTester_Test() {
	// this would be your testing function
	func(t *testing.T) {
		googleGetter := httpcheck.NewEndpointTester("GET", "google.com", fakeRouter{})
		t.Run("this tests an empty payload to a google.com with a GET request", googleGetter.Test(323, nil))
		t.Run("this is another option", httpcheck.NewEndpointTester("GET", "google.com", fakeRouter{}).Test(http.StatusAccepted, nil))
		httpcheck.NewEndpointTester("PUT", "myspace.com", fakeRouter{}).Test(http.StatusAccepted, nil)(t)
	}(&testing.T{})
}
