package fun

import (
	"fmt"
)

// Example for: Compose3

func ExampleCompose3_eg1() {
	f := func(a int) string { return fmt.Sprint(a) }
	g := func(b string) bool { return b != "" }
	h := func(c bool) string { return fmt.Sprint(c) }

	j := Compose3(f, g, h)

	fmt.Println(j(1) == h(g(f(1))))
	// Output: true
}
