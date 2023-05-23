package assert

import (
	"golang.org/x/exp/constraints"

	"github.com/peterldowns/testy/check"
	"github.com/peterldowns/testy/common"
)

// True passes if x == true, immediately failing the test if not.
func True(t common.T, x bool, args ...any) {
	t.Helper()
	check.True(t, x, args...)
	NoErrors(t)
}

// False passes if x == false, immediately failing the test if not.
func False(t common.T, x bool, args ...any) {
	t.Helper()
	check.False(t, x, args...)
	NoErrors(t)
}

// Equal passes if want == got, immediately failing the test if not.
//
// This is a typesafe check for inequality using go-cmp, allowing arguments only
// of the same type to be compared. Most of the time, this is the
// equality-checking method that you want.
//
// For more information on how this check is implemented, see
// [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func Equal[Type any](t common.T, want Type, got Type, args ...any) {
	t.Helper()
	check.Equal(t, want, got, args...)
	NoErrors(t)
}

// StrictEqual passes if want == got, immediately failing the test if not.
//
// This is a typesafe check that uses the golang == operator to compare its
// arguments. It may only be used if `want` and `got` are comparable types, defined as
//
// > (booleans, numbers, strings, pointers, channels, arrays of comparable types,
// structs whose fields are all comparable types).
//
// For "deep" or "semantic" equal comparison based on `go-cmp`, use [Equal].
func StrictEqual[Type comparable](t common.T, want Type, got Type, args ...any) {
	t.Helper()
	check.StrictEqual(t, want, got, args...)
	NoErrors(t)
}

// NotEqual passes if want != got, immediately failing the test if not.
//
// This is a typesafe check for inequality using go-cmp, allowing arguments only
// of the same type to be compared. Most of the time, this is the
// inequality-checking method that you want.
//
// For more information on how this check is implemented, see
// [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func NotEqual[Type any](t common.T, want Type, got Type, args ...any) {
	t.Helper()
	check.NotEqual(t, want, got, args...)
	NoErrors(t)
}

// StrictNotEqual passes if want != got, immediately failing the test if not.
//
// This is a typesafe check that uses the golang == operator to compare its
// arguments. It may only be used if `want` and `got` are comparable types, defined as
//
// > (booleans, numbers, strings, pointers, channels, arrays of comparable types,
// structs whose fields are all comparable types).
//
// For "deep" or "semantic" equal comparison based on `go-cmp`, use [NotEqual].
func StrictNotEqual[Type comparable](t common.T, want Type, got Type, args ...any) {
	t.Helper()
	check.StrictNotEqual(t, want, got, args...)
	NoErrors(t)
}

// LessThan passes if small < big, immediately failing the test if not.
func LessThan[Type constraints.Ordered](t common.T, small Type, big Type, args ...any) {
	t.Helper()
	check.LessThan(t, small, big, args...)
	NoErrors(t)
}

// LessThanOrEquals passes if small <= big, immediately failing the test if not.
func LessThanOrEquals[Type constraints.Ordered](t common.T, small Type, big Type, args ...any) {
	t.Helper()
	check.LessThanOrEquals(t, small, big, args...)
	NoErrors(t)
}

// GreaterThan passes if big > small, immediately failing the test if not.
func GreaterThan[Type constraints.Ordered](t common.T, big Type, small Type, args ...any) {
	t.Helper()
	check.GreaterThan(t, big, small, args...)
	NoErrors(t)
}

// GreaterThanOrEquals passes if big >= small, immediately failing the test if not.
func GreaterThanOrEquals[Type constraints.Ordered](t common.T, big Type, small Type, args ...any) {
	t.Helper()
	check.GreaterThanOrEquals(t, big, small, args...)
	NoErrors(t)
}

// Error passes if err != nil, immediately failing the test if not.
func Error(t common.T, err error, args ...any) {
	t.Helper()
	check.Error(t, err, args...)
	NoErrors(t)
}

// Nil passes if err == nil, immediately failing the test if not.
func Nil(t common.T, err error, args ...any) {
	t.Helper()
	check.Nil(t, err, args...)
	NoErrors(t)
}
