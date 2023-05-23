package check

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/constraints"

	"github.com/peterldowns/testy/common"
)

// True returns true if x == true, otherwise marks the test as failed and returns false.
func True(t common.T, x bool) bool {
	t.Helper()
	if x {
		return true
	}
	t.Error("expected true")
	return false
}

// False returns true if x == false, otherwise marks the test as failed and returns false.
func False(t common.T, x bool) bool {
	t.Helper()
	if !x {
		return true
	}
	t.Error("expected false")
	return false
}

// Equal returns true if want == got, otherwise marks the test as failed and returns false.
//
// This is a typesafe check for inequality using go-cmp, allowing arguments only
// of the same type to be compared. Most of the time, this is the
// equality-checking method that you want.
//
// You can change the behavior of the equality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func Equal[Type any](t common.T, want Type, got Type, opts ...cmp.Option) bool {
	t.Helper()
	diff := cmp.Diff(want, got, opts...)
	if diff == "" {
		return true
	}
	t.Error(fmt.Sprintf("expected want == got\n--- want\n+++ got\n%s", diff))
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
// This means that you cannot use StrictEqual on `map` or `slice` types.
// Instead, and for all other "deep" or "semantic" equality comparisons via
// `go-cmp`, use [Equal].
func StrictEqual[Type comparable](t common.T, want Type, got Type) bool {
	t.Helper()
	if want == got {
		return true
	}
	t.Error(fmt.Sprintf("expected want == got\n--- want\n+++ got\n- %+v\n+ %+v\n", want, got))
	return false
}

// NotEqual returns true if want != got, otherwise marks the test as failed and returns false.
//
// This is a typesafe check for inequality using go-cmp, allowing arguments only
// of the same type to be compared. Most of the time, this is the
// inequality-checking method that you want.
//
// You can change the behavior of the equality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func NotEqual[Type any](t common.T, want Type, got Type, opts ...cmp.Option) bool {
	t.Helper()
	if !cmp.Equal(want, got, opts...) {
		return true
	}
	t.Error(fmt.Sprintf("expected want != got\nwant: %+v\ngot: %+v", want, got))
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
// This means that you cannot use StrictNotEqual on `map` or `slice` types.
// Instead, and for all other "deep" or "semantic" equality comparisons via
// `go-cmp`, use [NotEqual].
func StrictNotEqual[Type comparable](t common.T, want Type, got Type) bool {
	t.Helper()
	if want != got {
		return true
	}
	t.Error(fmt.Sprintf("expected want != got\n%+v", want))
	return false
}

// LessThan returns true if small < big, otherwise marks the test as failed and returns false.
func LessThan[Type constraints.Ordered](t common.T, small Type, big Type) bool {
	t.Helper()
	if small < big {
		return true
	}
	t.Error(fmt.Sprintf("expected %v < %v", small, big))
	return false
}

// LessThanOrEqual returns true if small <= big, otherwise marks the test as failed and returns false.
func LessThanOrEqual[Type constraints.Ordered](t common.T, small Type, big Type) bool {
	t.Helper()
	if small <= big {
		return true
	}
	t.Error(fmt.Sprintf("expected %v <= %v", small, big))
	return false
}

// GreaterThan returns true if big > small, otherwise marks the test as failed and returns false.
func GreaterThan[Type constraints.Ordered](t common.T, big Type, small Type) bool {
	t.Helper()
	if big > small {
		return true
	}
	t.Error(fmt.Sprintf("expected %v > %v", big, small))
	return false
}

// GreaterThanOrEqual returns true if big >= small, otherwise marks the test as failed and returns false.
func GreaterThanOrEqual[Type constraints.Ordered](t common.T, big Type, small Type) bool {
	t.Helper()
	if big >= small {
		return true
	}
	t.Error(fmt.Sprintf("expected %v >= %v", big, small))
	return false
}

// Error returns true if err != nil, otherwise marks the test as failed and returns false.
func Error(t common.T, err error) bool {
	t.Helper()
	if err != nil {
		return true
	}
	t.Error("expected error, received <nil>")
	return false
}

// Nil returns true if err == nil, otherwise marks the test as failed and returns false.
func Nil(t common.T, err error) bool {
	t.Helper()
	if err == nil {
		return true
	}
	t.Error(fmt.Sprintf("expected <nil> error, received %v", err))
	return false
}
