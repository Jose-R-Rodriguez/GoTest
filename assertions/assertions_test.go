package assertions_test

import (
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
	t.Run("assert", func(t *testing.T) {
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
	t.Run("assertNow", func(t *testing.T) {
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
