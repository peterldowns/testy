package assert

import (
	"github.com/peterldowns/testy/check"
	"github.com/peterldowns/testy/common"
)

// NoFailures will exit the test immediately if any checks have previously
// failed. This includes any calls to `t.Fail` or `t.Error`, not just those from
// Testy.
func NoFailures(t common.T, thunks ...(func())) {
	t.Helper()
	for _, thunk := range thunks {
		thunk()
		if failNowIfFailed(t) {
			break
		}
	}
	failNowIfFailed(t)
}

func NoErrors(t common.T, thunks ...(func() error)) {
	t.Helper()
	for _, thunk := range thunks {
		check.Nil(t, thunk())
		if failNowIfFailed(t) {
			break
		}
	}
	failNowIfFailed(t)
}

// failNowIfFailed immediately stops test execution if it has failed at any
// point. Returns a boolean for the purposes of testing that this framework is
// working correrctly; when we call `t.FailNow()` on a `*testing.T` the code is
// guaranteed to stop executing at that point.
func failNowIfFailed(t common.T) bool {
	t.Helper()
	if t.Failed() {
		t.FailNow()
		return true // normally unreachable, used in testing
	}
	return false
}
