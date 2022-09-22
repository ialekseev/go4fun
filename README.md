Go4Fun - GO for FUNctional programming
======================================

`Option`, `Sequence`, `Tuple` types with familiar functions found in other functional-first languages: `Map`, `FlatMap`, `Filter`, `Fold`, `Reduce`, `Zip`, `UnZip`... Alongside many other handy functions.

## Examples

### Map & FlatMap
```go
r = Some("route").Map(func(a string) string { return a + "60" })
fmt.Println(r)
// Output: Some(route60)

r = Some("route").FlatMap(func(a string) Option[string] { return Some(a + "60") })
fmt.Println(r)
// Output: Some(route60)

r = Seq[string]{"a", "b", "c"}.Map(func(a string) string { return a + "!" })
fmt.Println(r)
// Output: [a! b! c!]

r = Seq[int]{1, 2}.FlatMap(func(a int) Seq[int] { return Seq[int]{a, a} })
fmt.Println(r)
// Output: [1 1 2 2]
```

### Filter
```go
r = Some(5).Filter(func(a int) bool { return a < 10 })
fmt.Println(r)
// Output: Some(5)

r = Some(10).Filter(func(a int) bool { return a > 10 })
fmt.Println(r)
// Output: None

r = Seq[int]{2, 3, 4, 5, 6}.Filter(func(a int) bool { return a%2 == 0 })
fmt.Println(r)
// Output: [2 4 6]
```

### Fold & Reduce
```go
r = Some(5).Fold(1, func(a int) int { return a * 2 })
fmt.Println(r)
// Output: 10

r = None[int]().Fold(1, func(a int) int { return a * 2 })
fmt.Println(r)
// Output: 1

r = Seq[string]{"r", "o", "b"}.Fold("hi ", func(a1, a2 string) string { return a1 + a2 })
fmt.Println(r)
// Output: hi rob

r = Seq[int]{1, 2, 3, 4}.Reduce(func(a1, a2 int) int { return a1 + a2 })
fmt.Println(r)
// Output: 10
```

### Zip & UnZip
```go
r = ZipOption(Some("route"), Some(60))
fmt.Println(r)
// Output: Some((route,60))

r = UnZipOption(Some(Tup2("route", 60)))
fmt.Println(r)
// Output: (Some(route),Some(60))

r = ZipSeq(Seq[int]{1, 2, 3}, Seq[string]{"a", "b", "c"})
fmt.Println(r)
// Output: [(1,a) (2,b) (3,c)]

r = UnZipSeq(Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b"), Tup2(3, "c")})
fmt.Println(r)
// Output: ([1 2 3],[a b c])
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