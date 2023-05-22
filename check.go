package check

// True returns x == true, marking the test as failed if not
func True(t T, x bool, args ...any) bool {
	if !x {
		t.Error(args...)
	}
	return x
}

// False returns x == false, marking the test as failed if not
func False(t T, x bool, args ...any) bool {
	return True(t, !x, args...)
}

// NoError returns x == nil, marking the test as failed if not
func NoError(t T, x error, args ...any) bool {
	if x != nil {
		t.Error(args...)
		return false
	}
	return true
}

// Error returns x != nil, marking the test as failed if not
func Error(t T, x error, args ...any) bool {
	if x == nil {
		t.Error(args...)
		return false
	}
	return true
}

// Enforce will check to see if the test has failed, and if so, immediately
// stop execution. It optionally accepts any number of thunks, which it runs in
// order, checking to make sure that no failure has occurred after each thunk.
func Enforce(t T, thunks ...(func() error)) {
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
func failNowIfFailed(t T) bool {
	if t.Failed() {
		t.FailNow()
		return true // normally unreachable, used in testing
	}
	return false
}
