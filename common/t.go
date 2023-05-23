package common

// T is an interface implemented by *testing.T, for compatibility
// and (lol) testing purposes.
type T interface {
	Failed() bool // yes if the test has failed

	Fail()    // mark as failed, continue
	FailNow() // mark as failed, exit

	Log(args ...any)
	Error(args ...any) // log a message, mark as failed, continue

	Helper() // mark as a helper
}

/*
output from require.Less(t, 3, 1, "something went wrong")

--- FAIL: TestEqual (0.00s)
	/Users/pd/code/check/check_test.go:175:
			Error Trace:	/Users/pd/code/check/check_test.go:175
			Error:      	"3" is not less than "1"
			Test:       	TestEqual
			Messages:   	something went wrong
*/

func Fail(t T, msg string, args ...any) {
	t.Helper()
	if msg != "" {
		t.Log(msg)
	}
	if len(args) > 0 {
		t.Log(args...)
	}
	t.Fail()
}
