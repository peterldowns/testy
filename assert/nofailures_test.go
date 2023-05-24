package assert_test

import (
	"fmt"
	"testing"

	"github.com/peterldowns/testy/assert"
	"github.com/peterldowns/testy/check"
	"github.com/peterldowns/testy/common"
)

func dummyAdd(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("cannot calculate with zero")
	}
	return a + b, nil
}

func TestNoFailuresPasses(t *testing.T) {
	// NoFailures() should pass on a brand new *testing.T
	assert.NoFailures(t)

	// NoFailures() should pass on a test that has only had successful assertions
	check.True(t, 2 == 1+1)
	check.False(t, 2 == 1-9)
	assert.NoFailures(t)
}

func TestNoErrorsPasses(t *testing.T) {
	// NoErrors() should pass on a brand new *testing.T
	assert.NoErrors(t)

	// NoErrors() should pass on a test that has only had successful assertions
	check.True(t, 2 == 1+1)
	check.False(t, 2 == 1-9)
	assert.NoErrors(t)
}

func TestNoFailuresDetectsNonCheckFailures(t *testing.T) {
	// NoFailures() should call FailNow() even if the failure
	// on the test was reported by a different framework.
	mt := &common.MockT{}
	mt.Error("something went wrong")
	assert.NoFailures(mt)
	check.True(t, mt.FailedNow())
}

func TestNoErrorsDetectsNonCheckFailures(t *testing.T) {
	// NoErrors() should call FailNow() even if the failure
	// on the test was reported by a different framework.
	mt := &common.MockT{}
	mt.Error("something went wrong")
	assert.NoErrors(mt)
	check.True(t, mt.FailedNow())
}

func TestNoFailuresCallsFailedNow(t *testing.T) {
	// Initially, the test hasn't failed, so NoFailures() doesn't call FailNow()
	mt := &common.MockT{}
	check.False(t, mt.Failed())
	assert.NoFailures(mt)
	check.False(t, mt.FailedNow())

	// Now cause the test to fail.
	check.True(mt, false)
	check.True(t, mt.Failed())
	check.False(t, mt.FailedNow())

	// NoFailures() should have called FailNow()
	assert.NoFailures(mt)
	check.True(t, mt.Failed())
	check.True(t, mt.FailedNow())
}

func TestNoErrorsCallsFailedNow(t *testing.T) {
	// Initially, the test hasn't failed, so NoErrors() doesn't call FailNow()
	mt := &common.MockT{}
	check.False(t, mt.Failed())
	assert.NoErrors(mt)
	check.False(t, mt.FailedNow())

	// Now cause the test to fail.
	check.True(mt, false)
	check.True(t, mt.Failed())
	check.False(t, mt.FailedNow())

	// NoErrors() should have called FailNow()
	assert.NoErrors(mt)
	check.True(t, mt.Failed())
	check.True(t, mt.FailedNow())
}

func TestNoFailuresCallsThunks(t *testing.T) {
	t.Parallel()
	t.Run("no failures", func(t *testing.T) {
		t.Parallel()
		calledFirst := false
		calledSecond := false
		assert.NoFailures(t, func() {
			_, err := dummyAdd(1, 2)
			check.Nil(t, err)
			res, err := dummyAdd(-1, 1)
			check.Nil(t, err)
			check.True(t, 0 == res)
			calledFirst = true
		}, func() {
			calledSecond = true
		})
		check.True(t, calledFirst)
		check.True(t, calledSecond)
	})
	t.Run("failure interrupts", func(t *testing.T) {
		t.Parallel()

		calledFirst := false
		calledSecond := false
		mt := &common.MockT{}
		assert.NoFailures(mt, func() {
			check.True(mt, false) // intentional failure
			calledFirst = true
		}, func() {
			// Never reached because of the check failure in the first thunk
			calledSecond = true
		})
		check.True(t, calledFirst)
		check.False(t, calledSecond)
	})
}

func TestNoErrorsCallsThunks(t *testing.T) {
	t.Parallel()
	t.Run("no failures", func(t *testing.T) {
		t.Parallel()
		calledFirst := false
		calledSecond := false
		assert.NoErrors(t, func() error {
			if _, err := dummyAdd(1, 2); err != nil {
				return err
			}
			res, err := dummyAdd(-1, 1)
			if err != nil {
				return err
			}
			check.True(t, 0 == res)

			calledFirst = true
			return nil
		}, func() error {
			calledSecond = true
			return nil
		})
		check.True(t, calledFirst)
		check.True(t, calledSecond)
	})
	t.Run("failure interrupts", func(t *testing.T) {
		t.Parallel()

		calledFirst := false
		calledSecond := false
		mt := &common.MockT{}
		assert.NoErrors(mt, func() error {
			calledFirst = true
			return fmt.Errorf("intentional failure")
		}, func() error {
			// Never reached because of the check failure in the first thunk
			calledSecond = true
			return nil
		})
		check.True(t, calledFirst)
		check.False(t, calledSecond)
	})
}

func TestNoErrorsNestsJustFine(t *testing.T) {
	// NoErrors() allows nesting, so you can stage your tests as you'd like
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
