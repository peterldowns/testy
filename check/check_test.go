package check_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/peterldowns/testy/check"
	"github.com/peterldowns/testy/common"
)

func TestTrue(t *testing.T) {
	t.Parallel()
	check.True(t, true)

	mt := &common.MockT{}
	check.True(mt, false)
	check.True(t, mt.Failed())
	check.False(t, mt.FailedNow())
}

func TestFalse(t *testing.T) {
	t.Parallel()
	check.False(t, false)

	mt := &common.MockT{}
	check.False(mt, true)
	check.True(t, mt.Failed())
	check.False(t, mt.FailedNow())
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
		check.Equal(t, 1, 1)

		mt := &common.MockT{}
		check.Equal(mt, 1, 2)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("slices", func(t *testing.T) {
		t.Parallel()
		emptySlice := []string{}
		check.Equal(t, emptySlice, emptySlice)

		mt := &common.MockT{}
		check.Equal(mt, emptySlice, []string{"hello"})
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("custom structs", func(t *testing.T) {
		t.Parallel()
		customStruct := person{Name: "peter"}
		check.Equal(t, customStruct, customStruct)

		mt := &common.MockT{}
		check.Equal(mt, customStruct, person{Name: "bob"})
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("maps", func(t *testing.T) {
		t.Parallel()
		mapdata := map[string]int{
			"hello":   1,
			"goodbye": 9,
		}
		check.Equal(t, mapdata, mapdata)

		mt := &common.MockT{}
		check.Equal(mt, mapdata, map[string]int{"hello": 1})
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("hidden struct fields with cmp.opts", func(t *testing.T) {
		t.Parallel()
		// Equal works with types that have hidden fields, you just need to use a
		// cmp.Option.
		customStructWithHiddenField := hiddenPerson{Name: "Peter", hidden: true}
		check.Equal(
			t,
			customStructWithHiddenField,
			customStructWithHiddenField,
			cmp.AllowUnexported(hiddenPerson{}),
		)

		mt := &common.MockT{}
		check.Equal(
			mt,
			hiddenPerson{Name: "Peter", hidden: true},
			hiddenPerson{Name: "Peter", hidden: false},
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("time.time structs with custom .equals", func(t *testing.T) {
		t.Parallel()
		// Shows that a custom .Equal method works -- a time.Time will be .Equal()
		// to a version of itself in a different timezone, even though the structs
		// contain strictly different values.
		timeStruct := time.Now()
		check.Equal(t, timeStruct, timeStruct.UTC())

		mt := &common.MockT{}
		check.Equal(
			mt,
			timeStruct,
			timeStruct.Add(1*time.Hour),
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
}

func TestNotEqual(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		check.NotEqual(t, 1, 2)

		mt := &common.MockT{}
		check.NotEqual(mt, 1, 1)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("slices", func(t *testing.T) {
		t.Parallel()
		check.NotEqual(t, []string{}, []string{"hello"})

		mt := &common.MockT{}
		emptySlice := []string{}
		check.NotEqual(mt, emptySlice, emptySlice)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("custom structs", func(t *testing.T) {
		t.Parallel()
		check.NotEqual(t, person{Name: "peter"}, person{Name: "bob"})

		mt := &common.MockT{}
		customStruct := person{Name: "peter"}
		check.NotEqual(mt, customStruct, customStruct)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("maps", func(t *testing.T) {
		t.Parallel()
		mapdata := map[string]int{
			"hello":   1,
			"goodbye": 9,
		}
		check.NotEqual(t, mapdata, map[string]int{"hello": 1})

		mt := &common.MockT{}
		check.NotEqual(mt, mapdata, mapdata)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("hidden struct fields with cmp.opts", func(t *testing.T) {
		t.Parallel()
		// Equal works with types that have hidden fields, you just need to use a
		// cmp.Option.
		customStructWithHiddenField := hiddenPerson{Name: "Peter", hidden: true}
		check.NotEqual(
			t,
			customStructWithHiddenField,
			hiddenPerson{Name: "Peter", hidden: false},
			cmp.AllowUnexported(hiddenPerson{}),
		)

		mt := &common.MockT{}
		check.NotEqual(
			mt,
			customStructWithHiddenField,
			customStructWithHiddenField,
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("time.time structs with custom .equals", func(t *testing.T) {
		t.Parallel()
		// Shows that a custom .Equal method works -- a time.Time will be .Equal()
		// to a version of itself in a different timezone, even though the structs
		// contain strictly different values.
		timeStruct := time.Now()
		check.NotEqual(t, timeStruct, timeStruct.Add(1*time.Hour))

		mt := &common.MockT{}
		check.NotEqual(
			mt,
			timeStruct,
			timeStruct.UTC(),
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
}

func TestLess(t *testing.T) {
	t.Parallel()
	t.Run("float", func(t *testing.T) {
		t.Parallel()
		check.LessThan(t, 1.0, 2.0)
		check.LessThanOrEqual(t, 1.0, 2.0)
		check.LessThanOrEqual(t, 1.0, 1.0)

		mt := &common.MockT{}
		check.LessThan(mt, 1.0, 1.0)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())

		mt = &common.MockT{}
		check.LessThanOrEqual(mt, 2.0, 1.0)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("int", func(t *testing.T) {
		t.Parallel()
		check.LessThan(t, 1, 2)
		check.LessThanOrEqual(t, 1, 2)
		check.LessThanOrEqual(t, 1, 1)

		mt := &common.MockT{}
		check.LessThan(mt, 1, 1)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())

		mt = &common.MockT{}
		check.LessThanOrEqual(mt, 2, 1)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		check.LessThan(t, "aaa", "bbb")
		check.LessThanOrEqual(t, "aaa", "bbb")
		check.LessThanOrEqual(t, "aaa", "aaa")

		mt := &common.MockT{}
		check.LessThan(mt, "aaa", "aaa")
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())

		mt = &common.MockT{}
		check.LessThanOrEqual(mt, "bbb", "aaa")
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("rune", func(t *testing.T) {
		t.Parallel()
		check.LessThan(t, 'a', 'b')
		check.LessThanOrEqual(t, 'a', 'b')
		check.LessThanOrEqual(t, 'a', 'a')

		mt := &common.MockT{}
		check.LessThan(mt, 'a', 'a')
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())

		mt = &common.MockT{}
		check.LessThanOrEqual(mt, 'b', 'a')
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
}

func TestGreater(t *testing.T) {
	t.Parallel()
	t.Run("float", func(t *testing.T) {
		t.Parallel()
		check.GreaterThan(t, 2.0, 1.0)
		check.GreaterThanOrEqual(t, 2.0, 1.0)
		check.GreaterThanOrEqual(t, 2.0, 2.0)

		mt := &common.MockT{}
		check.GreaterThan(mt, 2.0, 2.0)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())

		mt = &common.MockT{}
		check.GreaterThanOrEqual(mt, 1.0, 2.0)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("int", func(t *testing.T) {
		t.Parallel()
		check.GreaterThan(t, 2, 1)
		check.GreaterThanOrEqual(t, 2, 1)
		check.GreaterThanOrEqual(t, 1, 1)

		mt := &common.MockT{}
		check.GreaterThan(mt, 2, 2)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())

		mt = &common.MockT{}
		check.GreaterThanOrEqual(mt, 1, 2)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		check.GreaterThan(t, "bbb", "aaa")
		check.GreaterThanOrEqual(t, "bbb", "aaa")
		check.GreaterThanOrEqual(t, "bbb", "bbb")

		mt := &common.MockT{}
		check.GreaterThan(mt, "bbb", "bbb")
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())

		mt = &common.MockT{}
		check.GreaterThanOrEqual(mt, "aaa", "bbb")
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("rune", func(t *testing.T) {
		t.Parallel()
		check.GreaterThan(t, 'b', 'a')
		check.GreaterThanOrEqual(t, 'b', 'a')
		check.GreaterThanOrEqual(t, 'b', 'b')

		mt := &common.MockT{}
		check.GreaterThan(mt, 'b', 'b')
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())

		mt = &common.MockT{}
		check.GreaterThanOrEqual(mt, 'a', 'b')
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
}

func TestError(t *testing.T) {
	t.Parallel()
	t.Run("error", func(t *testing.T) {
		t.Parallel()
		check.Error(t, fmt.Errorf("new error"))

		mt := &common.MockT{}
		check.Error(mt, nil)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		check.Nil(t, nil)

		mt := &common.MockT{}
		check.Nil(mt, fmt.Errorf("new error"))
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("noerror", func(t *testing.T) {
		t.Parallel()
		check.NoError(t, nil)

		mt := &common.MockT{}
		check.NoError(mt, fmt.Errorf("new error"))
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
}

func TestIn(t *testing.T) {
	t.Parallel()
	t.Run("int", func(t *testing.T) {
		t.Parallel()
		check.In(t, 1, []int{1, 2, 3})

		mt := &common.MockT{}
		check.In(mt, 1, []int{4, 5, 6})
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("nil equality", func(t *testing.T) {
		t.Parallel()
		check.In(t, nil, []any{nil})

		mt := &common.MockT{}
		check.In(mt, nil, []any{})
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		check.In(t, "world", []string{"hello", "world"})

		mt := &common.MockT{}
		check.In(mt, "world", []string{"hello world"})
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("hidden struct fields with cmp.opts", func(t *testing.T) {
		t.Parallel()
		// Equal works with types that have hidden fields, you just need to use a
		// cmp.Option.
		customStructWithHiddenField := hiddenPerson{Name: "Peter", hidden: true}
		check.In(
			t,
			customStructWithHiddenField,
			[]hiddenPerson{
				customStructWithHiddenField,
			},
			cmp.AllowUnexported(hiddenPerson{}),
		)

		mt := &common.MockT{}
		check.In(
			mt,
			hiddenPerson{Name: "Peter", hidden: true},
			[]hiddenPerson{
				{Name: "Peter", hidden: false},
			},
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("time.time structs with custom .equals", func(t *testing.T) {
		t.Parallel()
		t1 := time.Now()
		t2 := t1.UTC()
		check.In(t, &t1, []*time.Time{nil, &t2})

		mt := &common.MockT{}
		check.In(
			mt,
			t1,
			[]time.Time{
				t1.Add(1 * time.Hour),
			},
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		mt := &common.MockT{}
		check.In(mt, 1, []int{})
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())

		mt = &common.MockT{}
		check.In(mt, 1, nil)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
}

func TestNotIn(t *testing.T) {
	t.Parallel()
	t.Run("int", func(t *testing.T) {
		t.Parallel()
		res := check.NotIn(t, 1, []int{4, 5, 6})
		check.True(t, res)

		mt := &common.MockT{}
		res = check.NotIn(mt, 1, []int{1, 2, 3})
		check.False(t, res)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("nil equality", func(t *testing.T) {
		t.Parallel()
		res := check.NotIn(t, nil, []any{})
		check.True(t, res)

		mt := &common.MockT{}
		res = check.NotIn(mt, nil, []any{nil})
		check.False(t, res)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		check.NotIn(t, "world", []string{"hello world"})

		mt := &common.MockT{}
		check.NotIn(mt, "world", []string{"hello", "world"})
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
	t.Run("hidden struct fields with cmp.opts", func(t *testing.T) {
		t.Parallel()
		// Equal works with types that have hidden fields, you just need to use a
		// cmp.Option.
		customStructWithHiddenField := hiddenPerson{Name: "Peter", hidden: true}
		check.NotIn(
			t,
			hiddenPerson{Name: "Peter", hidden: true},
			[]hiddenPerson{
				{Name: "Peter", hidden: false},
			},
			cmp.AllowUnexported(hiddenPerson{}),
		)

		mt := &common.MockT{}
		check.NotIn(
			mt,
			customStructWithHiddenField,
			[]hiddenPerson{
				customStructWithHiddenField,
			},
			cmp.AllowUnexported(hiddenPerson{}),
		)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("time.time structs with custom .equals", func(t *testing.T) {
		t.Parallel()
		t1 := time.Now()
		check.NotIn(
			t,
			t1,
			[]time.Time{
				t1.Add(1 * time.Hour),
			},
		)

		mt := &common.MockT{}
		t2 := t1.UTC()
		check.NotIn(mt, t1, []time.Time{t1, t2})
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		check.NotIn(t, 1, []int{})
		check.NotIn(t, 1, nil)
	})
}
