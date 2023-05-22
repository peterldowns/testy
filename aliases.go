package check

var (
	// True
	OK = True //nolint:gochecknoglobals
	// False
	NotOk = False //nolint:gochecknoglobals
	// Enforce
	NoFailures        = Enforce //nolint:gochecknoglobals
	EnforceNoFailures = Enforce //nolint:gochecknoglobals
	Require           = Enforce //nolint:gochecknoglobals
	RequireNoFailures = Enforce //nolint:gochecknoglobals
	Verify            = Enforce //nolint:gochecknoglobals
	VerifyNoFailures  = Enforce //nolint:gochecknoglobals
)
