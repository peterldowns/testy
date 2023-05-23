🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧

> **Warning**
> Under Construction, Work In Progress

🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧 🚧

# 😤 Testy

Testy is a library for writing meaningful and readable tests.
Testy is typesafe (using generics), based on [go-cmp](#), and designed as an alternative
to [testify](https://github.com/stretchr/testify), [gotools.test](https://github.com/gotestyourself/gotest.tools), and [is](https://github.com/matryer/is).

Major features:
- Typesafe, using generics, means that many failures are now caught at compiler time,
  rather than requiring you to run the test suite.
- Deep equality testing by default using [go-cmp](https://github.com/google/go-cmp), which
  means it works just like you'd expect, and gives you optional control if you want something different.
- Soft (calls `t.Fail()`) and hard (calls `t.FailNow()`) versions of every check method gives full control over which checks are run, in which stages, during your tests.
- Additional nice-to-have helpers for structuring tests in more readable ways.

```go
package example_test

import (
    "github.com/peterldowns/testy/check"
    "github.com/peterldowns/testy/assert"
)
func TestExample(t *testing.T) {
    // If a given check fails, the test will be marked as failed but continue
    // executing.  All failures are reported when the test stops executing,
    // either at the end of the test or when someone calls t.FailNow().
    check.True(t, true)
    check.False(t, false)
    check.Equal(t, []string{"hello"}, []string{"hello"})
    check.StrictEqual(t, "hello", "hello")
    check.NotEqual(t, map[string]int{"hello": 1}, nil)
    check.StrictNotEqual(t, 5, 0)
    check.LessThan(t, 1, 4)
    check.LessThanOrEquals(t, 4, 4)
    check.GreaterThan(t, 8, 6)
    check.GreaterThanOrequals(t, 6, 6)
    check.Error(t, fmt.Errorf("oh no!")
    check.Nil(t, nil)
    // If a given assert fails, the test will immediately be marked as failed
    // stop executing, and report all failures.
    assert.True(t, true)
    assert.False(t, false)
    assert.Equal(t, []string{"hello"}, []string{"hello"})
    assert.StrictEqual(t, "hello", "hello")
    assert.NotEqual(t, map[string]int{"hello": 1}, nil)
    assert.StrictNotEqual(t, 5, 0)
    assert.LessThan(t, 1, 4)
    assert.LessThanOrEquals(t, 4, 4)
    assert.GreaterThan(t, 8, 6)
    assert.GreaterThanOrequals(t, 6, 6)
    assert.Error(t, fmt.Errorf("oh no!")
    assert.Nil(t, nil)
}
```

## motivation
Testy helps you write more usable, readable, and meaningful tests. By allowing
you to determine between *checks* (soft assertions) and *asserts* (hard
assertions), and providing structuring helpers, you can make debugging much
easier on your teammates and your future self.

Testy attempts to strike a balance between allowing explicit and meaningful
tests, encouraging standard golang coding conventions, avoiding implementation
magic, and minimizing the number of check methods to choose between when writing
tests.  Here's everything you need to know:
## `check`
`check` contains methods for checking a condition, marking a test failed if
the condition is not met.  This is a "soft" style assert, equivalent to the
methods in `testify/assert` or the `Check` method in `gotest.tools/assert`.

Each `check` method returns a boolean, which is `true` if the check passed, and `false` otherwise. You can use this to conditionally run other logic in your code.

```go
func TestExample(t *testing.T) {
    var f *MyFoo
    var err error
    f, err = ServiceThatGetsAFoo()
    // f is only meaningful if err == nil
    if check.Nil(t, err) {
        check.Equals(f.Name, "peter")
    }
}
```

## `assert`
`assert` contains methods for checking a condition, failing the test
immediately if the condition is not met. This is a "hard" or traditional
assert, equivalent to the methods in `testify/require` or the `Assert` method
in `gotest.tools/assert`.

```go
func TestExample(t *testing.T) {
    var f *MyFoo
    var err error
    f, err = ServiceThatGetsAFoo()
    assert.Nil(t, err) // if err != nil, the test will end here
    assert.Equals(f.Name, "peter")
}
```

## methods
The following methods are available on both `check` and `assert`:

- `True(t, x)` checks if its argument is `true`
- `False(t, x)` checks if its argument is `false`
- `Equal(t, want, got)` checks if its arguments are equal using [go-cmp](https://github.com/google/go-cmp)
- `StrictEqual(t, want, got)` checks if its arguments are equal using `==`
- `NotEqual(t, want, got)` checks if its arguments are not equal using [go-cmp](https://github.com/google/go-cmp)
- `StrictNotEqual(t, want, got)` checks if its arguments are not equal using `!=`
- `LessThan(t, small, big)` checks if `small < big`
- `LessThanOrEquals(t, small, big)` checks if `small <= big`
- `GreaterThan(t, big, small)` checks if `big > small`
- `GreaterThanOrEquals(t, big, small)` checks if `big >= small`
- `Error(t, err)` checks if `err != nil`
- `Nil(t, err)` checks if `err == nil`


## structuring helpers
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

```go
func TestStructuringHelpers(t *testing.T) {
    // Here's an easy way to "stage" your tests, where
    // - each stage runs all of the checks
    // - at the end, report any failures.
    // - if there were any failures, end the test.
    check.Equal(t, 2, 2)
    check.LessThanOrEquals(t, 2, 3)
    check.GreaterThan(t, 3, 1)
    assert.NoFailures(t)

    // This is another, equivalent, way to express the same logic
    assert.Group(t, func() error {
        check.Equal(t, 2, 2)
        check.LessThanOrEquals(t, 2, 3)
        check.GreaterThan(t, 3, 1)
        return nil
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
    assert.Nil(err)
    _, err = anotherF(bar)
    assert.Nil(err)
}
```

## more examples

Beyond the examples presented in this README, see `main_test.go` for additional examples.


# FAQ

## why create this library? aren't the existing libraries way better and more polished?
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

Most tests are run and refactored many, many more times than they are written. I believe that
means we should primarily write tests to serve their readers and future users. Testy is an improvement over testify/gotest.tools/is in the following ways:

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


## why not add more helper methods like testify?
When I was working on a real-life, multi-year, multi-developer project,
I regularly heard that testify was confusing because it wasn't clear which
methods to use when writing a test. I did some grepping/analysis, and found that
most tests were easily expressed using `Error/NoError`, `Equal/NotEqual`,
`True/False`, `Nil/NotNil`, `Error/NoError`, `Zero/NotZero`, `Empty/NotEmpty`.
Testy handles all of these cases gracefully with a much reduced surface area.
Hopefully this means it is easier to learn and use.

## what should I do if I rely on testify methods that aren't present here?
First off, sorry, I know they do make life more convenient. My recommendation is
you should change how your test is expressed, or reimplement the helper method
yourself.

Most codebases have their own domain-specific helpers, and this is a Good Thing.
Tests should be written like normal code, and involve plenty of helper methods
and functions for making it easy to express your domain specific problems.

That said, I could be convinced to create a big library of these, like
`gotest.tools` has done, if enough people think that's the move. Let me know by
opening up an issue/PR or contacting me via email.

## i wrote custom test helpers like you recommended, but now they show up in the testing output and ruin the stacktrace. how can i avoid this?
In any testing helpers you create, just call `t.Helper()` to exclude them from the stacktrace. This is what Testy does (look at the code!)

```go
func myTestHelper(t *testing.T, otherArgs ...any) {
    t.Helper() // <- excludes this function from the stacktraces reported during test failures
    // ... actual logic goes here
}
```

## why use `go-cmp` instead of `reflect.DeepEquals`?
There were a bunch of github issues about it being better, and in general it seems to give developers more control over how the comparison is implemented, including which fields to include/ignore. This seems to make a big difference particularly when comparing `time.Time` objects. For more information, see these discussions:

- https://github.com/stretchr/testify/issues/535
- https://github.com/matryer/is/issues/53

## how did you decide on "check" and "assert"?

Most languages have an "assert" concept that halts program execution if the
condition being asserted fails. So it makes sense to me to call them asserts.

Not many languages have the ability to easily perform soft asserts. One of the nice
things about Go is that it does, with the `t.Fail()`/`t.Error()` methods of its builtin `testing` library/framework/tool. There is no one name that everyone uses for this, but "check" makes sense to me and it's also used by `gotest.tools`.

`testify`.... ugh! It actually did the most, in my opinion, to popularize the use of both soft and hard style asserts. But it calls the soft style "asserts" and the hard style "requires". Their minds! What were they thinking! 😤