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

const (
	// FailImmediately is a simple constant to make function calls more verbose
	FailImmediately = true
	// FailLater is a simple constant to make function calls more verbose
	FailLater = false
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, exitOnFail, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		failAssertion(tb, exitOnFail)
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, exitOnFail bool, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		failAssertion(tb, exitOnFail)
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exitOnFail bool, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		failAssertion(tb, exitOnFail)
	}
}

// FailAssertion receives wether we want to fail our assertions immediately or not
func failAssertion(tb testing.TB, failTime bool) {
	if failTime == FailImmediately {
		tb.FailNow()
	} else {
		tb.Fail()
	}
}
