Go4Fun - GO for FUNctional programming
======================================

`Option`, `Sequence`, `Either`, `Tuple` types with familiar functions found in other functional-first languages: `Map`, `FlatMap`, `Filter`, `Fold`, `Reduce`, `Zip`, `UnZip`... Alongside many other handy functions.

## Examples

### Option
```go
r = Some("route").Map(func(a string) string { return a + "60" })
fmt.Println(r)
// Output: Some(route60)

r = Some("route").FlatMap(func(a string) Option[string] { return Some(a + "60") })
fmt.Println(r)
// Output: Some(route60)

r = Some("route").FlatMap(func(a string) Option[string] { return None[string]() })
fmt.Println(r)
// Output: None

r = Some(5).Filter(func(a int) bool { return a < 10 })
fmt.Println(r)
// Output: Some(5)

r = Some(10).Filter(func(a int) bool { return a > 10 })
fmt.Println(r)
// Output: None

r = Some(5).Fold(1, func(a int) int { return a * 2 })
fmt.Println(r)
// Output: 10

r = None[int]().Fold(1, func(a int) int { return a * 2 })
fmt.Println(r)
// Output: 1

r = ZipOption(Some("route"), Some(60))
fmt.Println(r)
// Output: Some((route,60))

r = UnZipOption(Some(Tup2("route", 60)))
fmt.Println(r)
// Output: (Some(route),Some(60))
```

### Sequence
```go
r = Seq[string]{"a", "b", "c"}.Map(func(a string) string { return a + "!" })
fmt.Println(r)
// Output: [a! b! c!]

r = Seq[int]{1, 2}.FlatMap(func(a int) Seq[int] { return Seq[int]{a, a} })
fmt.Println(r)
// Output: [1 1 2 2]

r = Seq[int]{2, 3, 4, 5, 6}.Filter(func(a int) bool { return a%2 == 0 })
fmt.Println(r)
// Output: [2 4 6]

r = Seq[string]{"r", "o", "b"}.Fold("hi ", func(a1, a2 string) string { return a1 + a2 })
fmt.Println(r)
// Output: hi rob

r = Seq[int]{1, 2, 3, 4}.Reduce(func(a1, a2 int) int { return a1 + a2 })
fmt.Println(r)
// Output: 10

r = ZipSeq(Seq[int]{1, 2, 3}, Seq[string]{"a", "b", "c"})
fmt.Println(r)
// Output: [(1,a) (2,b) (3,c)]

r = UnZipSeq(Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b"), Tup2(3, "c")})
fmt.Println(r)
// Output: ([1 2 3],[a b c])
```

### Either
```go
r = Right[int]("60").Map(func(r string) string { return "route" + r })
fmt.Println(r)
// Output: Right(route60)

r = Right[int]("60").FlatMap(func(r string) Either[int, string] { return Right[int]("route" + r) })
fmt.Println(r)
// Output: Right(route60)

r = Right[int]("60").FlatMap(func(r string) Either[int, string] { return Left[int, string](-1) })
fmt.Println(r)
// Output: Left(-1)

r = Right[int]("john lennon").ToOption()
fmt.Println(r)
// Output: Some(john lennon)

r = Left[int, string](-1).ToOption()
fmt.Println(r)
// Output: None
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