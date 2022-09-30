package fun

import (
	"fmt"
	"time"
)

// Examples for: Map, FlatMap, Apply, OnComplete

func ExampleFuture_Map_eg1() {
	future := FutureValue(func() string {
		time.Sleep(time.Millisecond * 20)
		return "abc"
	})

	r := future.Map(func(a string) string { return a + "def" })

	time.Sleep(time.Millisecond * 30)

	fmt.Println(r.Result())
	// Output: abcdef
}

func ExampleFuture_FlatMap_eg1() {
	future := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})

	r := future.FlatMap(func(a string) Future[string] {
		return FutureValue(func() string {
			time.Sleep(time.Millisecond * 10)
			return a + "def"
		})
	})

	time.Sleep(time.Millisecond * 50)

	fmt.Println(r.Result())
	// Output: abcdef
}

func ExampleApplyFuture2_eg1() {
	future1 := FutureValue(func() int {
		time.Sleep(time.Millisecond * 10)
		return 123
	})

	future2 := FutureValue(func() bool {
		time.Sleep(time.Millisecond * 10)
		return true
	})

	r := ApplyFuture2(future1, future2, func(a int, b bool) Future[string] {
		return FutureValue(func() string {
			time.Sleep(time.Millisecond * 10)
			return fmt.Sprint(a) + " " + fmt.Sprint(b)
		})
	})

	time.Sleep(time.Millisecond * 60)

	fmt.Println(r.Result())
	// Output: 123 true
}

func ExampleFuture_OnComplete_eg1() {
	future := FutureValue(func() string {
		time.Sleep(time.Millisecond * 20)
		return "abc"
	})

	future.OnComplete(func(a string) { fmt.Println(a + "def") })

	time.Sleep(time.Millisecond * 30)

	// Output: abcdef
}
