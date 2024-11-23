package common

// T is an interface implemented by *testing.T, for compatibility
// and (lol) testing purposes.
type T interface {
	// These are the methods on *testing.T that Testy actually uses.
	Error(args ...any) // log a message, mark as failed, continue
	Fail()             // mark as failed, continue
	FailNow()          // mark as failed, exit

	// These methods are needed to test Testy. (Say that three times fast)
	Helper()      // mark as a helper
	Failed() bool // yes if the test has failed
}

// MockT is designed to be used in tests to make sure that Testy fails in the
// appropriate ways.
type MockT struct {
	failed    bool
	failednow bool
}

func (t *MockT) Failed() bool {
	return t.failed
}

func (t *MockT) FailedNow() bool {
	return t.failednow
}

func (t *MockT) Fail() {
	t.failed = true
}

func (t *MockT) FailNow() {
	t.Fail()
	t.failednow = true
}

func (*MockT) Log(_ ...any) {
	// no-op
}

func (t *MockT) Error(_ ...any) {
	t.Fail()
}

func (*MockT) Helper() {
	// no-op
}
