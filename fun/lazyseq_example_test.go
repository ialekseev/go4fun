package fun

import (
	"fmt"
	"strings"
)

func ExampleLazySeq_eg1() {
	// Strict (Regular) Sequence eagerly evaluates its elements.
	// Below code calculates a result in multiple iterations:
	r1 := Seq[int]{-2, -1, 0, 1, 2, 3, 4, 5, 6}.
		Filter(func(a int) bool { return a > 0 }).
		Filter(func(a int) bool { return a%2 == 0 }).
		Map(func(a int) int { return a / 2 }).
		Reduce(func(a1, a2 int) int { return a1 + a2 })

	fmt.Println(r1)

	// Lazy Sequence evaluates elements only when they are needed.
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
}

func ExampleLazySeq_eg2() {
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
}
