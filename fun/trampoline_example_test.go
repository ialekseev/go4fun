package fun

import (
	"fmt"
)

// Summation of n natural numbers (see: https://en.wikipedia.org/wiki/Summation).
// For example, a summation of the first 100 natural numbers is 1 + 2 + 3 + 4 + ⋯ + 99 + 100.
// Calling this function with a big n like 100000000 will fail with a stack overflow.
func summation(n, current uint) uint {
	if n < 1 {
		return current
	}
	return summation(n-1, n+current)
}

// Summation with Trampoline.
// Running this function with a big n like 100000000 will still be OK.
func summationT(n, current uint) Trampoline[uint] {
	if n < 1 {
		return DoneTrampolining(current)
	}
	return MoreTrampolining(func() Trampoline[uint] {
		return summationT(n-1, n+current)
	})
}

func ExampleTrampoline_Run_eg1() {
	/*
			func summationT(n, current uint) Trampoline[uint] {
				if n < 1 {
					return DoneTrampolining(current)
				}
				return MoreTrampolining(func() Trampoline[uint] {
					return summationT(n-1, n+current)
				})
		}
	*/
	r := summationT(5, 0).Run() //Change 5 to 100000000. It Will still be ok.

	fmt.Println(r)
	// Output: 15
}

func ExampleTrampoline_Run_eg2() {
	/*
		func summation(n, current uint) uint {
			if n < 1 {
				return current
			}
			return summation(n-1, n+current)
		}
	*/

	r := summation(5, 0) //Change 5 to 100000000. It Will fail with a stack overflow.

	fmt.Println(r)
	// Output: 15
}
