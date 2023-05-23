package main_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/peterldowns/testy/assert"
	"github.com/peterldowns/testy/check"
)

type person struct {
	Name string
	Age  int32
}

func TestEquality(t *testing.T) {
	// These structs are not equal under deep or strict equality, they are
	// simply different.
	peter := &person{Name: "peter", Age: int32(29)}
	johan := &person{Name: "johan", Age: int32(28)}
	assert.NotEqual(t, peter, johan)
	assert.StrictNotEqual(t, peter, johan)

	// These two Time objects are equal due to the logic in their .Equal()
	// method, but they are not strictly equal because the structs contain
	// fields with different values.
	t1 := time.Now()
	t2 := t1.UTC()
	assert.Equal(t, t1, t2)
	assert.StrictNotEqual(t, t1, t2)
}

func TestStructuringHelpers(t *testing.T) {
	// This is one way to execute a series of checks
	// and exit the test if any of them have failed.
	check.Equal(t, 2, 2)
	check.LessThanOrEquals(t, 2, 3)
	check.GreaterThan(t, 3, 1)
	assert.NoFailures(t)

	// This is another, equivalent, way to do the same thing
	assert.NoFailures(t, func() error {
		check.Equal(t, 2, 2)
		check.LessThanOrEquals(t, 2, 3)
		check.GreaterThan(t, 3, 1)
		return nil
	})

	// This is a way to execute a series of checks
	// and exit the test immediately when one fails.
	_, err := myHelper(1, 2)
	assert.Nil(t, err)
	_, err = myHelper(3, -1)
	assert.Nil(t, err)
	_, err = myHelper(5, 99)
	assert.Nil(t, err)
	x, err := myHelper(10, 0)
	assert.Nil(t, err)
	assert.Equal(t, 0, x)

	// This is another, equivalent, way to do the same thing, with a more
	// standard golang style.
	assert.NoErrors(t, func() error {
		if _, err := myHelper(1, 2); err != nil {
			return err
		}
		if _, err := myHelper(3, -1); err != nil {
			return err
		}
		if _, err := myHelper(5, 99); err != nil {
			return err
		}
		x, err := myHelper(10, 0)
		if err != nil {
			return err
		}
		assert.Equal(t, 0, x)
		return nil
	})
}

func myHelper(a, b int) (int, error) {
	if a == 0 {
		return -1, fmt.Errorf("a ccannot be 0")
	}
	if b == 10 {
		return -2, fmt.Errorf("b cannot be 10")
	}
	return ((a + b) / (b - 10)) / a, nil
}

// You can use Equal and NotEqual to check if an error is valid or nil, and type
// inference handles everything correctly. Because this is so common, you can
// also use Nil and Error as shorthand.
func TestErrorComparisons(t *testing.T) {
	var x error
	check.Equal(t, nil, x)
	x = fmt.Errorf("something went wrong")
	check.NotEqual(t, nil, x)

	var y error
	check.Nil(t, y)
	y = fmt.Errorf("something went wrong")
	check.Error(t, y)
}

// You can use Equal and NotEqual to check if a pointer is valid or nil,
// and type inference handles everything correctly.
func TestPointerComparisons(t *testing.T) {
	y := "hello"
	var x *string
	check.Equal(t, nil, x)
	x = &y
	check.NotEqual(t, nil, x)
}
