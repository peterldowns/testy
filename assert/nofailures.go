package assert

import (
	"github.com/peterldowns/testy/check"
	"github.com/peterldowns/testy/common"
)

var (
	// NoErrors is an alias for [NoFailures]
	NoErrors = NoFailures //nolint:gochecknoglobals
	// Group is an alias for [NoFailures]
	Group = NoFailures //nolint:gochecknoglobals
	// All is an alias for [NoFailures]
	All = NoFailures //nolint:gochecknoglobals
)

// NoFailures will exit the test immediately if any checks
// have failed.
//
// NoFailures optionally accepts any number of thunks, which it runs in order,
// checking to make sure that no failure has occurred after each thunk.  This
// allows for more standard golang coding style while writing tests.
func NoFailures(t common.T, thunks ...(func() error)) {
	t.Helper()
	if failNowIfFailed(t) {
		return
	}
	for _, thunk := range thunks {
		check.Nil(t, thunk())
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
