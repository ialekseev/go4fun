package fun

import (
	"fmt"
)

// Examples for: Apply3Partial, Compose2, Curry3, UnCurry3

func ExampleApply3_eg1() {
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c)
	}

	// function `f` is applied only to the 1st and the 2nd argument.
	// resulting function `p` has only 1 remaining argument.
	p := Apply3Partial_1_2(f, 10, true)

	fmt.Println(p(5.5))
	// Output: 10 true 5.5
}

func ExampleCompose2_eg1() {
	f := func(a int) string { return fmt.Sprint(a) }
	g := func(b string) bool { return b != "" }

	h := Compose2(f, g)

	fmt.Println(h(1) == g(f(1)))
	// Output: true
}

func ExampleCurry3_eg1() {
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c)
	}

	curriedF := Curry3(f)
	r := curriedF(1)(true)(5.5)

	fmt.Println(r)
	// Output: 1 true 5.5
}

func ExampleUnCurry3_eg1() {
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
}
