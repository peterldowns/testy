# 😤 Testy

![Latest Version](https://badgers.space/badge/latest%20version/v0.0.8/blueviolet?corner_radius=m)
![Golang](https://badgers.space/badge/golang/1.21+/blue?corner_radius=m)

Testy is a library for writing meaningful, readable, and maintainable tests.
Testy is typesafe (using generics), based on [go-cmp](https://github.com/google/go-cmp), and designed as an alternative
to [testify](https://github.com/stretchr/testify), [gotools.test](https://github.com/gotestyourself/gotest.tools), and [is](https://github.com/matryer/is).

Major features:
- Typesafe comparisons mean that when you refactor your code, your tests
  can be easily updated.
- A limited number of methods makes it easier to write tests by constraining your options.
- Extendable: if you need a more complicated assertion method, write your own helper functions.
- Deep equality testing by default using [go-cmp](https://github.com/google/go-cmp).
- **checks**: soft assertions using `t.Error`, the test is marked as failed but will continue running.
- **asserts**: hard assertions using `t.FailNow`, the test is marked as failed and stops immediately.
- Optional helpers for structuring tests in more readable ways.

## Install

```shell
go get github.com/peterldowns/testy@latest
```

## Documentation
- [The github README, https://github.com/peterldowns/testy](https://github.com/peterldowns/testy)
- [The go.dev docs, pkg.go.dev/github.com/peterldowns/testy](https://pkg.go.dev/github.com/peterldowns/testy)

The github README is the primary source for documentation. The code itself is
supposed to be well-organized, and each function has a meaningful docstring, so
you should be able to explore it quite easily using an LSP plugin, reading the
code, or clicking through the go.dev docs. 


## API

- `True(t, x)` checks if `x == true`
- `False(t, x)` checks if `x == false`
- `Equal(t, want, got)` checks if its arguments are equal using [go-cmp](https://github.com/google/go-cmp)
- `NotEqual(t, want, got)` checks if its arguments are not equal using [go-cmp](https://github.com/google/go-cmp)
- `LessThan(t, small, big)` checks if `small < big`
- `LessThanOrEqual(t, small, big)` checks if `small <= big`
- `GreaterThan(t, big, small)` checks if `big > small`
- `GreaterThanOrEqual(t, big, small)` checks if `big >= small`
- `Error(t, err)` checks if `err == nil`
- `NoError(t, err)` checks if `err != nil`
- `In(t, item, slice)` checks if `item in slice` using [go-cmp](https://github.com/google/go-cmp)
- `NotIn(t, item, slice)` checks if `item not in slice` using [go-cmp](https://github.com/google/go-cmp)
- `Nil(t, val)` checks if `val == nil` using reflection to support any nilable value.
- `NotNil(t, val)` checks if `val != nil` using reflection to support any nilable value.

```go
package api_test

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
```

# Details

## `check` methods call `t.Error`
`check` contains methods for checking a condition, marking the test as failed
but allowing it to continue running if the condition is not met.  This is a
"soft" style assert, equivalent to the methods in `testify/assert` or the
`Check` method in `gotest.tools/assert`.  If a check fails, testy calls
`t.Error`.

Each `check` method returns a boolean, which is `true` if the check passed, and
`false` otherwise. You can use this to conditionally run other logic in your
code.

```go
func TestExample(t *testing.T) {
    var f *MyFoo
    var err error
    f, err = ServiceThatGetsAFoo()
    // f is only meaningful if err == nil
    if check.NoError(t, err) {
        check.Equal(f.Name, "peter")
    }
}
```

## `assert` methods call `t.FailNow`
`assert` contains methods for asserting a condition, marking the test as failed
and immediately exiting the test if the condition is not met. This is a "hard"
or traditional assert, equivalent to the methods in `testify/require` or the
`Assert` method in `gotest.tools/assert`. If an assertion fails, testy calls
`t.FailNow`.

```go
func TestExample(t *testing.T) {
    var f *MyFoo
    var err error
    f, err = ServiceThatGetsAFoo()
    assert.NoError(t, err) // if err != nil, the test will end here
    assert.Equal(f.Name, "peter")
}
```

## Structuring Helpers
The `assert` package also provides helpers for structuring your tests and making
them more expressive. You can use these helpers to determine which checks are
run in parallel, and which checks should halt test execution.

- `assert.NoFailures(t)` and `assert.NoErrors(t)` will instantly fail the test
if any previous `check` has failed, or any other code has called
`t.Fail()`/`t.Error()` for any reason.
- `assert.NoFailures(t, thunks...(func()))` will run each thunk function. After
each thunk, it will check that no failures have occurred. If they have, the test
immediately exits without running any of the following thunks.
- `assert.NoErrors(t, thunks...(func() error))` will run each thunk function. If
the thunk returns an error, or any check failure has occurred, the test
immediately exits without running any of the following thunks.

You can use these to stage your tests and make them more useful.

```go
func TestStructuringHelpers(t *testing.T) {
    // Here's an easy way to "stage" your tests, where
    // - each stage runs all of the checks
    // - at the end, report any failures.
    // - if there were any failures, end the test.
    check.Equal(t, 2, 2)
    check.LessThanOrEqual(t, 2, 3)
    check.GreaterThan(t, 3, 1)
    assert.NoFailures(t)

    // This is another, equivalent, way to express the same logic
    assert.NoFailures(t, func() {
        check.Equal(t, 2, 2)
        check.LessThanOrEqual(t, 2, 3)
        check.GreaterThan(t, 3, 1)
    })

    // You can also use the helpers to adopt a standard error-returning style
    assert.NoErrors(t, func() error {
        if _, err := myHelper(baz, bar); err != nil {
            return err
        }
        if _, err := anotherF(bar); err != nil {
            return err
        }
        return nil
    })

    // This is equivalent to
    _, err := myHelper(baz, bar)
    assert.NoError(err)
    _, err = anotherF(bar)
    assert.NoError(err)
}
```

## More Examples

Beyond the examples presented in this README, please read the code and its tests
to see `testy` in action.

# FAQ

## Why create this library?
Each of the existing libraries (testify, gotest.tools, is) are phenomenal projects and Testy
has been deeply influenced by each of them, and could not exist without them. While writing Testy, I used those libraries as reference, and I have great respect for their authors and contributors.

That said, these libraries are (in my opinion) too limited in terms of expressiveness. Tests serve multiple purposes:
- they prove the ground-level correctness of the system
- they prove that a bug was fixed 
    - therefore, they act as a historical record of all discovered bugs
- they are documentation of the current implementation of the system and how to use its components
- they are documentation of the business goals and invariants of the system
    - some of which are OK to change
    - some of which should never be changed without a big discussion
- they are guard-rails to prevent accidental changes

Testy is designed to make it easier to write tests for all of these different purposes. By making it easier to write tests, more tests get written. By giving the author more control over when to halt the test (checks vs. asserts, structuring helpers), debugging failing tests becomes easier.

Explicitly, these other libraries have the following problems:

- `testify`
    - Has a massive surface area, so many methods make it confusing to know which one to use
    - Not typesafe, and most likely never will be (v2 seems abandoned)
    - Uses `reflect.DeepEquals` instead of `go-cmp`
- `gotest.tools`
    - Not typesafe (although may be soon?)
    - Makes `Equal` non-deep by default
    - Does not provide a soft/check version of `Equal` and `DeepEqual`
    - Magic implementation using ast-walking to determine comparison types is very cool but hard to understand
- `is`
    - Missing common and useful methods like `NotEqual`.
    - Magic implementation using ast-walking to determine comment messages is very cool but hard to understand

Only Testy is typesafe and able to perform both soft and hard style assertions,
with a reasonable API surface area, with `go-cmp`-powered deep equality by
default.

## Why not add more helper methods like testify?
When I was working on a real-life, multi-year, multi-developer project,
I regularly heard that testify was confusing because it wasn't clear which
methods to use when writing a test. I did some grepping/analysis, and found that
most tests were easily expressed using `Error/NoError`, `Equal/NotEqual`,
`True/False`, `Nil/NotNil`, `Zero/NotZero`, `Empty/NotEmpty`.
Testy handles all of these cases gracefully with a much reduced API surface area.
Hopefully this means it is easier to learn and use.

## What should I do if I rely on testify methods that aren't present here?
You have the following options:

- Rewrite the test to not use those methods. This sounds scary but is often not
  that hard.
- Write your own helper method for your specific check. Most codebases have
  their own test helpers anyway because it makes the tests more fluent to have
  domain-specific helpers.
- File an issue or PR to add the helper to this library. I'm open to the idea of
  adding a few more methods.

## I wrote custom test helpers like you recommended, but now they show up in the testing output and ruin the stacktrace. How can I avoid this?
In any testing helpers you create, just call `t.Helper()` to exclude them from the stacktrace. This is what Testy does (look at the code!)

```go
func myTestHelper(t *testing.T, otherArgs ...any) {
    t.Helper() // <- excludes this function from the stacktraces reported during test failures
    // ... actual logic goes here
}
```

## Why use `go-cmp` instead of `reflect.DeepEquals`?
There were a bunch of github issues about it being better, and in general it seems to give developers more control over how the comparison is implemented, including which fields to include/ignore. This seems to make a big difference particularly when comparing `time.Time` objects. For more information, see these discussions:

- https://github.com/stretchr/testify/issues/535
- https://github.com/matryer/is/issues/53

## How did you decide on "check" and "assert"?

Most languages have an "assert" concept that halts program execution if the
condition being asserted fails. So it makes sense to me to call them asserts.

Not many languages have the ability to easily perform soft asserts. One of the nice
things about Go is that it does, with the `t.Fail()`/`t.Error()` methods of its builtin `testing` library/framework/tool. There is no one name that everyone uses for this, but "check" makes sense to me and it's also used by `gotest.tools`.

`testify`.... ugh! It actually did the most, in my opinion, to popularize the use of both soft and hard style asserts. But it calls the soft style "asserts" and the hard style "requires". Their minds! What were they thinking! 😤

## How can I contribute?

Testy is a standard golang project, you'll need a working golang environment.

If you use VSCode, the repo comes with suggested extensions and settings.

Testing and linting scripts are defined with Just, see the Justfile to see how
to run those commands manually. There are also Github Actions that will lint and test
your PRs.

Contributions are more than welcome!
