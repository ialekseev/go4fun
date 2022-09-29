package fun

import (
	"fmt"
)

// Examples for: Map, FlatMap, Apply, Filter, Fold, Zip, UnZip

func ExampleOption_Map_eg1() {
	r := Some("route").Map(func(a string) string { return a + "60" })
	fmt.Println(r)
	// Output: Some(route60)
}

func ExampleOption_Map_eg2() {
	r := None[string]().Map(func(a string) string { return a + "60" })
	fmt.Println(r)
	// Output: None
}

func ExampleMapOption_eg1() {
	r := MapOption(Some(60), func(a int) string { return "route" + fmt.Sprint(a) })
	fmt.Println(r)
	// Output: Some(route60)
}

func ExampleMapOption_eg2() {
	r := MapOption(None[int](), func(a int) string { return "route" + fmt.Sprint(a) })
	fmt.Println(r)
	// Output: None
}

func ExampleOption_FlatMap_eg1() {
	r := Some("route").FlatMap(func(a string) Option[string] { return Some(a + "60") })
	fmt.Println(r)
	// Output: Some(route60)
}

func ExampleOption_FlatMap_eg2() {
	r := Some("route").FlatMap(func(a string) Option[string] { return None[string]() })
	fmt.Println(r)
	// Output: None
}

func ExampleFlatMapOption_eg1() {
	r := FlatMapOption(Some(60), func(a int) Option[string] { return Some("route" + fmt.Sprint(a)) })
	fmt.Println(r)
	// Output: Some(route60)
}

func ExampleFlatMapOption_eg2() {
	r := FlatMapOption(Some(60), func(a int) Option[string] { return None[string]() })
	fmt.Println(r)
	// Output: None
}

func ExampleFlatMapOption_eg3() {
	r := FlatMapOption(None[int](), func(a int) Option[string] { return Some("route" + fmt.Sprint(a)) })
	fmt.Println(r)
	// Output: None
}

func ExampleApplyOption3_eg1() {
	r := ApplyOption3(Some(true), Some(10), Some("abc"), func(a bool, b int, c string) string {
		return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c)
	})
	fmt.Println(r)
	// Output: Some(true 10 abc)
}

func ExampleApplyOption3_eg2() {
	r := ApplyOption3(None[bool](), Some(10), Some("abc"), func(a bool, b int, c string) string {
		return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c)
	})
	fmt.Println(r)
	// Output: None
}

func ExampleOption_Filter_eg1() {
	r := Some(5).Filter(func(a int) bool { return a < 10 })
	fmt.Println(r)
	// Output: Some(5)
}

func ExampleOption_Filter_eg2() {
	r := Some(10).Filter(func(a int) bool { return a > 10 })
	fmt.Println(r)
	// Output: None
}

func ExampleOption_Filter_eg3() {
	r := None[int]().Filter(func(a int) bool { return a < 10 })
	fmt.Println(r)
	// Output: None
}

func ExampleOption_Fold_eg1() {
	r := Some(5).Fold(1, func(a int) int { return a * 2 })
	fmt.Println(r)
	// Output: 10
}

func ExampleOption_Fold_eg2() {
	r := None[int]().Fold(1, func(a int) int { return a * 2 })
	fmt.Println(r)
	// Output: 1
}

func ExampleFoldOption_eg1() {
	r := FoldOption(Some(60), "route0", func(a int) string { return "route" + fmt.Sprint(a) })
	fmt.Println(r)
	// Output: route60
}

func ExampleFoldOption_eg2() {
	r := FoldOption(None[int](), "route0", func(a int) string { return "route" + fmt.Sprint(a) })
	fmt.Println(r)
	// Output: route0
}

func ExampleZipOption_eg1() {
	r := ZipOption(Some("route"), Some(60))
	fmt.Println(r)
	// Output: Some((route,60))
}

func ExampleZipOption_eg2() {
	r := ZipOption(None[string](), Some(60))
	fmt.Println(r)
	// Output: None
}

func ExampleZipOption_eg3() {
	r := ZipOption(Some("route"), None[int]())
	fmt.Println(r)
	// Output: None
}

func ExampleUnZipOption_eg1() {
	r := UnZipOption(Some(Tup2("route", 60)))
	fmt.Println(r)
	// Output: (Some(route),Some(60))
}
