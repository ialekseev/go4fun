package fun

import (
	"fmt"
)

// Examples for: Map, FlatMap, Filter, Fold, Reduce, Zip, UnZip

func ExampleSeq_Map_eg1() {
	r := Seq[string]{"a", "b", "c"}.Map(func(a string) string { return a + "!" })
	fmt.Println(r)
	// Output: [a! b! c!]
}

func ExampleMapSeq_eg1() {
	r := MapSeq(Seq[int]{1, 2, 3}, func(a int) string { return "tick" + fmt.Sprint(a) })
	fmt.Println(r)
	// Output: [tick1 tick2 tick3]
}

func ExampleSeq_FlatMap_eg1() {
	r := Seq[int]{1, 2}.FlatMap(func(a int) Seq[int] { return Seq[int]{a, a} })
	fmt.Println(r)
	// Output: [1 1 2 2]
}

func ExampleFlatMapSeq_eg1() {
	r := FlatMapSeq(Seq[int]{1, 2}, func(a int) Seq[string] { return Seq[string]{"tick" + fmt.Sprint(a), "tack" + fmt.Sprint(a)} })
	fmt.Println(r)
	// Output: [tick1 tack1 tick2 tack2]
}
