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
	t.Error(fmt.Sprintf("expected want != got\nwant: %+v\n got: %+v", want, got))
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

// NoError returned true if err == nil, otherwise marks the test as failed and returns false.
//
// NoError is an alias for [Nil]
func NoError(t common.T, err error) bool {
	t.Helper()
	return Nil(t, err)
}

// In returns true if element is in slice, otherwise marks the test as failed
// and returns false.
func In[Type any](t common.T, element Type, slice []Type, opts ...cmp.Option) bool {
	t.Helper()
	for _, value := range slice {
		if cmp.Equal(element, value, opts...) {
			return true
		}
	}
	t.Error(fmt.Sprintf("expected slice to contain element:\nelement: %+v\n", element))
	return false
}

// NotIn returns true if element is not in slice, otherwise marks the test as
// failed and returns false.
func NotIn[Type any](t common.T, element Type, slice []Type, opts ...cmp.Option) bool {
	t.Helper()
	for _, value := range slice {
		if cmp.Equal(element, value, opts...) {
			t.Error(fmt.Sprintf("expected slice to not contain element\nelement: %+v\n  found: %+v", element, value))
			return false
		}
	}
	return true
}
