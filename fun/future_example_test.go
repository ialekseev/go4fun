package fun

import (
	"fmt"
	"time"
)

// Examples for: Map, FlatMap, OnComplete

func ExampleFuture_Map_eg1() {
	future := FutureValue(func() string {
		time.Sleep(time.Millisecond * 20)
		return "abc"
	})

	r := future.Map(func(a string) string { return a + "def" })

	time.Sleep(time.Millisecond * 30)

	fmt.Println(r.Value())
	// Output: Some(abcdef)
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

	fmt.Println(r.Value())
	// Output: Some(abcdef)
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
