package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOption_ApplyOption1(t *testing.T) {
	assert.Equal(t, Some("5A"), ApplyOption1(Some(5), func(a int) Option[string] { return Some(fmt.Sprint(a) + "A") }))
	assert.Equal(t, None[string](), ApplyOption1(None[int](), func(a int) Option[string] { return Some(fmt.Sprint(a) + "A") }))
}

func TestOption_ApplyOption2(t *testing.T) {
	assert.Equal(t, Some("true 10"), ApplyOption2(Some(true), Some(10), func(a bool, b int) Option[string] { return Some(fmt.Sprint(a) + " " + fmt.Sprint(b)) }))
	assert.Equal(t, None[string](), ApplyOption2(None[bool](), Some(10), func(a bool, b int) Option[string] { return Some(fmt.Sprint(a) + " " + fmt.Sprint(b)) }))
	assert.Equal(t, None[string](), ApplyOption2(Some(true), None[int](), func(a bool, b int) Option[string] { return Some(fmt.Sprint(a) + " " + fmt.Sprint(b)) }))
	assert.Equal(t, None[string](), ApplyOption2(None[bool](), None[int](), func(a bool, b int) Option[string] { return Some(fmt.Sprint(a) + " " + fmt.Sprint(b)) }))
}

func TestOption_ApplyOption3(t *testing.T) {
	assert.Equal(t, Some("true 10 abc"), ApplyOption3(Some(true), Some(10), Some("abc"), func(a bool, b int, c string) Option[string] {
		return Some(fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c))
	}))
	assert.Equal(t, None[string](), ApplyOption3(None[bool](), Some(10), Some("abc"), func(a bool, b int, c string) Option[string] {
		return Some(fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c))
	}))
	assert.Equal(t, None[string](), ApplyOption3(Some(true), None[int](), Some("abc"), func(a bool, b int, c string) Option[string] {
		return Some(fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c))
	}))
	assert.Equal(t, None[string](), ApplyOption3(Some(true), Some(10), None[string](), func(a bool, b int, c string) Option[string] {
		return Some(fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c))
	}))
	assert.Equal(t, None[string](), ApplyOption3(None[bool](), None[int](), None[string](), func(a bool, b int, c string) Option[string] {
		return Some(fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c))
	}))
}

func TestOption_ContainsInOption(t *testing.T) {
	assert.True(t, ContainsInOption(Some(5), 5))
	assert.False(t, ContainsInOption(Some(5), 6))
	assert.False(t, ContainsInOption(None[int](), 5))
}

func TestOption_Exists(t *testing.T) {
	assert.True(t, Some(5).Exists(func(a int) bool { return a < 6 }))
	assert.False(t, Some(5).Exists(func(a int) bool { return a < 5 }))
	assert.False(t, None[int]().Exists(func(a int) bool { return a == 5 }))
}

func TestOption_Filter(t *testing.T) {
	assert.Equal(t, Some(5), Some(5).Filter(func(a int) bool { return a < 6 }))
	assert.Equal(t, None[int](), Some(5).Filter(func(a int) bool { return a < 5 }))
	assert.Equal(t, None[int](), None[int]().Filter(func(a int) bool { return a == 5 }))
}

func TestOption_FilterNot(t *testing.T) {
	assert.Equal(t, Some(5), Some(5).FilterNot(func(a int) bool { return a < 5 }))
	assert.Equal(t, None[int](), Some(5).FilterNot(func(a int) bool { return a < 6 }))
	assert.Equal(t, None[int](), None[int]().FilterNot(func(a int) bool { return a == 5 }))
}

func TestOption_FlatMap(t *testing.T) {
	assert.Equal(t, Some("abcdef"), Some("abc").FlatMap(func(a string) Option[string] { return Some(a + "def") }))
	assert.Equal(t, None[string](), Some("abc").FlatMap(func(a string) Option[string] { return None[string]() }))
	assert.Equal(t, None[string](), None[string]().FlatMap(func(a string) Option[string] { return Some(a + "def") }))
}

func TestOption_FlatMapOption(t *testing.T) {
	assert.Equal(t, Some("123"), FlatMapOption(Some(123), func(a int) Option[string] { return Some("123") }))
	assert.Equal(t, None[string](), FlatMapOption(Some(123), func(a int) Option[string] { return None[string]() }))
	assert.Equal(t, None[string](), FlatMapOption(None[int](), func(a int) Option[string] { return Some("123") }))
}

