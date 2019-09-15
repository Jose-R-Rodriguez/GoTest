package assertions_test

import (
	"errors"

	"github.com/jramonrod/go-test/assertions"
)

func ExampleOk() {
	assertions.Ok(&mockTB{}, errors.New("example error"))
	// Output: [31massertions.go:35: unexpected error: example error[39m
}

func ExampleOkNow() {
	assertions.Ok(&mockTB{}, errors.New("example error"))
	// Output: [31massertions.go:35: unexpected error: example error[39m
}

func ExampleAssert() {
	assertions.Assert(&mockTB{}, 3 == 4, "custom message %s%s", "custom str 2", "custom strings")
	// Output: [31massertions.go:21: custom message custom str 2custom strings[39m
}

func ExampleAssertNow() {
	assertions.Assert(&mockTB{}, 3 == 4, "custom message %s%s", "custom str 2", "custom strings")
	// Output: [31massertions.go:21: custom message custom str 2custom strings[39m
}

func ExampleEquals_int() {
	assertions.Equals(&mockTB{}, 44, 23)
	/* Output: [31massertions.go:55:

	exp: 44

	got: 23[39m
	*/
}

func ExampleEquals_map() {
	map1 := map[string]int{"dsds": 33, "lol": 420}
	map2 := map[string]int{"dsds": 3, "ll": 420}
	assertions.Equals(&mockTB{}, map1, map2)
	/* Output: [31massertions.go:55:

	exp: map[string]int{"dsds":33, "lol":420}

	got: map[string]int{"dsds":3, "ll":420}[39m
	*/
}
