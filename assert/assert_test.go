package assert_test

import (
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
}

func TestFalse(t *testing.T) {
	t.Parallel()
	assert.False(t, false)
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

func TestEquality(t *testing.T) {
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

func TestStrictEqual(t *testing.T) {
	t.Parallel()
	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		assert.StrictEqual(t, 1, 1)

		mt := &common.MockT{}
		assert.StrictEqual(mt, 1, 2)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("custom structs", func(t *testing.T) {
		t.Parallel()
		customStruct := person{Name: "peter"}
		assert.StrictEqual(t, customStruct, customStruct)

		mt := &common.MockT{}
		assert.StrictEqual(mt, customStruct, person{Name: "bob"})
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("hidden struct fields", func(t *testing.T) {
		t.Parallel()
		// Equal works with types that have hidden fields, you just need to use a
		// cmp.Option.
		customStructWithHiddenField := hiddenPerson{Name: "Peter", hidden: true}
		assert.StrictEqual(
			t,
			customStructWithHiddenField,
			customStructWithHiddenField,
		)

		mt := &common.MockT{}
		assert.StrictEqual(
			mt,
			hiddenPerson{Name: "Peter", hidden: true},
			hiddenPerson{Name: "Peter", hidden: false},
		)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})

	t.Run("time.time structs with custom .equals", func(t *testing.T) {
		t.Parallel()
		// While Equal will consider these two structs to be the same
		// because they .Equal() each other, StrictEqual just does struct
		// comparison, and will find that they differ.
		mt := &common.MockT{}
		timeStruct := time.Now()
		assert.StrictEqual(mt, timeStruct, timeStruct.UTC())
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())

		mt = &common.MockT{}
		assert.StrictEqual(
			mt,
			timeStruct,
			timeStruct.Add(1*time.Hour),
		)
		check.True(t, mt.Failed())
		check.True(t, mt.FailedNow())
	})
}
