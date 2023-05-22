package check

// T is an interface implemented by *testing.T, for compatibility
// and (lol) testing purposes.
type T interface {
	Failed() bool // yes if the test has failed

	Fail()    // mark as failed, continue
	FailNow() // mark as failed, exit

	Error(args ...any) // log a message, mark as failed, continue
}
