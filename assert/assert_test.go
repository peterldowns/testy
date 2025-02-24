package assert_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/peterldowns/testy/assert"
	"github.com/peterldowns/testy/check"
	"github.com/peterldowns/testy/common"
)

func TestTrue(t *testing.T) {
	t.Parallel()
	assert.True(t, true)

	mt := &common.MockT{}
	assert.True(mt, false)
	check.True(t, mt.Failed())
	check.True(t, mt.FailedNow())
}

func TestFalse(t *testing.T) {
	t.Parallel()
	assert.False(t, false)

	mt := &common.MockT{}
	assert.False(mt, true)
	check.True(t, mt.Failed())
	check.True(t, mt.FailedNow())
}

type person struct {
	Name string
}

type hiddenPerson struct {
	Name string
	// When using Equal / NotEqual, this unexported field will require a custom
	// cmp.Option.
	hidden bool
}

func TestEqual(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, 1, 1)

		mt := &common.MockT{}
		assert.Equal(mt, 1, 2)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("slices", func(t *testing.T) {
		t.Parallel()
		emptySlice := []string{}
		assert.Equal(t, emptySlice, emptySlice)

		mt := &common.MockT{}
		assert.Equal(mt, emptySlice, []string{"hello"})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("custom structs", func(t *testing.T) {
		t.Parallel()
		customStruct := person{Name: "peter"}
		assert.Equal(t, customStruct, customStruct)

		mt := &common.MockT{}
		assert.Equal(mt, customStruct, person{Name: "bob"})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("maps", func(t *testing.T) {
		t.Parallel()
		mapdata := map[string]int{
			"hello":   1,
			"goodbye": 9,
		}
		assert.Equal(t, mapdata, mapdata)

		mt := &common.MockT{}
		assert.Equal(mt, mapdata, map[string]int{"hello": 1})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("hidden struct fields with cmp.opts", func(t *testing.T) {
		t.Parallel()
		// Equal works with types that have hidden fields, you just need to use a
		// cmp.Option.
		customStructWithHiddenField := hiddenPerson{Name: "Peter", hidden: true}
		assert.Equal(
			t,
			customStructWithHiddenField,
			customStructWithHiddenField,
			cmp.AllowUnexported(hiddenPerson{}),
		)

		mt := &common.MockT{}
		assert.Equal(
			mt,
			hiddenPerson{Name: "Peter", hidden: true},
			hiddenPerson{Name: "Peter", hidden: false},
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("time.time structs with custom .equals", func(t *testing.T) {
		t.Parallel()
		// Shows that a custom .Equal method works -- a time.Time will be .Equal()
		// to a version of itself in a different timezone, even though the structs
		// contain strictly different values.
		timeStruct := time.Now()
		assert.Equal(t, timeStruct, timeStruct.UTC())

		mt := &common.MockT{}
		assert.Equal(
			mt,
			timeStruct,
			timeStruct.Add(1*time.Hour),
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
}

func TestNotEqual(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		assert.NotEqual(t, 1, 2)

		mt := &common.MockT{}
		assert.NotEqual(mt, 1, 1)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("slices", func(t *testing.T) {
		t.Parallel()
		assert.NotEqual(t, []string{}, []string{"hello"})

		mt := &common.MockT{}
		emptySlice := []string{}
		assert.NotEqual(mt, emptySlice, emptySlice)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("custom structs", func(t *testing.T) {
		t.Parallel()
		assert.NotEqual(t, person{Name: "peter"}, person{Name: "bob"})

		mt := &common.MockT{}
		customStruct := person{Name: "peter"}
		assert.NotEqual(mt, customStruct, customStruct)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("maps", func(t *testing.T) {
		t.Parallel()
		mapdata := map[string]int{
			"hello":   1,
			"goodbye": 9,
		}
		assert.NotEqual(t, mapdata, map[string]int{"hello": 1})

		mt := &common.MockT{}
		assert.NotEqual(mt, mapdata, mapdata)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("hidden struct fields with cmp.opts", func(t *testing.T) {
		t.Parallel()
		// Equal works with types that have hidden fields, you just need to use a
		// cmp.Option.
		customStructWithHiddenField := hiddenPerson{Name: "Peter", hidden: true}
		assert.NotEqual(
			t,
			customStructWithHiddenField,
			hiddenPerson{Name: "Peter", hidden: false},
			cmp.AllowUnexported(hiddenPerson{}),
		)

		mt := &common.MockT{}
		assert.NotEqual(
			mt,
			customStructWithHiddenField,
			customStructWithHiddenField,
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("time.time structs with custom .equals", func(t *testing.T) {
		t.Parallel()
		// Shows that a custom .Equal method works -- a time.Time will be .Equal()
		// to a version of itself in a different timezone, even though the structs
		// contain strictly different values.
		timeStruct := time.Now()
		assert.NotEqual(t, timeStruct, timeStruct.Add(1*time.Hour))

		mt := &common.MockT{}
		assert.NotEqual(
			mt,
			timeStruct,
			timeStruct.UTC(),
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
}

func TestLess(t *testing.T) {
	t.Parallel()
	t.Run("float", func(t *testing.T) {
		t.Parallel()
		assert.LessThan(t, 1.0, 2.0)
		assert.LessThanOrEqual(t, 1.0, 2.0)
		assert.LessThanOrEqual(t, 1.0, 1.0)

		mt := &common.MockT{}
		assert.LessThan(mt, 1.0, 1.0)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())

		mt = &common.MockT{}
		assert.LessThanOrEqual(mt, 2.0, 1.0)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("int", func(t *testing.T) {
		t.Parallel()
		assert.LessThan(t, 1, 2)
		assert.LessThanOrEqual(t, 1, 2)
		assert.LessThanOrEqual(t, 1, 1)

		mt := &common.MockT{}
		assert.LessThan(mt, 1, 1)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())

		mt = &common.MockT{}
		assert.LessThanOrEqual(mt, 2, 1)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		assert.LessThan(t, "aaa", "bbb")
		assert.LessThanOrEqual(t, "aaa", "bbb")
		assert.LessThanOrEqual(t, "aaa", "aaa")

		mt := &common.MockT{}
		assert.LessThan(mt, "aaa", "aaa")
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())

		mt = &common.MockT{}
		assert.LessThanOrEqual(mt, "bbb", "aaa")
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("rune", func(t *testing.T) {
		t.Parallel()
		assert.LessThan(t, 'a', 'b')
		assert.LessThanOrEqual(t, 'a', 'b')
		assert.LessThanOrEqual(t, 'a', 'a')

		mt := &common.MockT{}
		assert.LessThan(mt, 'a', 'a')
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())

		mt = &common.MockT{}
		assert.LessThanOrEqual(mt, 'b', 'a')
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
}

func TestGreater(t *testing.T) {
	t.Parallel()
	t.Run("float", func(t *testing.T) {
		t.Parallel()
		assert.GreaterThan(t, 2.0, 1.0)
		assert.GreaterThanOrEqual(t, 2.0, 1.0)
		assert.GreaterThanOrEqual(t, 2.0, 2.0)

		mt := &common.MockT{}
		assert.GreaterThan(mt, 2.0, 2.0)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())

		mt = &common.MockT{}
		assert.GreaterThanOrEqual(mt, 1.0, 2.0)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("int", func(t *testing.T) {
		t.Parallel()
		assert.GreaterThan(t, 2, 1)
		assert.GreaterThanOrEqual(t, 2, 1)
		assert.GreaterThanOrEqual(t, 1, 1)

		mt := &common.MockT{}
		assert.GreaterThan(mt, 2, 2)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())

		mt = &common.MockT{}
		assert.GreaterThanOrEqual(mt, 1, 2)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		assert.GreaterThan(t, "bbb", "aaa")
		assert.GreaterThanOrEqual(t, "bbb", "aaa")
		assert.GreaterThanOrEqual(t, "bbb", "bbb")

		mt := &common.MockT{}
		assert.GreaterThan(mt, "bbb", "bbb")
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())

		mt = &common.MockT{}
		assert.GreaterThanOrEqual(mt, "aaa", "bbb")
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("rune", func(t *testing.T) {
		t.Parallel()
		assert.GreaterThan(t, 'b', 'a')
		assert.GreaterThanOrEqual(t, 'b', 'a')
		assert.GreaterThanOrEqual(t, 'b', 'b')

		mt := &common.MockT{}
		assert.GreaterThan(mt, 'b', 'b')
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())

		mt = &common.MockT{}
		assert.GreaterThanOrEqual(mt, 'a', 'b')
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
}

func TestError(t *testing.T) {
	t.Parallel()
	t.Run("error", func(t *testing.T) {
		t.Parallel()
		assert.Error(t, fmt.Errorf("new error"))

		mt := &common.MockT{}
		assert.Error(mt, nil)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		assert.Nil(t, nil)

		mt := &common.MockT{}
		assert.Nil(mt, fmt.Errorf("new error"))
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("noerror", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, nil)

		mt := &common.MockT{}
		assert.NoError(mt, fmt.Errorf("new error"))
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
}

func TestIn(t *testing.T) {
	t.Parallel()
	t.Run("int", func(t *testing.T) {
		t.Parallel()
		assert.In(t, 1, []int{1, 2, 3})

		mt := &common.MockT{}
		assert.In(mt, 1, []int{4, 5, 6})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("nil equality", func(t *testing.T) {
		t.Parallel()
		assert.In(t, nil, []any{nil})

		mt := &common.MockT{}
		assert.In(mt, nil, []any{})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		assert.In(t, "world", []string{"hello", "world"})

		mt := &common.MockT{}
		assert.In(mt, "world", []string{"hello world"})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("hidden struct fields with cmp.opts", func(t *testing.T) {
		t.Parallel()
		// Equal works with types that have hidden fields, you just need to use a
		// cmp.Option.
		customStructWithHiddenField := hiddenPerson{Name: "Peter", hidden: true}
		assert.In(
			t,
			customStructWithHiddenField,
			[]hiddenPerson{
				customStructWithHiddenField,
			},
			cmp.AllowUnexported(hiddenPerson{}),
		)

		mt := &common.MockT{}
		assert.In(
			mt,
			hiddenPerson{Name: "Peter", hidden: true},
			[]hiddenPerson{
				{Name: "Peter", hidden: false},
			},
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("time.time structs with custom .equals", func(t *testing.T) {
		t.Parallel()
		t1 := time.Now()
		t2 := t1.UTC()
		assert.In(t, &t1, []*time.Time{nil, &t2})

		mt := &common.MockT{}
		assert.In(
			mt,
			t1,
			[]time.Time{
				t1.Add(1 * time.Hour),
			},
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		mt := &common.MockT{}
		assert.In(mt, 1, []int{})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())

		mt = &common.MockT{}
		assert.In(mt, 1, nil)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
}

func TestNotIn(t *testing.T) {
	t.Parallel()
	t.Run("int", func(t *testing.T) {
		t.Parallel()
		assert.NotIn(t, 1, []int{4, 5, 6})

		mt := &common.MockT{}
		assert.NotIn(mt, 1, []int{1, 2, 3})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("nil equality", func(t *testing.T) {
		t.Parallel()
		assert.NotIn(t, nil, []any{})

		mt := &common.MockT{}
		assert.NotIn(mt, nil, []any{nil})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		assert.NotIn(t, "world", []string{"hello world"})

		mt := &common.MockT{}
		assert.NotIn(mt, "world", []string{"hello", "world"})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
	t.Run("hidden struct fields with cmp.opts", func(t *testing.T) {
		t.Parallel()
		// Equal works with types that have hidden fields, you just need to use a
		// cmp.Option.
		customStructWithHiddenField := hiddenPerson{Name: "Peter", hidden: true}
		assert.NotIn(
			t,
			hiddenPerson{Name: "Peter", hidden: true},
			[]hiddenPerson{
				{Name: "Peter", hidden: false},
			},
			cmp.AllowUnexported(hiddenPerson{}),
		)

		mt := &common.MockT{}
		assert.NotIn(
			mt,
			customStructWithHiddenField,
			[]hiddenPerson{
				customStructWithHiddenField,
			},
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("time.time structs with custom .equals", func(t *testing.T) {
		t.Parallel()
		t1 := time.Now()
		assert.NotIn(
			t,
			t1,
			[]time.Time{
				t1.Add(1 * time.Hour),
			},
		)

		mt := &common.MockT{}
		t2 := t1.UTC()
		assert.NotIn(mt, t1, []time.Time{t1, t2})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		assert.NotIn(t, 1, []int{})
		assert.NotIn(t, 1, nil)
	})
}
