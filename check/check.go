package check

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/constraints"

	"github.com/peterldowns/testy/common"
)

// True returns true if x == true, otherwise marks the test as failed and returns false.
func True(t common.T, x bool, args ...any) bool {
	t.Helper()
	if !x {
		common.Fail(t, "expected true", args...)
	}
	return x
}

// False returns true if x == false, otherwise marks the test as failed and returns false.
func False(t common.T, x bool, args ...any) bool {
	t.Helper()
	if x {
		common.Fail(t, "expected false", args...)
	}
	return !x
}

// Equal returns true if want == got, otherwise marks the test as failed and returns false.
//
// This is a typesafe check for inequality using go-cmp, allowing arguments only
// of the same type to be compared. Most of the time, this is the
// equality-checking method that you want.
//
// For more information on how this check is implemented, see
// [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func Equal[Type any](t common.T, want Type, got Type, args ...any) bool {
	t.Helper()
	diff := cmp.Diff(want, got)
	if diff == "" {
		return true
	}
	msg := fmt.Sprintf("expected want == got\n--- want\n+++ got\n%s", diff)
	common.Fail(t, msg, args...)
	return false
}

// StrictEqual returns true if want == got, otherwise marks the test as failed and returns false.
//
// This is a typesafe check that uses the golang == operator to compare its
// arguments. It may only be used if `want` and `got` are comparable types, defined as
//
// > (booleans, numbers, strings, pointers, channels, arrays of comparable types,
// structs whose fields are all comparable types).
//
// For "deep" or "semantic" equal comparison based on `go-cmp`, use [Equal].
func StrictEqual[Type comparable](t common.T, want Type, got Type, args ...any) bool {
	t.Helper()
	if want == got {
		return true
	}
	msg := fmt.Sprintf("expected want == got\n--- want\n+++ got\n- %+v\n+ %+v\n", want, got)
	common.Fail(t, msg, args...)
	return false
}

// NotEqual returns true if want != got, otherwise marks the test as failed and returns false.
//
// This is a typesafe check for inequality using go-cmp, allowing arguments only
// of the same type to be compared. Most of the time, this is the
// inequality-checking method that you want.
//
// For more information on how this check is implemented, see
// [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func NotEqual[Type any](t common.T, want Type, got Type, args ...any) bool {
	t.Helper()
	if !cmp.Equal(want, got) {
		return true
	}
	msg := fmt.Sprintf("expected want != got\nwant: %+v\ngot: %+v", want, got)
	common.Fail(t, msg, args...)
	return false
}

// StrictNotEqual returns true if want != got, otherwise marks the test as failed and returns false.
//
// This is a typesafe check that uses the golang == operator to compare its
// arguments. It may only be used if `want` and `got` are comparable types, defined as
//
// > (booleans, numbers, strings, pointers, channels, arrays of comparable types,
// structs whose fields are all comparable types).
//
// For "deep" or "semantic" equal comparison based on `go-cmp`, use [NotEqual].
func StrictNotEqual[Type comparable](t common.T, want Type, got Type, args ...any) bool {
	t.Helper()
	if want != got {
		return true
	}
	msg := fmt.Sprintf("expected want != got\n%+v", want)
	common.Fail(t, msg, args...)
	return false
}

// LessThan returns true if small < big, otherwise marks the test as failed and returns false.
func LessThan[Type constraints.Ordered](t common.T, small Type, big Type, args ...any) bool {
	t.Helper()
	if small < big {
		return true
	}
	msg := fmt.Sprintf("expected %v < %v", small, big)
	common.Fail(t, msg, args...)
	return false
}

// LessThanOrEquals returns true if small <= big, otherwise marks the test as failed and returns false.
func LessThanOrEquals[Type constraints.Ordered](t common.T, small Type, big Type, args ...any) bool {
	t.Helper()
	if small <= big {
		return true
	}
	msg := fmt.Sprintf("expected %v <= %v", small, big)
	common.Fail(t, msg, args...)
	return false
}

// GreaterThan returns true if big > small, otherwise marks the test as failed and returns false.
func GreaterThan[Type constraints.Ordered](t common.T, big Type, small Type, args ...any) bool {
	t.Helper()
	if big > small {
		return true
	}
	msg := fmt.Sprintf("expected %v > %v", big, small)
	common.Fail(t, msg, args...)
	return false
}

// GreaterThanOrEquals returns true if big >= small, otherwise marks the test as failed and returns false.
func GreaterThanOrEquals[Type constraints.Ordered](t common.T, big Type, small Type, args ...any) bool {
	t.Helper()
	if big >= small {
		return true
	}
	msg := fmt.Sprintf("expected %v >= %v", big, small)
	common.Fail(t, msg, args...)
	return false
}

// Error returns true if err != nil, otherwise marks the test as failed and returns false.
func Error(t common.T, err error, args ...any) bool {
	t.Helper()
	if err != nil {
		return true
	}
	msg := "expected error, received <nil>"
	common.Fail(t, msg, args...)
	return false
}

// Nil returns true if err == nil, otherwise marks the test as failed and returns false.
func Nil(t common.T, err error, args ...any) bool {
	t.Helper()
	if err == nil {
		return true
	}
	msg := fmt.Sprintf("expected <nil> error, received %v", err)
	common.Fail(t, msg, args...)
	return false
}
