package main_test

import (
	"fmt"
	"testing"

	"github.com/peterldowns/testy/assert"
	"github.com/peterldowns/testy/check"
)

func TestChecks(t *testing.T) {
	t.Parallel()
	check.True(t, true)
	check.False(t, false)
	check.Equal(t, []string{"hello"}, []string{"hello"})
	check.NotEqual(t,
		map[string]int{"hello": 1},
		map[string]int{"goodbye": 2},
	)
	check.LessThan(t, 1, 4)
	check.LessThanOrEqual(t, 4, 4)
	check.GreaterThan(t, 8, 6)
	check.GreaterThanOrEqual(t, 6, 6)
	check.Error(t, fmt.Errorf("oh no"))
	check.NoError(t, nil)
	check.In(t, 4, []int{2, 3, 4, 5})
	check.NotIn(t, "hello", []string{"goodbye", "world"})

	var nilm map[string]string
	check.Nil(t, nilm)
	nilm = map[string]string{"hello": "world"}
	check.NotNil(t, nilm)
}

func TestAsserts(t *testing.T) {
	t.Parallel()
	assert.True(t, true)
	assert.False(t, false)
	assert.Equal(t, []string{"hello"}, []string{"hello"})
	assert.NotEqual(t,
		map[string]int{"hello": 1},
		map[string]int{"goodbye": 2},
	)
	assert.LessThan(t, 1, 4)
	assert.LessThanOrEqual(t, 4, 4)
	assert.GreaterThan(t, 8, 6)
	assert.GreaterThanOrEqual(t, 6, 6)
	assert.Error(t, fmt.Errorf("oh no"))
	assert.NoError(t, nil)
	assert.In(t, 4, []int{2, 3, 4, 5})
	assert.NotIn(t, "hello", []string{"goodbye", "world"})

	var nilm map[string]string
	assert.Nil(t, nilm)
	nilm = map[string]string{"hello": "world"}
	assert.NotNil(t, nilm)
}
