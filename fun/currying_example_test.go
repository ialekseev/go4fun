package fun

import (
	"fmt"
)

// Examples for: Curry3, UnCurry3

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
