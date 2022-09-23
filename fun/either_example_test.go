package fun

import (
	"fmt"
)

// Examples for: Map, FlatMap, ToOption

func ExampleEither_Map_eg1() {
	r := Right[int]("60").Map(func(r string) string { return "route" + r })
	fmt.Println(r)
	// Output: Right(route60)
}

func ExampleEither_Map_eg2() {
	r := Left[int, string](-1).Map(func(r string) string { return "route" + r })
	fmt.Println(r)
	// Output: Left(-1)
}

func ExampleMapEither_eg1() {
	r := MapEither(Right[bool](60), func(r int) string { return "route" + fmt.Sprint(r) })
	fmt.Println(r)
	// Output: Right(route60)
}

func ExampleMapEither_eg2() {
	r := MapEither(Left[bool, int](false), func(r int) string { return "route" + fmt.Sprint(r) })
	fmt.Println(r)
	// Output: Left(false)
}

func ExampleEither_FlatMap_eg1() {
	r := Right[int]("60").FlatMap(func(r string) Either[int, string] { return Right[int]("route" + r) })
	fmt.Println(r)
	// Output: Right(route60)
}

func ExampleEither_FlatMap_eg2() {
	r := Right[int]("60").FlatMap(func(r string) Either[int, string] { return Left[int, string](-1) })
	fmt.Println(r)
	// Output: Left(-1)
}

func ExampleEither_FlatMap_eg3() {
	r := Left[int, string](-1).FlatMap(func(r string) Either[int, string] { return Right[int]("route" + r) })
	fmt.Println(r)
	// Output: Left(-1)
}

func ExampleFlatMapEither_eg1() {
	r := FlatMapEither(Right[bool](60), func(r int) Either[bool, string] { return Right[bool]("route" + fmt.Sprint(r)) })
	fmt.Println(r)
	// Output: Right(route60)
}

func ExampleFlatMapEither_eg2() {
	r := FlatMapEither(Right[bool](60), func(r int) Either[bool, string] { return Left[bool, string](false) })
	fmt.Println(r)
	// Output: Left(false)
}

func ExampleFlatMapEither_eg3() {
	r := FlatMapEither(Left[bool, int](false), func(r int) Either[bool, string] { return Right[bool]("route" + fmt.Sprint(r)) })
	fmt.Println(r)
	// Output: Left(false)
}

func ExampleEither_ToOption_eg1() {
	r := Right[int]("john lennon").ToOption()
	fmt.Println(r)
	// Output: Some(john lennon)
}

func ExampleEither_ToOption_eg2() {
	r := Left[int, string](-1).ToOption()
	fmt.Println(r)
	// Output: None
}
