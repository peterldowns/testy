package assert

import (
	"golang.org/x/exp/constraints"

	"github.com/google/go-cmp/cmp"

	"github.com/peterldowns/testy/check"
	"github.com/peterldowns/testy/common"
)

// True passes if x == true, immediately failing the test if not.
func True(t common.T, x bool) {
	t.Helper()
	if !check.True(t, x) {
		t.FailNow()
	}
}

// False passes if x == false, immediately failing the test if not.
func False(t common.T, x bool) {
	t.Helper()
	if !check.False(t, x) {
		t.FailNow()
	}
}

// Equal passes if want == got, immediately failing the test if not.
//
// This is a typesafe check for inequality using go-cmp, allowing arguments only
// of the same type to be compared. Most of the time, this is the
// equality-checking method that you want.
//
// You can change the behavior of the equality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func Equal[Type any](t common.T, want Type, got Type, opts ...cmp.Option) {
	t.Helper()
	if !check.Equal(t, want, got, opts...) {
		t.FailNow()
	}
}

// StrictEqual passes if want == got, immediately failing the test if not.
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
func StrictEqual[Type comparable](t common.T, want Type, got Type) {
	t.Helper()
	if !check.StrictEqual(t, want, got) {
		t.FailNow()
	}
}

// NotEqual passes if want != got, immediately failing the test if not.
//
// This is a typesafe check for inequality using go-cmp, allowing arguments only
// of the same type to be compared. Most of the time, this is the
// inequality-checking method that you want.
//
// You can change the behavior of the equality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func NotEqual[Type any](t common.T, want Type, got Type, opts ...cmp.Option) {
	t.Helper()
	if !check.NotEqual(t, want, got, opts...) {
		t.FailNow()
	}
}

// StrictNotEqual passes if want != got, immediately failing the test if not.
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
func StrictNotEqual[Type comparable](t common.T, want Type, got Type) {
	t.Helper()
	if !check.StrictNotEqual(t, want, got) {
		t.FailNow()
	}
}

// LessThan passes if small < big, immediately failing the test if not.
func LessThan[Type constraints.Ordered](t common.T, small Type, big Type) {
	t.Helper()
	if !check.LessThan(t, small, big) {
		t.FailNow()
	}
}

// LessThanOrEquals passes if small <= big, immediately failing the test if not.
func LessThanOrEquals[Type constraints.Ordered](t common.T, small Type, big Type) {
	t.Helper()
	if !check.LessThanOrEquals(t, small, big) {
		t.FailNow()
	}
}

// GreaterThan passes if big > small, immediately failing the test if not.
func GreaterThan[Type constraints.Ordered](t common.T, big Type, small Type) {
	t.Helper()
	if !check.GreaterThan(t, big, small) {
		t.FailNow()
	}
}

// GreaterThanOrEquals passes if big >= small, immediately failing the test if not.
func GreaterThanOrEquals[Type constraints.Ordered](t common.T, big Type, small Type) {
	t.Helper()
	if !check.GreaterThanOrEquals(t, big, small) {
		t.FailNow()
	}
}

// Error passes if err != nil, immediately failing the test if not.
func Error(t common.T, err error) {
	t.Helper()
	if !check.Error(t, err) {
		t.FailNow()
	}
}

// Nil passes if err == nil, immediately failing the test if not.
func Nil(t common.T, err error) {
	t.Helper()
	if !check.Nil(t, err) {
		t.FailNow()
	}
}
