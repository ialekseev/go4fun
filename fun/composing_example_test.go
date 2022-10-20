package fun

import (
	"fmt"
)

// Example for: Compose2

func ExampleCompose2_eg1() {
	f := func(a int) string { return fmt.Sprint(a) }
	g := func(b string) bool { return b != "" }

	h := Compose2(f, g)

	fmt.Println(h(1) == g(f(1)))
	// Output: true
}
