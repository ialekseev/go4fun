package fun

import (
	"fmt"
)

func ExampleOption_Map_eg1() {
	r := Some("abc").Map(func(a string) string { return a + "def" })
	fmt.Println(r)
	// Output: Some(abcdef)
}

func ExampleOption_Map_eg2() {
	r := None[string]().Map(func(a string) string { return a + "def" })
	fmt.Println(r)
	// Output: None
}

func ExampleOption_FlatMap_eg1() {
	r := Some("abc").FlatMap(func(a string) Option[string] { return Some(a + "def") })
	fmt.Println(r)
	// Output: Some(abcdef)
}

func ExampleOption_FlatMap_eg2() {
	r := Some("abc").FlatMap(func(a string) Option[string] { return None[string]() })
	fmt.Println(r)
	// Output: None
}
