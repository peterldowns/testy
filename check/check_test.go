package check_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/peterldowns/testy/assert"
	"github.com/peterldowns/testy/check"
	"github.com/peterldowns/testy/common"
)

func TestTrue(t *testing.T) {
	check.True(t, true)

	mt := &common.MockT{}
	check.True(mt, false)
	check.True(t, mt.Failed())
}

func TestFalse(t *testing.T) {
	check.False(t, false)

	mt := &common.MockT{}
	check.True(mt, true)
	check.True(mt, mt.Failed())
}

func TestEnforcePassesIfNoFailures(t *testing.T) {
	// Enforce() should pass on a brand new *testing.T
	assert.NoErrors(t)

	// Enforce() should pass on a test that has only had successful assertions
	check.True(t, 2 == 1+1)
	check.False(t, 2 == 1-9)
	assert.NoErrors(t)
}

func TestEnforceDetectsNonCheckFailures(t *testing.T) {
	// Enforce() should call FailNow() even if the failure
	// on the test was reported by a different framework.
	mt := &common.MockT{}
	mt.Error("something went wrong")
	assert.NoErrors(mt)
	check.True(t, mt.FailedNow())
}

func TestEnforceCallsFailedNow(t *testing.T) {
	// Initially, the test hasn't failed, so Enforce() doesn't call FailNow()
	mt := &common.MockT{}
	check.False(t, mt.Failed())
	assert.NoErrors(mt)
	check.False(t, mt.FailedNow())

	// Now cause the test to fail.
	check.True(mt, false)
	check.True(t, mt.Failed())
	check.False(t, mt.FailedNow())

	// Enforce() should have called FailNow()
	assert.NoErrors(mt)
	check.True(t, mt.Failed())
	check.True(t, mt.FailedNow())
}

func TestEnforceCallsThunks(t *testing.T) {
	assert.NoErrors(t, func() error {
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
	assert.NoErrors(t)
	assert.NoErrors(t, func() error {
		check.True(t, true)
		check.False(t, false)
		assert.NoErrors(t)

		assert.NoErrors(t, func() error {
			check.Nil(t, nil)
			return nil
		})

		return nil
	})
}

func TestEnforce(t *testing.T) {
	mt := &common.MockT{}
	_, err := dummyAdd(1, 2)
	check.Nil(mt, err) // passes
	_, err2 := dummyAdd(1, 0)
	check.Nil(mt, err2) // fails!
	assert.NoErrors(mt, func() error {
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

func TestError(t *testing.T) {
	var err error
	check.Nil(t, err)
	check.Nil(t, err)
	err = fmt.Errorf("example")
	check.Error(t, err)
}

type person struct {
	Name string
	Age  int
	Data map[string]any
}

func TestEqual(t *testing.T) {
	peter := person{Name: "peter", Age: 29, Data: map[string]any{
		"foo": "bar",
	}}
	johan := person{Name: "peter", Age: 29, Data: map[string]any{}}

	check.NotEqual(t, peter, johan)
	// check.Equal(t, peter, johan)

	check.True(t, true)

	t1 := time.Now()
	t2 := t1.UTC()
	check.Equal(t, t1, t2)
	if check.Equal(t, t1, t2) {
		t.Log("I knew it")
	} else {
		t.Log("what the fuck?")
	}
	// assert.DeepEqual(t, t1, t2)

	// TODO:
	// assert, require = check.Helpers()
	// check, assert = check.Helpers()
	// assert.NoFailures(t)
	// check.AssertNoFailures(t)
}
