package check_test

import (
	"fmt"
	"testing"

	"github.com/peterldowns/check"
)

type mockT struct {
	failed    bool
	failednow bool
}

func (t *mockT) Failed() bool {
	return t.failed
}

func (t *mockT) Fail() {
	t.failed = true
}

func (t *mockT) FailNow() {
	t.Fail()
	t.failednow = true
}

func (t *mockT) Error(_ ...any) {
	t.Fail()
}

func TestTrue(t *testing.T) {
	check.True(t, true)

	mt := &mockT{}
	check.True(mt, false)
	check.True(t, mt.Failed())
}

func TestFalse(t *testing.T) {
	check.False(t, false)

	mt := &mockT{}
	check.True(mt, true)
	check.True(mt, mt.Failed())
}

func TestEnforcePassesIfNoFailures(t *testing.T) {
	// Enforce() should pass on a brand new *testing.T
	check.Enforce(t)

	// Enforce() should pass on a test that has only had successful assertions
	check.True(t, 2 == 1+1)
	check.False(t, 2 == 1-9)
	check.Enforce(t)
}

func TestEnforceDetectsNonCheckFailures(t *testing.T) {
	// Enforce() should call FailNow() even if the failure
	// on the test was reported by a different framework.
	mt := &mockT{}
	mt.Error("something went wrong")
	check.Enforce(mt)
	check.True(t, mt.failednow)
}

func TestEnforceCallsFailedNow(t *testing.T) {
	// Initially, the test hasn't failed, so Enforce() doesn't call FailNow()
	mt := &mockT{}
	check.False(t, mt.Failed())
	check.Enforce(mt)
	check.False(t, mt.failednow)

	// Now cause the test to fail.
	check.True(mt, false)
	check.True(t, mt.Failed())
	check.False(t, mt.failednow)

	// Enforce() should have called FailNow()
	check.Enforce(mt)
	check.True(t, mt.Failed())
	check.True(t, mt.failednow)
}

func TestEnforceCallsThunks(t *testing.T) {
	check.Enforce(t, func() error {
		// Enforce() allows easily calling other functions and asserting
		// no error, with the same error checking patterns that you use in
		// the rest of your non-test code.
		if _, err := dummyAdd(1, 2); err != nil {
			return err
		}
		// Enforce() allows the use of the standard error checking
		// pattern instead of using NoError()
		res, err := dummyAdd(-1, 1)
		if err != nil {
			return err
		}
		check.True(t, 0 == res)

		if 1 == 2 {
			return fmt.Errorf("math has gone insane today")
		}
		return nil
	})
}

func TestEnforceNestsJustFine(t *testing.T) {
	// Enforce() allows nesting, so you can stage your tests as you'd like
	// if it helps you make the test more readable.
	check.Enforce(t)
	check.Enforce(t, func() error {
		check.True(t, true)
		check.False(t, false)
		check.Enforce(t)

		check.Enforce(t, func() error {
			check.NoError(t, nil)
			return nil
		})

		return nil
	})
}

func TestEnforce(t *testing.T) {
	mt := &mockT{}
	_, err := dummyAdd(1, 2)
	check.NoError(mt, err) // passes
	_, err2 := dummyAdd(1, 0)
	check.NoError(mt, err2) // fails!
	check.Enforce(mt, func() error {
		// This code should not be reached because of the earlier failure.
		t.Fatal("should not have been reached")
		return nil
	})
}

func dummyAdd(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("cannot calculate with zero")
	}
	return a + b, nil
}
