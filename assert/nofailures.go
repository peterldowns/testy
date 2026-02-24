package assert

import (
	"github.com/peterldowns/testy/common"
)

// NoFailures will exit the test immediately if any checks have previously
// failed. Then, if supplied, thunks will be called. If the code in the thunk
// resulted in a non-fatal test error, the test is immediately failed.
func NoFailures(t common.T, thunks ...(func())) {
	t.Helper()
	if failNowIfFailed(t) {
		return
	}
	for _, thunk := range thunks {
		thunk()
		if failNowIfFailed(t) {
			break
		}
	}
}

// NoErrors will exit the test immediately if any checks have previously
// failed. Then, if supplied, thunks will be called. If the code in the thunk
// resulted in a non-fatal test error, the test is immediately failed. If the thunk
// returns an error value, the test is immediately failed.
func NoErrors(t common.T, thunks ...(func() error)) {
	t.Helper()
	if failNowIfFailed(t) {
		return
	}
	for _, thunk := range thunks {
		NoError(t, thunk())
		if failNowIfFailed(t) {
			break
		}
	}
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
