Go4Fun - GO for FUNctional programming
======================================
![build-test](https://github.com/ialekseev/go4fun/actions/workflows/main.yml/badge.svg)

`Option`, `Sequence`, `Future`, `Either`, `Tuple` types with familiar combinators found in other functional-first languages: `Map`, `FlatMap`, `Apply (Applicative)`, `Filter`, `Fold`, `Reduce`, `Zip`, `UnZip`... alongside many other handy functions. And also: `Trampoline`, `Currying`, `Function Composition`...

# Examples
- [Option](https://github.com/ialekseev/go4fun#option)
- [Sequence](https://github.com/ialekseev/go4fun#sequence)
- [Future](https://github.com/ialekseev/go4fun#future)
- [Either](https://github.com/ialekseev/go4fun#either)
- [Trampoline](https://github.com/ialekseev/go4fun#trampoline)
- [Currying](https://github.com/ialekseev/go4fun#currying)
- [Function Composition](https://github.com/ialekseev/go4fun#function-composition)

## Option
#### Map
```go
r = Some("route").Map(func(a string) string { return a + "60" })
fmt.Println(r)
// Output: Some(route60)
```
#### FlatMap
```go
r = Some("route").FlatMap(func(a string) Option[string] { return Some(a + "60") })
fmt.Println(r)
// Output: Some(route60)

r = Some("route").FlatMap(func(a string) Option[string] { return None[string]() })
fmt.Println(r)
// Output: None
```
#### Apply (Applicative)
```go
r = ApplyOption3(Some(true), Some(10), Some("abc"), func(a bool, b int, c string) Option[string] {
	return Some(fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c))
})
fmt.Println(r)
// Output: Some(true 10 abc)

r = ApplyOption3(None[bool](), Some(10), Some("abc"), func(a bool, b int, c string) Option[string] {
	return Some(fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c))
})
fmt.Println(r)
// Output: None
```
#### Filter
```go
r = Some(5).Filter(func(a int) bool { return a < 10 })
fmt.Println(r)
// Output: Some(5)

r = Some(10).Filter(func(a int) bool { return a > 10 })
fmt.Println(r)
// Output: None
```
#### Fold
```go
r = Some(5).Fold(1, func(a int) int { return a * 2 })
fmt.Println(r)
// Output: 10

r = None[int]().Fold(1, func(a int) int { return a * 2 })
fmt.Println(r)
// Output: 1
```
#### Zip & UnZip
```go
r = ZipOption(Some("route"), Some(60))
fmt.Println(r)
// Output: Some((route,60))

r = UnZipOption(Some(Tup2("route", 60)))
fmt.Println(r)
// Output: (Some(route),Some(60))
```

## Sequence
#### Map
```go
r = Seq[string]{"a", "b", "c"}.Map(func(a string) string { return a + "!" })
fmt.Println(r)
// Output: [a! b! c!]
```
#### FlatMap
```go
r = Seq[int]{1, 2}.FlatMap(func(a int) Seq[int] { return Seq[int]{a, a} })
fmt.Println(r)
// Output: [1 1 2 2]
```
#### Filter
```go
r = Seq[int]{2, 3, 4, 5, 6}.Filter(func(a int) bool { return a%2 == 0 })
fmt.Println(r)
// Output: [2 4 6]
```
#### Fold
```go
r = Seq[string]{"r", "o", "b"}.Fold("hi ", func(a1, a2 string) string { return a1 + a2 })
fmt.Println(r)
// Output: hi rob
```
#### Reduce
```go
r = Seq[int]{1, 2, 3, 4}.Reduce(func(a1, a2 int) int { return a1 + a2 })
fmt.Println(r)
// Output: 10
```
#### Zip & UnZip
```go
r = ZipSeq(Seq[int]{1, 2, 3}, Seq[string]{"a", "b", "c"})
fmt.Println(r)
// Output: [(1,a) (2,b) (3,c)]

r = UnZipSeq(Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b"), Tup2(3, "c")})
fmt.Println(r)
// Output: ([1 2 3],[a b c])
```

## Future
#### Map
```go
future = FutureValue(func() string {
	time.Sleep(time.Millisecond * 20)
	return "abc"
})

r = future.Map(func(a string) string { return a + "def" })

time.Sleep(time.Millisecond * 30)

fmt.Println(r.Result())
// Output: abcdef
```
#### FlatMap
```go
future = FutureValue(func() string {
	time.Sleep(time.Millisecond * 10)
	return "abc"
})

r = future.FlatMap(func(a string) Future[string] {
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
future1 = FutureValue(func() int {
	time.Sleep(time.Millisecond * 10)
	return 123
})

future2 = FutureValue(func() bool {
	time.Sleep(time.Millisecond * 10)
	return true
})

r = ApplyFuture2(future1, future2, func(a int, b bool) Future[string] {
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
future = FutureValue(func() string {
	time.Sleep(time.Millisecond * 20)
	return "abc"
})

future.OnComplete(func(a string) { fmt.Println(a + "def") })

time.Sleep(time.Millisecond * 30)

// Output: abcdef
```

## Either
#### Map
```go
r = Right[int]("60").Map(func(r string) string { return "route" + r })
fmt.Println(r)
// Output: Right(route60)
```
#### FlatMap
```go
r = Right[int]("60").FlatMap(func(r string) Either[int, string] { return Right[int]("route" + r) })
fmt.Println(r)
// Output: Right(route60)

r = Right[int]("60").FlatMap(func(r string) Either[int, string] { return Left[int, string](-1) })
fmt.Println(r)
// Output: Left(-1)
```
#### ToOption
```go
r = Right[int]("john lennon").ToOption()
fmt.Println(r)
// Output: Some(john lennon)

r = Left[int, string](-1).ToOption()
fmt.Println(r)
// Output: None
```

## Trampoline
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
#### Curry
```go
f = func(a int, b bool, c float64) string {
	return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c)
}

curriedF := Curry3(f)
r = curriedF(1)(true)(5.5)

fmt.Println(r)
// Output: 1 true 5.5
```
#### UnCurry
```go
f = func(a int) func(bool) func(float64) string {
	return func(b bool) func(float64) string {
		return func(c float64) string {
			return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c)
		}
	}
}

unCurriedF := UnCurry3(f)
r = unCurriedF(1, true, 5.5)

fmt.Println(r)
// Output: 1 true 5.5
```

## Function Composition
```go
f = func(a int) string { return fmt.Sprint(a) }
g = func(b string) bool { return b != "" }
h = func(c bool) string { return fmt.Sprint(c) }

j = Compose3(f, g, h)

fmt.Println(j(1) == h(g(f(1))))
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