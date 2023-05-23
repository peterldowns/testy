> **Warning**
> 🚧 Under Construction, Work In Progress 🚧

# 😤 Testy

Testy is a library for writing meaningful and readable tests.
Testy is typesafe (using generics), based on [go-cmp](#), and designed as an alternative
to [testify](https://github.com/stretchr/testify), [gotools.test](https://github.com/gotestyourself/gotest.tools), and [is](https://github.com/matryer/is).

Major features:
- typesafe using generics
- deep equality testing by default using [go-cmp](#)
- "soft" (calls `t.Fail()`) and "hard" (calls `t.FailNow()`) versions of every check method
- helpers for structuring tests

Developers working on complicated systems rely on tests to help educate them
about the code under test, often even more-so than the related documentation.
Tests are more useful when they convey more-specific information about the
system under test, especially when that system's behavior changes and causes
the tests to fail. Tests are code, too, but most developers seem willing to accept
substandard and inexpressive systems.

Testy attempts to strike a balance between allowing explicit and meaningful tests,
encouraging standard coding conventions, reducing implementation magic, and
minimizing the number of methods to learn. Here's everything you need to know:
### `check`
`check` contains methods for checking a condition, marking a test failed if
the condition is not met.  This is a "soft" style assert, equivalent to the
methods in `testify/assert` or the `Check` method in `gotest.tools/assert`.

### `assert`
`assert` contains methods for checking a condition, failing the test
immediately if the condition is not met. This is a "hard" or traditional
assert, equivalent to the methods in `testify/require` or the `Assert` method
in `gotest.tools/assert`.

### standard methods
The following methods are available on both `check` and `assert`:
- `True(t, x)` checks if its argument is `true`
- `False(t, x)` checks if its argument is `false`
- `Equal(t, want, got)` checks if its arguments are equal using [go-cmp](#)
- `StrictEqual(t, want, got)` checks if its arguments are equal using `==`
- `NotEqual(t, want, got)` checks if its arguments are not equal using [go-cmp](#)
- `StrictNotEqual(t, want, got)` checks if its arguments are not equal using `!=`
- `LessThan(t, small, big)` checks if `small < big`
- `LessThanOrEquals(t, small, big)` checks if `small <= big`
- `GreaterThan(t, big, small)` checks if `big > small`
- `GreaterThanOrEquals(t, big, small)` checks if `big >= small`
- `Error(t, err)` checks if `err != nil`
- `Nil(t, err)` checks if `err == nil`

### structuring helpers
The `assert` package also provides helpers for structuring your tests and making them more expressive. You can use these helpers to determine which checks are run in parallel, and which checks should halt test execution.

- `assert.NoFailures(t)` will instantly fail the test if any previous `check` has failed,
or any other code has called `t.Fail()`/`t.Error()` for any reason.
- `assert.NoFailures(t, thunk...(func() error))` will do the following:
    - for each thunk:
        - fail immediately if any previous `check` has failed.
        - execute the thunk
        - fail immediately if the thunk returns a non-nil error
    - fail immediately if check has failed
- `assert.NoErrors == assert.Group == assert.NoFailures` &mdash; pick the name that makes most sense in context.

# examples

See `main_test.go` for a full series of runnable and working examples. Here are a few excerpts that should serve to explain why this library exists:

```go
package main_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/peterldowns/testy/assert"
	"github.com/peterldowns/testy/check"
)

func TestEquality(t *testing.T) {
	// These two Time objects are equal due to the logic in their .Equal()
	// method, but they are not strictly equal because the structs contain
	// fields with different values.
	t1 := time.Now()
	t2 := t1.UTC()
	assert.Equal(t, t1, t2)
	assert.StrictNotEqual(t, t1, t2)
}

func TestStructuringHelpers(t *testing.T) {
    // Here's an easy way to "stage" your tests, where
    // - each stage runs all of the checks
    // - if any check fails, don't go on to the next stage.

    // stage 1
    check.Equal(t, 2, 2)
    check.LessThanOrEquals(t, 2, 3)
    check.GreaterThan(t, 3, 1)
    assert.NoFailures(t)

    // This is another, equivalent, way to express a stage
    // stage 2
    assert.Group(t, func() error {
        check.Equal(t, 2, 2)
        check.LessThanOrEquals(t, 2, 3)
        check.GreaterThan(t, 3, 1)
        return nil
    })
```



# FAQ

## why create this library? aren't the existing libraries way better and more polished?
Each of the existing libraries (testify, gotest.tools, is) are phenomenal projects and Testy
has been deeply influenced by each of them, and could not exist without them. While writing Testy, I used those libraries as reference, and I have great respect for their authors and contributors.

That said, these libraries are (in my opinion) too limited in terms of expressiveness. Tests serve multiple purposes:
- ground-level correctness of the system
- proof that a bug was fixed 
    - therefore, as a historical record of all discovered bugs
- documentation of the current implementation of the system and how to use its components
- documentation of the business goals and invariants of the system
    - some of which are OK to change
    - some of which should never be changed without a big discussion
- guard-rails to prevent accidental changes

Most tests are run and refactored many, many more times than they are written. I believe that
means we should primarily write tests to serve their readers and future users. Testy is an improvement over testify/gotest.tools/is in the following ways:

- Typesafe testy/assert methods make it easier to write and refactor tests without having to run them, "shifting left" the detection of problems. This is generally agreed to be a good thing.
- By default, equality testing is deep and uses [go-cmp](#), because most of the time you just want to check if two things are equal, and it should work correctly.
- The test-structuring helpers and the "soft" (check) and "hard" (assert) version of every test method means that you can write tests that are easier to understand when debugging.

Specifically:
- Testify
    - Has a massive surface area, so many methods make it confusing to know which one to use
    - Not typesafe
    - Uses `reflect.DeepEquals`
- gotest.tools
    - Not typesafe (although may be soon?)
    - Does not provide a "soft" method of `Equal` and `DeepEqual`
    - Makes `Equal` non-deep by default
    - Magic implementation using ast-walking to determine comparison types is very cool but hard to understand
- is
    - Magic implementation using ast-walking to determine comment messages is very cool but hard to understand
    - Missing common and useful methods like `NotEqual`.



## why not more helper methods like testify?
Based on a real-life, multi-year, multi-developer project I was part of,
it seemed like the most important methods were `Error/NoError`, `Equal/NotEqual`, `True/False`, `Nil/NotNil`, `Error/NoError`, `Zero/NotZero`, `Empty/NotEmpty`. Testy handles all of these cases gracefully with a much reduced surface area. 

## what should I do if I rely on testify methods that aren't present here?
First off, sorry, I know they do make life more convenient. My recommendation is you should change how your test is expressed, or reimplement the helper method yourself. Most codebases
have their own domain-specific helpers anyway. I could be convinced to create a big library of these, like `gotest.tools` has done, if enough people think that's the move. let me know by opening up an issue/PR or contacting me via email.

## my custom helpers show up in the testing output and ruin the stacktrace, how can I avoid that?
In any testing helpers you create, just call `t.Helper()` to exclude them from the stacktrace. This is what Testy does.

## why use `go-cmp` instead of `reflect.DeepEquals`?
There were a bunch of github issues about it being better, and in general it seems to give developers more control over how the comparison is implemented, including which fields to include/ignore. This seems to make a big difference particularly when comparing `time.Time` objects. For more information, see these discussions:

- https://github.com/stretchr/testify/issues/535
- https://github.com/matryer/is/issues/53
