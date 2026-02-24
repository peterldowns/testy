package assert

import (
	"cmp"

	gocmp "github.com/google/go-cmp/cmp"

	"github.com/peterldowns/testy/check"
	"github.com/peterldowns/testy/common"
)

// True passes if x == true.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
func True(t common.T, x bool) {
	t.Helper()
	if !check.True(t, x) {
		t.FailNow()
	}
}

// False passes if x == false.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
func False(t common.T, x bool) {
	t.Helper()
	if !check.False(t, x) {
		t.FailNow()
	}
}

// Equal passes if want == got.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
//
// This is a typesafe check for equality using go-cmp, allowing arguments only
// of the same type to be compared.
//
// You can change the behavior of the equality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func Equal[Type any](t common.T, want Type, got Type, opts ...gocmp.Option) {
	t.Helper()
	if !check.Equal(t, want, got, opts...) {
		t.FailNow()
	}
}

// NotEqual passes if want != got.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
//
// This is a typesafe check for inequality using go-cmp, allowing arguments only
// of the same type to be compared.
//
// You can change the behavior of the inequality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func NotEqual[Type any](t common.T, want Type, got Type, opts ...gocmp.Option) {
	t.Helper()
	if !check.NotEqual(t, want, got, opts...) {
		t.FailNow()
	}
}

// LessThan passes if small < big.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
func LessThan[Type cmp.Ordered](t common.T, small Type, big Type) {
	t.Helper()
	if !check.LessThan(t, small, big) {
		t.FailNow()
	}
}

// LessThanOrEqual passes if small <= big.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
func LessThanOrEqual[Type cmp.Ordered](t common.T, small Type, big Type) {
	t.Helper()
	if !check.LessThanOrEqual(t, small, big) {
		t.FailNow()
	}
}

// GreaterThan passes if big > small.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
func GreaterThan[Type cmp.Ordered](t common.T, big Type, small Type) {
	t.Helper()
	if !check.GreaterThan(t, big, small) {
		t.FailNow()
	}
}

// GreaterThanOrEqual passes if big >= small.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
func GreaterThanOrEqual[Type cmp.Ordered](t common.T, big Type, small Type) {
	t.Helper()
	if !check.GreaterThanOrEqual(t, big, small) {
		t.FailNow()
	}
}

// Error passes if err != nil.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
func Error(t common.T, err error) {
	t.Helper()
	if !check.Error(t, err) {
		t.FailNow()
	}
}

// NoError passes if err == nil
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
func NoError(t common.T, err error) {
	t.Helper()
	if !check.NoError(t, err) {
		t.FailNow()
	}
}

// In passes if want is an element of slice.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
//
// You can change the behavior of the equality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func In[Type any](t common.T, want Type, slice []Type, opts ...gocmp.Option) {
	t.Helper()
	if !check.In(t, want, slice, opts...) {
		t.FailNow()
	}
}

// NotIn passes if want is not an element of slice.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
//
// You can change the behavior of the inequality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func NotIn[Type any](t common.T, want Type, slice []Type, opts ...gocmp.Option) {
	t.Helper()
	if !check.NotIn(t, want, slice, opts...) {
		t.FailNow()
	}
}

// Nil passes if val == nil.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
//
// Uses reflection because Go doesn't have a type constraint for "nilable".
// Can return false for the following types:
//
//   - error
//   - pointer
//   - interface
//   - map
//   - slice
//   - channel
//   - function
//   - unsafe.Pointer
func Nil(t common.T, val any) {
	t.Helper()
	if !check.Nil(t, val) {
		t.FailNow()
	}
}

// NotNil passes if val != nil.
//
// Otherwise, the test is immediately failed and stopped with t.FailNow().
//
// Uses reflection because Go doesn't have a type constraint for "nilable".
// Can return false for the following types:
//
//   - error
//   - pointer
//   - interface
//   - map
//   - slice
//   - channel
//   - function
//   - unsafe.Pointer
func NotNil(t common.T, val any) {
	t.Helper()
	if !check.NotNil(t, val) {
		t.FailNow()
	}
}
