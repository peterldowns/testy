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

func TestStrictEqual(t *testing.T) {
	t.Parallel()
	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		check.StrictEqual(t, 1, 1)

		mt := &common.MockT{}
		check.StrictEqual(mt, 1, 2)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("custom structs", func(t *testing.T) {
		t.Parallel()
		customStruct := person{Name: "peter"}
		check.StrictEqual(t, customStruct, customStruct)

		mt := &common.MockT{}
		check.StrictEqual(mt, customStruct, person{Name: "bob"})
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("hidden struct fields", func(t *testing.T) {
		t.Parallel()
		// Equal works with types that have hidden fields, you just need to use a
		// cmp.Option.
		customStructWithHiddenField := hiddenPerson{Name: "Peter", hidden: true}
		check.StrictEqual(
			t,
			customStructWithHiddenField,
			customStructWithHiddenField,
		)

		mt := &common.MockT{}
		check.StrictEqual(
			mt,
			hiddenPerson{Name: "Peter", hidden: true},
			hiddenPerson{Name: "Peter", hidden: false},
		)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("time.time structs with custom .equals", func(t *testing.T) {
		t.Parallel()
		// While Equal will consider these two structs to be the same
		// because they .Equal() each other, StrictEqual just does struct
		// comparison, and will find that they differ.
		mt := &common.MockT{}
		timeStruct := time.Now()
		check.StrictEqual(mt, timeStruct, timeStruct.UTC())
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())

		mt = &common.MockT{}
		check.StrictEqual(
			mt,
			timeStruct,
			timeStruct.Add(1*time.Hour),
		)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})
}

func TestStrictNotEqual(t *testing.T) {
	t.Parallel()
	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		check.StrictNotEqual(t, 1, 2)

		mt := &common.MockT{}
		check.StrictNotEqual(mt, 1, 1)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("custom structs", func(t *testing.T) {
		t.Parallel()
		customStruct := person{Name: "peter"}
		check.StrictNotEqual(t, customStruct, person{Name: "bob"})

		mt := &common.MockT{}
		check.StrictNotEqual(mt, customStruct, customStruct)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("hidden struct fields", func(t *testing.T) {
		t.Parallel()
		// Equal works with types that have hidden fields, you just need to use a
		// cmp.Option.
		customStructWithHiddenField := hiddenPerson{Name: "Peter", hidden: true}
		check.StrictNotEqual(
			t,
			customStructWithHiddenField,
			hiddenPerson{Name: "Peter", hidden: false},
		)

		mt := &common.MockT{}
		check.StrictNotEqual(
			mt,
			customStructWithHiddenField,
			customStructWithHiddenField,
		)
		check.True(t, mt.Failed())
		check.False(t, mt.FailedNow())
	})

	t.Run("time.time structs with custom .equals", func(t *testing.T) {
		t.Parallel()
		// While Equal will consider these two structs to be the same
		// because they .Equal() each other, StrictEqual just does struct
		// comparison, and will find that they differ.
		timeStruct := time.Now()
		check.StrictNotEqual(t, timeStruct, timeStruct.UTC())
		check.StrictNotEqual(
			t,
			timeStruct,
			timeStruct.Add(1*time.Hour),
		)
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
}
