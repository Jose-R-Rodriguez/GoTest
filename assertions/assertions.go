package assertions

/*
Original code and idea by https://github.com/benbjohnson/testing
*/
import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// AssertNow fails the test if the condition is false with tb.FailNow
func AssertNow(tb testing.TB, condition bool, msg string, v ...interface{}) {
	assert(tb, true, condition, msg, v)
}

// Assert fails the test if the condition is false with tb.Fail
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	assert(tb, false, condition, msg, v)
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, exitOnFail, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		failAssertion(tb, exitOnFail)
	}
}

// Ok checks for unexpected errors (if err is not nil)
func Ok(tb testing.TB, err error) {
	ok(tb, false, err)
}

// OkNow checks for unexpected errors (if err is not nil)
func OkNow(tb testing.TB, err error) {
	ok(tb, true, err)
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, exitOnFail bool, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		failAssertion(tb, exitOnFail)
	}
}

// Equals fails the test if exp is not equal to act
func Equals(tb testing.TB, exp, act interface{}) {
	equals(tb, false, exp, act)
}

// EqualsNow fails the test if exp is not equal to act
func EqualsNow(tb testing.TB, exp, act interface{}) {
	equals(tb, true, exp, act)
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exitOnFail bool, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		failAssertion(tb, exitOnFail)
	}
}

// FailAssertion receives wether we want to fail our assertions immediately or not
func failAssertion(tb testing.TB, isFailNow bool) {
	if isFailNow == true {
		tb.FailNow()
	} else {
		tb.Fail()
	}
}
