package fun

import (
	"fmt"
	"time"
)

func ExampleMemo1_eg1() {
	// an expensive function (with 1 argument) is wrapped into a memo function of the same signature.
	// it would cache results of function calls and return back a cached result for the same input, if requested again.
	var memoF = Memo1(func(a int) string {
		// expensive computation:
		time.Sleep(time.Millisecond * time.Duration(a))
		return fmt.Sprint(a)
	})

	r := memoF(2) // the first call is slow
	r = memoF(2)  // other calls are fast

	fmt.Println(r)
	// Output: 2
}

func ExampleMemo2_eg1() {
	// an expensive function (with 2 arguments) is wrapped into a memo function of the same signature.
	// it would cache results of function calls and return back a cached result for the same inputs, if requested again.
	var memoF = Memo2(func(a, b int) string {
		// expensive computation:
		time.Sleep(time.Millisecond * time.Duration(a+b))
		return fmt.Sprint(a) + fmt.Sprint(b)
	})

	r := memoF(2, 2) // the first call is slow
	r = memoF(2, 2)  // other calls are fast

	fmt.Println(r)
	// Output: 22
}
