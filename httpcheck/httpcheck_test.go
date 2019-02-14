package httpcheck_test

import (
	"net/http"
	"testing"

	"github.com/Jose-R-Rodriguez/Golang_Assertions/httpcheck"
)

type fakeRouter struct{}

func (fr fakeRouter) ServeHTTP(http.ResponseWriter, *http.Request) {}
func ExampleTestEndpoint() {
	// this would be your testing function
	func(t *testing.T) {
		googleGetter := httpcheck.TestEndpoint("GET", "google.com", fakeRouter{})
		t.Run("this tests an empty payload to a google.com with a GET request", googleGetter(http.StatusOK, nil))
		t.Run("You could also call it to test sites dynamically like", httpcheck.TestEndpoint("POST", "facebook.com", fakeRouter{})(http.StatusAccepted, nil))
		// of course you could just use the testing function yourself like so
		httpcheck.TestEndpoint("PUT", "myspace.com", fakeRouter{})(http.StatusAccepted, nil)(t)
	}(&testing.T{})
}