func TestOption_Flatten(t *testing.T) {
	assert.Equal(t, Some(5), FlattenOption(Some(Some(5))))
	assert.Equal(t, None[int](), FlattenOption(Some(None[int]())))
}

func TestOption_Fold(t *testing.T) {
	assert.Equal(t, 6, Some(5).Fold(-1, func(a int) int { return a + 1 }))
	assert.Equal(t, -1, None[int]().Fold(-1, func(a int) int { return a + 1 }))
}

func TestOption_FoldOption(t *testing.T) {
	assert.Equal(t, "5A", FoldOption(Some(5), "", func(a int) string { return fmt.Sprint(a) + "A" }))
	assert.Equal(t, "", FoldOption(None[int](), "", func(a int) string { return fmt.Sprint(a) + "A" }))
}

func TestOption_ForAll(t *testing.T) {
	assert.True(t, Some("abc").ForAll(func(a string) bool { return a == "abc" }))
	assert.True(t, None[string]().ForAll(func(a string) bool { return a == "abc" }))
	assert.False(t, Some("abc").ForAll(func(a string) bool { return a == "def" }))
}

func TestOption_Foreach(t *testing.T) {
	//given
	e := 0
	//when
	Some(5).Foreach(func(a int) { e = a })
	//then
	assert.Equal(t, 5, e)
}

func TestOption_Get(t *testing.T) {
	assert.Equal(t, "abc", Some("abc").Get())
	assert.Equal(t, "", None[string]().Get())
	assert.Equal(t, 0, None[int]().Get())
}

func TestOption_GetOrElse(t *testing.T) {
	assert.Equal(t, "abc", Some("abc").GetOrElse(""))
	assert.Equal(t, "", None[string]().GetOrElse(""))
}

func TestOption_IsDefined(t *testing.T) {
	assert.True(t, Some(5).IsDefined())
	assert.False(t, None[int]().IsDefined())
}

func TestOption_IsEmpty(t *testing.T) {
	assert.True(t, None[int]().IsEmpty())
	assert.False(t, Some(5).IsEmpty())
}

func TestOption_Map(t *testing.T) {
	assert.Equal(t, Some(6), Some(5).Map(func(a int) int { return a + 1 }))
	assert.Equal(t, None[int](), None[int]().Map(func(a int) int { return a + 1 }))
}

func TestOption_MapOption(t *testing.T) {
	assert.Equal(t, Some("5A"), MapOption(Some(5), func(a int) string { return fmt.Sprint(a) + "A" }))
	assert.Equal(t, None[string](), MapOption(None[int](), func(a int) string { return fmt.Sprint(a) + "A" }))
}

func TestOption_None(t *testing.T) {
	//when
	option := None[string]()
	//then
	assert.Equal(t, false, option.defined)
	assert.Equal(t, "", option.value)
}

func TestOption_NonEmpty(t *testing.T) {
	assert.True(t, Some(5).NonEmpty())
	assert.False(t, None[int]().NonEmpty())
}

func TestOption_OrElse(t *testing.T) {
	assert.Equal(t, Some("abc"), Some("abc").OrElse(Some("def")))
	assert.Equal(t, Some("def"), None[string]().OrElse(Some("def")))
	assert.Equal(t, None[string](), None[string]().OrElse(None[string]()))
}

func TestOption_Some(t *testing.T) {
	//when
	option := Some(5)
	//then
	assert.Equal(t, true, option.defined)
	assert.Equal(t, 5, option.value)
}

func TestOption_String(t *testing.T) {
	assert.Equal(t, "Some(5)", Some(5).String())
	assert.Equal(t, "None", None[int]().String())
}

func TestOption_ToSeq(t *testing.T) {
	assert.Equal(t, Seq[int]{5}, Some(5).ToSeq())
	assert.Nil(t, None[int]().ToSeq())
}

func TestOption_UnZipOption(t *testing.T) {
	assert.Equal(t, Tup2(Some(5), Some("abc")), UnZipOption(Some(Tup2(5, "abc"))))
	assert.Equal(t, Tup2(None[int](), None[string]()), UnZipOption(None[Tuple2[int, string]]()))
}

func TestOption_ZipOption(t *testing.T) {
	assert.Equal(t, Some(Tup2(5, "123")), ZipOption(Some(5), Some("123")))
	assert.Equal(t, None[Tuple2[int, string]](), ZipOption(Some(5), None[string]()))
	assert.Equal(t, None[Tuple2[int, string]](), ZipOption(None[int](), Some("123")))
	assert.Equal(t, None[Tuple2[int, string]](), ZipOption(None[int](), None[string]()))
}
