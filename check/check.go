package check

import (
	"cmp"
	"fmt"
	"reflect"

	gocmp "github.com/google/go-cmp/cmp"

	"github.com/peterldowns/testy/common"
)

// True passes and returns true if x == true.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
func True(t common.T, x bool) bool {
	t.Helper()
	if x {
		return true
	}
	t.Error("expected true")
	return false
}

// False passes and returns true if x == false.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
func False(t common.T, x bool) bool {
	t.Helper()
	if !x {
		return true
	}
	t.Error("expected false")
	return false
}

// Equal passes and returns true if want == got.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
//
// This is a typesafe check for equality using go-cmp, allowing arguments only
// of the same type to be compared.
//
// You can change the behavior of the equality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func Equal[Type any](t common.T, want Type, got Type, opts ...gocmp.Option) bool {
	t.Helper()
	diff := gocmp.Diff(want, got, opts...)
	if diff == "" {
		return true
	}
	t.Error(fmt.Sprintf("expected want == got\n--- want\n+++ got\n%+v", diff))
	return false
}

// NotEqual passes and returns true if want != got.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
//
// This is a typesafe check for inequality using go-cmp, allowing arguments only
// of the same type to be compared.
//
// You can change the behavior of the inequality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func NotEqual[Type any](t common.T, want Type, got Type, opts ...gocmp.Option) bool {
	t.Helper()
	if !gocmp.Equal(want, got, opts...) {
		return true
	}
	t.Error(fmt.Sprintf("expected want != got\nwant: %+v\n got: %+v", want, got))
	return false
}

// LessThan passes and returns true if small < big.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
func LessThan[Type cmp.Ordered](t common.T, small Type, big Type) bool {
	t.Helper()
	if small < big {
		return true
	}
	t.Error(fmt.Sprintf("expected %+v < %+v", small, big))
	return false
}

// LessThanOrEqual passes and returns true if small <= big.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
func LessThanOrEqual[Type cmp.Ordered](t common.T, small Type, big Type) bool {
	t.Helper()
	if small <= big {
		return true
	}
	t.Error(fmt.Sprintf("expected %+v <= %+v", small, big))
	return false
}

// GreaterThan passes and returns true if big > small.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
func GreaterThan[Type cmp.Ordered](t common.T, big Type, small Type) bool {
	t.Helper()
	if big > small {
		return true
	}
	t.Error(fmt.Sprintf("expected %+v > %+v", big, small))
	return false
}

// GreaterThanOrEqual passes and returns true if big >= small.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
func GreaterThanOrEqual[Type cmp.Ordered](t common.T, big Type, small Type) bool {
	t.Helper()
	if big >= small {
		return true
	}
	t.Error(fmt.Sprintf("expected %+v >= %+v", big, small))
	return false
}

// Error passes and returns true if err != nil.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
func Error(t common.T, err error) bool {
	t.Helper()
	if err != nil {
		return true
	}
	t.Error(fmt.Sprintf("expected non-<nil> error, received %+v", err))
	return false
}

// NoError passes and returns true if err == nil
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
func NoError(t common.T, err error) bool {
	t.Helper()
	if err == nil {
		return true
	}
	t.Error(fmt.Sprintf("expected <nil> error, received %+v", err))
	return false
}

// In passes and returns true if want is an element of slice.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
//
// You can change the behavior of the equality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func In[Type any](t common.T, element Type, slice []Type, opts ...gocmp.Option) bool {
	t.Helper()
	for _, value := range slice {
		if gocmp.Equal(element, value, opts...) {
			return true
		}
	}
	t.Error(fmt.Sprintf("expected slice to contain element:\nelement: %+v\n", element))
	return false
}

// NotIn passes and returns true if want is not an element of slice.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
//
// You can change the behavior of the inequality checking using the go-cmp/cmp
// Options system. For more information, see [the go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal).
func NotIn[Type any](t common.T, element Type, slice []Type, opts ...gocmp.Option) bool {
	t.Helper()
	for _, value := range slice {
		if gocmp.Equal(element, value, opts...) {
			t.Error(fmt.Sprintf("expected slice to not contain element\nelement: %+v\n  found: %+v", element, value))
			return false
		}
	}
	return true
}

// Nil passes and returns true if val == nil.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
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
func Nil(t common.T, val any) bool {
	t.Helper()
	if isNil(val) {
		return true
	}
	t.Error(fmt.Sprintf("expected <nil>, received %+v", val))
	return false
}

// NotNil passes and returns true if val != nil.
//
// Otherwise, the test is marked as failed with t.Error(), this function returns
// false, and the test continues running.
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
func NotNil(t common.T, v any) bool {
	t.Helper()
	if !isNil(v) {
		return true
	}
	t.Error(fmt.Sprintf("expected non-<nil> value, received %+v", v))
	return false
}

// reflection-based implementation
func isNil(object any) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	switch value.Kind() {
	case
		reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
		reflect.UnsafePointer:
		return value.IsNil()
	default:
		return false
	}
}
