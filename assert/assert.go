package assert

import (
	"golang.org/x/exp/constraints"

	"github.com/google/go-cmp/cmp"

	"github.com/peterldowns/testy/check"
	"github.com/peterldowns/testy/common"
)

// True passes if x == true, otherwise immediately failing the test.
func True(t common.T, x bool) {
	t.Helper()
	if !check.True(t, x) {
		t.FailNow()
	}
}

// False passes if x == false, otherwise immediately failing the test.
func False(t common.T, x bool) {
	t.Helper()
	if !check.False(t, x) {
		t.FailNow()
	}
}

// Equal passes if want == got, otherwise immediately failing the test.
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

// NotEqual passes if want != got, otherwise immediately failing the test.
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

// LessThan passes if small < big, otherwise immediately failing the test.
func LessThan[Type constraints.Ordered](t common.T, small Type, big Type) {
	t.Helper()
	if !check.LessThan(t, small, big) {
		t.FailNow()
	}
}

// LessThanOrEqual passes if small <= big, otherwise immediately failing the
// test.
func LessThanOrEqual[Type constraints.Ordered](t common.T, small Type, big Type) {
	t.Helper()
	if !check.LessThanOrEqual(t, small, big) {
		t.FailNow()
	}
}

// GreaterThan passes if big > small, otherwise immediately failing the test.
func GreaterThan[Type constraints.Ordered](t common.T, big Type, small Type) {
	t.Helper()
	if !check.GreaterThan(t, big, small) {
		t.FailNow()
	}
}

// GreaterThanOrEqual passes if big >= small, otherwise immediately failing the
// test.
func GreaterThanOrEqual[Type constraints.Ordered](t common.T, big Type, small Type) {
	t.Helper()
	if !check.GreaterThanOrEqual(t, big, small) {
		t.FailNow()
	}
}

// Error passes if err != nil, otherwise immediately failing the test.
func Error(t common.T, err error) {
	t.Helper()
	if !check.Error(t, err) {
		t.FailNow()
	}
}

// Nil passes if err == nil, otherwise immediately failing the test.
func Nil(t common.T, err error) {
	t.Helper()
	if !check.Nil(t, err) {
		t.FailNow()
	}
}

// NoError passes if err == nil otherwise immediately failing the test.
//
// NoError is an alias for [Nil]
func NoError(t common.T, err error) {
	t.Helper()
	Nil(t, err)
}

// In passes if want is an element of slice, otherwise immediately failing the
// test.
func In[Type any](t common.T, want Type, slice []Type, opts ...cmp.Option) {
	t.Helper()
	if !check.In(t, want, slice, opts...) {
		t.FailNow()
	}
}

// NotIn passes if want is an not element of slice, otherwise immediately
// failing the test.
func NotIn[Type any](t common.T, want Type, slice []Type, opts ...cmp.Option) {
	t.Helper()
	if !check.NotIn(t, want, slice, opts...) {
		t.FailNow()
	}
}
