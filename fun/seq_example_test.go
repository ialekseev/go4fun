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

func ExampleSeq_Filter_eg1() {
	r := Seq[int]{2, 3, 4, 5, 6}.Filter(func(a int) bool { return a%2 == 0 })
	fmt.Println(r)
	// Output: [2 4 6]
}

func ExampleSeq_Fold_eg1() {
	r := Seq[string]{"r", "o", "b"}.Fold("hi ", func(a1, a2 string) string { return a1 + a2 })
	fmt.Println(r)
	// Output: hi rob
}

func ExampleSeq_Fold_eg2() {
	r := Seq[int]{1, 2, 3}.Fold(10, func(a1, a2 int) int { return a1 + a2 })
	fmt.Println(r)
	// Output: 16
}

func ExampleFoldSeq_eg1() {
	r := FoldSeq(Seq[int]{1, 2, 3}, "0", func(b string, a int) string { return b + fmt.Sprint(a) })
	fmt.Println(r)
	// Output: 0123
}

func ExampleSeq_Reduce_eg1() {
	r := Seq[int]{1, 2, 3, 4}.Reduce(func(a1, a2 int) int { return a1 + a2 })
	fmt.Println(r)
	// Output: 10
}

func ExampleZipSeq_eg1() {
	r := ZipSeq(Seq[int]{1, 2, 3}, Seq[string]{"a", "b", "c"})
	fmt.Println(r)
	// Output: [(1,a) (2,b) (3,c)]
}

func ExampleZipSeq_eg2() {
	r := ZipSeq(Seq[int]{1, 2}, Seq[string]{"a", "b", "c"})
	fmt.Println(r)
	// Output: [(1,a) (2,b)]
}

func ExampleUnZipSeq_eg1() {
	r := UnZipSeq(Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b"), Tup2(3, "c")})
	fmt.Println(r)
	// Output: ([1 2 3],[a b c])
}
