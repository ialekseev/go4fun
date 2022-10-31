Go4Fun - GO for FUNctional programming
======================================
![build-test](https://github.com/ialekseev/go4fun/actions/workflows/main.yml/badge.svg)

`Option`, `Sequence`, `Lazy Sequence`, `Future`, `Either`, `Tuple` types with familiar combinators found in other functional-first languages: `Map`, `FlatMap`, `Apply (Applicative)`, `Filter`, `Fold`, `Reduce`, `Zip`, `UnZip`... alongside many other handy functions. And also: `Memoization`, `Trampoline`, `Currying`, `Function Composition`...

# Examples
- [Option](https://github.com/ialekseev/go4fun#option)
- [Sequence](https://github.com/ialekseev/go4fun#sequence)
- [Lazy Sequence](https://github.com/ialekseev/go4fun#lazy-sequence)
- [Future](https://github.com/ialekseev/go4fun#future)
- [Either](https://github.com/ialekseev/go4fun#either)
- [Memoization](https://github.com/ialekseev/go4fun#memoization)
- [Trampoline](https://github.com/ialekseev/go4fun#trampoline)
- [Currying](https://github.com/ialekseev/go4fun#currying)
- [Function Composition](https://github.com/ialekseev/go4fun#function-composition)

## Option
`Option` type represents optional values. If a value exists - it is wrapped as `Some(value)`. If not - it is `None`. It is a FP way of representing a (possibly) non-existing value. Functions like `Map`, `FlatMap`, `Apply`, `Filter` etc. allow complex chaining of `Option` values without having to check for the existence of a value.
#### Map
```go
r := Some("route").Map(func(a string) string { return a + "60" })
fmt.Println(r)
// Output: Some(route60)
```
#### FlatMap
```go
r := Some("route").FlatMap(func(a string) Option[string] { return Some(a + "60") })
fmt.Println(r)
// Output: Some(route60)

r := Some("route").FlatMap(func(a string) Option[string] { return None[string]() })
fmt.Println(r)
// Output: None
```
#### Apply (Applicative)
```go
r := ApplyOption3(Some(true), Some(10), Some("abc"), func(a bool, b int, c string) Option[string] {
	return Some(fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c))
})
fmt.Println(r)
// Output: Some(true 10 abc)

r := ApplyOption3(None[bool](), Some(10), Some("abc"), func(a bool, b int, c string) Option[string] {
	return Some(fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c))
})
fmt.Println(r)
// Output: None
```
#### Filter
```go
r := Some(5).Filter(func(a int) bool { return a < 10 })
fmt.Println(r)
// Output: Some(5)

r := Some(10).Filter(func(a int) bool { return a > 10 })
fmt.Println(r)
// Output: None
```
#### Fold
```go
r := Some(5).Fold(1, func(a int) int { return a * 2 })
fmt.Println(r)
// Output: 10

r := None[int]().Fold(1, func(a int) int { return a * 2 })
fmt.Println(r)
// Output: 1
```
#### Zip & UnZip
```go
r := ZipOption(Some("route"), Some(60))
fmt.Println(r)
// Output: Some((route,60))

r := UnZipOption(Some(Tup2("route", 60)))
fmt.Println(r)
// Output: (Some(route),Some(60))
```

## Sequence
`Sequence` type is based on Go slices with common FP functions added on top of that: `Map`, `FlatMap`, `Filter`, `Fold`, `Reduce`, `Zip`, `UnZip` etc.
#### Map
```go
r := Seq[string]{"a", "b", "c"}.Map(func(a string) string { return a + "!" })
fmt.Println(r)
// Output: [a! b! c!]
```
#### FlatMap
```go
r := Seq[int]{1, 2}.FlatMap(func(a int) Seq[int] { return Seq[int]{a, a} })
fmt.Println(r)
// Output: [1 1 2 2]
```
#### Filter
```go
r := Seq[int]{2, 3, 4, 5, 6}.Filter(func(a int) bool { return a%2 == 0 })
fmt.Println(r)
// Output: [2 4 6]
```
#### Fold
```go
r := Seq[string]{"r", "o", "b"}.Fold("hi ", func(a1, a2 string) string { return a1 + a2 })
fmt.Println(r)
// Output: hi rob
```
#### Reduce
```go
r := Seq[int]{1, 2, 3, 4}.Reduce(func(a1, a2 int) int { return a1 + a2 })
fmt.Println(r)
// Output: 10
```
#### Zip & UnZip
```go
r := ZipSeq(Seq[int]{1, 2, 3}, Seq[string]{"a", "b", "c"})
fmt.Println(r)
// Output: [(1,a) (2,b) (3,c)]

r := UnZipSeq(Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b"), Tup2(3, "c")})
fmt.Println(r)
// Output: ([1 2 3],[a b c])
```

## Lazy Sequence
`Lazy Sequence` iterates elements only when they are needed (unlike the regular `Sequence` that does it eagerly). It has the same functions as `Sequence` but many of them are "Lazy" and would not trigger any processing until that is needed.
```go
// Strict (Regular) Sequence eagerly iterates its elements on each operation.
// Below code calculates a result in multiple iterations:
r1 := Seq[int]{-2, -1, 0, 1, 2, 3, 4, 5, 6}.
	Filter(func(a int) bool { return a > 0 }).
	Filter(func(a int) bool { return a%2 == 0 }).
	Map(func(a int) int { return a / 2 }).
	Reduce(func(a1, a2 int) int { return a1 + a2 })

fmt.Println(r1)

// Lazy Sequence iterates elements only when they are needed.
// In this case, it's when the last materializing call happens (Reduce).
// Other calls (Filter, Map) are "lazy" and don't result in any computation.
// Below code calculates a result in 1 iteration:
r2 := Seq[int]{-2, -1, 0, 1, 2, 3, 4, 5, 6}.Lazy().
	Filter(func(a int) bool { return a > 0 }).
	Filter(func(a int) bool { return a%2 == 0 }).
	Map(func(a int) int { return a / 2 }).
	Reduce(func(a1, a2 int) int { return a1 + a2 })

fmt.Println(r2)

// Output:
// 6
// 6
```
```go
lazySeq := Seq[string]{"b", "c", "d", "e", "f"}.Lazy().
	Map(func(a string) string { return strings.ToUpper(a) })

// The same Lazy Sequence is re-used below for different computations:

r1 := lazySeq.FlatMap(func(a string) LazySeq[string] { return Seq[string]{a, a}.Lazy() }).Strict()
fmt.Println(r1)

r2 := lazySeq.Map(func(a string) string { return strings.ToUpper(a) }).
	Fold("A", func(a string, b string) string { return a + b })

fmt.Println(r2)
// Output:
// [B B C C D D E E F F]
// ABCDEF
```

## Future
`Future` represents a value that may not be yet available, but should become available at some point when the underlying asynchronous computation is completed. It also has functions (`Map`, `FlatMap`, `Apply`...) that allow chaining of `Future` values without having to check for the availability of a value.
#### Map
```go
future := FutureValue(func() string {
	time.Sleep(time.Millisecond * 20)
	return "abc"
})

r := future.Map(func(a string) string { return a + "def" })

time.Sleep(time.Millisecond * 30)

fmt.Println(r.Result())
// Output: abcdef
```
#### FlatMap
```go
future := FutureValue(func() string {
	time.Sleep(time.Millisecond * 10)
	return "abc"
})

r := future.FlatMap(func(a string) Future[string] {
	return FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return a + "def"
	})
})

time.Sleep(time.Millisecond * 50)

fmt.Println(r.Result())
// Output: abcdef
```
#### Apply (Applicative)
```go
future1 := FutureValue(func() int {
	time.Sleep(time.Millisecond * 10)
	return 123
})

future2 := FutureValue(func() bool {
	time.Sleep(time.Millisecond * 10)
	return true
})

r := ApplyFuture2(future1, future2, func(a int, b bool) Future[string] {
	return FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return fmt.Sprint(a) + " " + fmt.Sprint(b)
	})
})

time.Sleep(time.Millisecond * 60)

fmt.Println(r.Result())
// Output: 123 true
```
#### OnComplete
```go
future := FutureValue(func() string {
	time.Sleep(time.Millisecond * 20)
	return "abc"
})

future.OnComplete(func(a string) { fmt.Println(a + "def") })

time.Sleep(time.Millisecond * 30)

// Output: abcdef
```

## Either
`Either` represents a value of one of two possible cases: it is either `Left` or `Right`. A common use of `Either` is as an alternative to `Option` for dealing with (possibly) missing values. In this case, `Left` is used instead of `None` and can additionally contain useful information, and `Right` is used instead of `Some`. Thus, usually `Left` is used for a failure and `Right` for a success.
#### Map
```go
r := Right[int]("60").Map(func(r string) string { return "route" + r })
fmt.Println(r)
// Output: Right(route60)
```
#### FlatMap
```go
r := Right[int]("60").FlatMap(func(r string) Either[int, string] { return Right[int]("route" + r) })
fmt.Println(r)
// Output: Right(route60)

r := Right[int]("60").FlatMap(func(r string) Either[int, string] { return Left[int, string](-1) })
fmt.Println(r)
// Output: Left(-1)
```
#### ToOption
```go
r := Right[int]("john lennon").ToOption()
fmt.Println(r)
// Output: Some(john lennon)

r := Left[int, string](-1).ToOption()
fmt.Println(r)
// Output: None
```

## Memoization
Memoization is an optimization technique where an expensive function is wrapped into a memoized `Memo` function of the same signature. It would cache results of function calls and return back a cached result for the same input, if requested again.
```go
// an expensive function (with 1 argument) is wrapped into a memoized function of the same signature.
var memoF = Memo1(func(a int) string {
	// expensive computation:
	time.Sleep(time.Millisecond * time.Duration(a))
	return fmt.Sprint(a)
})

r := memoF(2) // the first call is slow
r = memoF(2)  // other calls are fast

fmt.Println(r)
// Output: 2
```
```go
// an expensive function (with 2 arguments) is wrapped into a memoized function of the same signature.
var memoF = Memo2(func(a, b int) string {
	// expensive computation:
	time.Sleep(time.Millisecond * time.Duration(a+b))
	return fmt.Sprint(a) + fmt.Sprint(b)
})

r := memoF(2, 2) // the first call is slow
r = memoF(2, 2)  // other calls are fast

fmt.Println(r)
// Output: 22
```

## Trampoline
Trampoline allows to preserve a recursive structure of the code while avoiding a possible Stack Overflow problem. A recursive function becomes a description of the recursive computation which we need to run later to actually produce a real result. `MoreTrampolining` call is used to wrap a deferred call and `DoneTrampolining` to wrap a final result.
```go
// Recursion without Trampoline:
func summation(n, current uint64) uint64 {
	if n < 1 {
		return current
	}
	return summation(n-1, n+current)
}
summation(100000000, 0)
// fatal error: stack overflow

// Recursion with Trampoline:
func summationT(n, current uint64) Trampoline[uint64] {
	if n < 1 {
		return DoneTrampolining(current)
	}
	return MoreTrampolining(func() Trampoline[uint64] {
		return summationT(n-1, n+current)
	})
}
summationT(100000000, 0).Run()
// Output: 5000000050000000
```

## Currying
Currying is a technique of converting a function that takes multiple arguments into a sequence of functions that each takes a single argument.
#### Curry
```go
f := func(a int, b bool, c float64) string {
	return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c)
}

curriedF := Curry3(f)
r := curriedF(1)(true)(5.5)

fmt.Println(r)
// Output: 1 true 5.5
```
#### UnCurry
```go
f := func(a int) func(bool) func(float64) string {
	return func(b bool) func(float64) string {
		return func(c float64) string {
			return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c)
		}
	}
}

unCurriedF := UnCurry3(f)
r := unCurriedF(1, true, 5.5)

fmt.Println(r)
// Output: 1 true 5.5
```

## Function Composition
Function composition is an operation `Compose2` that takes two functions `f` and `g`, and produces a function `h` such that `h(x) = g(f(x))`. The concept could be extended beyond `Compose2` to chain more than 2 functions: `Compose3` etc.
```go
f := func(a int) string { return fmt.Sprint(a) }
g := func(b string) bool { return b != "" }

h := Compose2(f, g)

fmt.Println(h(1) == g(f(1)))
// Output: true
```

Installation
============

To install Go4Fun use `go get`:

    go get github.com/ialekseev/go4fun


Import `fun` package:
```go
import (
  "github.com/ialekseev/go4fun/fun"
)
```

Staying up to date
==================

To update to the latest version use `go get -u github.com/ialekseev/go4fun`.


License
=======

This project is licensed under the terms of the MIT license.