package assertions_test

import (
	"errors"
	"testing"

	"github.com/jramonrod/go-test/assertions"
)

// mocks the testing.TB interface
type mockTB struct {
	failCalls    int
	failNowCalls int
	testing.TB
}

func (mt *mockTB) Fail() {
	mt.failCalls++
}

func (mt *mockTB) FailNow() {
	mt.failNowCalls++
}

func TestAssert(t *testing.T) {
	t.Run("Assert", func(t *testing.T) {
		t.Run("fails, calls failure", func(t *testing.T) {
			mockTB := mockTB{}
			assertions.Assert(&mockTB, false, "")
			if mockTB.failCalls != 1 {
				t.Fail()
			}
		})
		t.Run("passes, doesn't call failure", func(t *testing.T) {
			mockTB := mockTB{}
			assertions.Assert(&mockTB, true, "")
			if mockTB.failCalls > 0 || mockTB.failNowCalls > 0 {
				t.Fail()
			}
		})
	})
	t.Run("AssertNow", func(t *testing.T) {
		t.Run("fails, calls failure", func(t *testing.T) {
			mockTB := mockTB{}
			assertions.AssertNow(&mockTB, false, "")
			if mockTB.failNowCalls != 1 {
				t.Fail()
			}
		})
		t.Run("passes, doesn't call failure", func(t *testing.T) {
			mockTB := mockTB{}
			assertions.AssertNow(&mockTB, true, "")
			if mockTB.failCalls > 0 || mockTB.failNowCalls > 0 {
				t.Fail()
			}
		})
	})
}

func TestOk(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		t.Run("fails, calls failure", func(t *testing.T) {
			mockTB := mockTB{}
			assertions.Ok(&mockTB, errors.New("Error"))
			if mockTB.failCalls != 1 {
				t.Fail()
			}
		})
		t.Run("passes, doesn't call failure", func(t *testing.T) {
			mockTB := mockTB{}
			assertions.Ok(&mockTB, nil)
			if mockTB.failCalls > 0 || mockTB.failNowCalls > 0 {
				t.Fail()
			}
		})
	})
	t.Run("OkNow", func(t *testing.T) {
		t.Run("fails, calls failure", func(t *testing.T) {
			mockTB := mockTB{}
			assertions.OkNow(&mockTB, errors.New("Error"))
			if mockTB.failNowCalls != 1 {
				t.Fail()
			}
		})
		t.Run("passes, doesn't call failure", func(t *testing.T) {
			mockTB := mockTB{}
			assertions.OkNow(&mockTB, nil)
			if mockTB.failCalls > 0 || mockTB.failNowCalls > 0 {
				t.Fail()
			}
		})
	})
}
