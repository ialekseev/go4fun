package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionApplyOption1(t *testing.T) {
	assert.Equal(t, Some("5A"), ApplyOption1(Some(5), func(a int) string { return fmt.Sprint(a) + "A" }))
	assert.Equal(t, None[string](), ApplyOption1(None[int](), func(a int) string { return fmt.Sprint(a) + "A" }))
}

func TestOptionApplyOption2(t *testing.T) {
	assert.Equal(t, Some("true 10"), ApplyOption2(Some(true), Some(10), func(a bool, b int) string { return fmt.Sprint(a) + " " + fmt.Sprint(b) }))
	assert.Equal(t, None[string](), ApplyOption2(None[bool](), Some(10), func(a bool, b int) string { return fmt.Sprint(a) + " " + fmt.Sprint(b) }))
	assert.Equal(t, None[string](), ApplyOption2(Some(true), None[int](), func(a bool, b int) string { return fmt.Sprint(a) + " " + fmt.Sprint(b) }))
	assert.Equal(t, None[string](), ApplyOption2(None[bool](), None[int](), func(a bool, b int) string { return fmt.Sprint(a) + " " + fmt.Sprint(b) }))
}

func TestOptionApplyOption3(t *testing.T) {
	assert.Equal(t, Some("true 10 abc"), ApplyOption3(Some(true), Some(10), Some("abc"), func(a bool, b int, c string) string { return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c) }))
	assert.Equal(t, None[string](), ApplyOption3(None[bool](), Some(10), Some("abc"), func(a bool, b int, c string) string { return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c) }))
	assert.Equal(t, None[string](), ApplyOption3(Some(true), None[int](), Some("abc"), func(a bool, b int, c string) string { return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c) }))
	assert.Equal(t, None[string](), ApplyOption3(Some(true), Some(10), None[string](), func(a bool, b int, c string) string { return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c) }))
	assert.Equal(t, None[string](), ApplyOption3(None[bool](), None[int](), None[string](), func(a bool, b int, c string) string { return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c) }))
}

func TestOptionContainsInOption(t *testing.T) {
	assert.True(t, ContainsInOption(Some(5), 5))
	assert.False(t, ContainsInOption(Some(5), 6))
	assert.False(t, ContainsInOption(None[int](), 5))
}

func TestOptionExists(t *testing.T) {
	assert.True(t, Some(5).Exists(func(a int) bool { return a < 6 }))
	assert.False(t, Some(5).Exists(func(a int) bool { return a < 5 }))
	assert.False(t, None[int]().Exists(func(a int) bool { return a == 5 }))
}

func TestOptionFilter(t *testing.T) {
	assert.Equal(t, Some(5), Some(5).Filter(func(a int) bool { return a < 6 }))
	assert.Equal(t, None[int](), Some(5).Filter(func(a int) bool { return a < 5 }))
	assert.Equal(t, None[int](), None[int]().Filter(func(a int) bool { return a == 5 }))
}

func TestOptionFilterNot(t *testing.T) {
	assert.Equal(t, Some(5), Some(5).FilterNot(func(a int) bool { return a < 5 }))
	assert.Equal(t, None[int](), Some(5).FilterNot(func(a int) bool { return a < 6 }))
	assert.Equal(t, None[int](), None[int]().FilterNot(func(a int) bool { return a == 5 }))
}

func TestOptionFlatMap(t *testing.T) {
	assert.Equal(t, Some("abcdef"), Some("abc").FlatMap(func(a string) Option[string] { return Some(a + "def") }))
	assert.Equal(t, None[string](), Some("abc").FlatMap(func(a string) Option[string] { return None[string]() }))
	assert.Equal(t, None[string](), None[string]().FlatMap(func(a string) Option[string] { return Some(a + "def") }))
}

func TestOptionFlatMapOption(t *testing.T) {
	assert.Equal(t, Some("123"), FlatMapOption(Some(123), func(a int) Option[string] { return Some("123") }))
	assert.Equal(t, None[string](), FlatMapOption(Some(123), func(a int) Option[string] { return None[string]() }))
	assert.Equal(t, None[string](), FlatMapOption(None[int](), func(a int) Option[string] { return Some("123") }))
}

func TestOptionFlatten(t *testing.T) {
	assert.Equal(t, Some(5), FlattenOption(Some(Some(5))))
	assert.Equal(t, None[int](), FlattenOption(Some(None[int]())))
}

func TestOptionFold(t *testing.T) {
	assert.Equal(t, 6, Some(5).Fold(-1, func(a int) int { return a + 1 }))
	assert.Equal(t, -1, None[int]().Fold(-1, func(a int) int { return a + 1 }))
}

func TestOptionFoldOption(t *testing.T) {
	assert.Equal(t, "5A", FoldOption(Some(5), "", func(a int) string { return fmt.Sprint(a) + "A" }))
	assert.Equal(t, "", FoldOption(None[int](), "", func(a int) string { return fmt.Sprint(a) + "A" }))
}

func TestOptionForAll(t *testing.T) {
	assert.True(t, Some("abc").ForAll(func(a string) bool { return a == "abc" }))
	assert.True(t, None[string]().ForAll(func(a string) bool { return a == "abc" }))
	assert.False(t, Some("abc").ForAll(func(a string) bool { return a == "def" }))
}

func TestOptionForeach(t *testing.T) {
	//given
	e := 0
	//when
	Some(5).Foreach(func(a int) { e = a })
	//then
	assert.Equal(t, 5, e)
}

func TestOptionGet(t *testing.T) {
	assert.Equal(t, "abc", Some("abc").Get())
	assert.Equal(t, "", None[string]().Get())
	assert.Equal(t, 0, None[int]().Get())
}

func TestOptionGetOrElse(t *testing.T) {
	assert.Equal(t, "abc", Some("abc").GetOrElse(""))
	assert.Equal(t, "", None[string]().GetOrElse(""))
}

func TestOptionIsDefined(t *testing.T) {
	assert.True(t, Some(5).IsDefined())
	assert.False(t, None[int]().IsDefined())
}

func TestOptionIsEmpty(t *testing.T) {
	assert.True(t, None[int]().IsEmpty())
	assert.False(t, Some(5).IsEmpty())
}

func TestOptionMap(t *testing.T) {
	assert.Equal(t, Some(6), Some(5).Map(func(a int) int { return a + 1 }))
	assert.Equal(t, None[int](), None[int]().Map(func(a int) int { return a + 1 }))
}

func TestOptionMapOption(t *testing.T) {
	assert.Equal(t, Some("5A"), MapOption(Some(5), func(a int) string { return fmt.Sprint(a) + "A" }))
	assert.Equal(t, None[string](), MapOption(None[int](), func(a int) string { return fmt.Sprint(a) + "A" }))
}

func TestOptionNone(t *testing.T) {
	//when
	option := None[string]()
	//then
	assert.Equal(t, false, option.defined)
	assert.Equal(t, "", option.value)
}

func TestOptionNonEmpty(t *testing.T) {
	assert.True(t, Some(5).NonEmpty())
	assert.False(t, None[int]().NonEmpty())
}

func TestOptionOrElse(t *testing.T) {
	assert.Equal(t, Some("abc"), Some("abc").OrElse(Some("def")))
	assert.Equal(t, Some("def"), None[string]().OrElse(Some("def")))
	assert.Equal(t, None[string](), None[string]().OrElse(None[string]()))
}

func TestOptionSome(t *testing.T) {
	//when
	option := Some(5)
	//then
	assert.Equal(t, true, option.defined)
	assert.Equal(t, 5, option.value)
}

func TestOptionString(t *testing.T) {
	assert.Equal(t, "Some(5)", Some(5).String())
	assert.Equal(t, "None", None[int]().String())
}

func TestOptionToSeq(t *testing.T) {
	assert.Equal(t, Seq[int]{5}, Some(5).ToSeq())
	assert.Nil(t, None[int]().ToSeq())
}

func TestOptionUnZipOption(t *testing.T) {
	assert.Equal(t, Tup2(Some(5), Some("abc")), UnZipOption(Some(Tup2(5, "abc"))))
	assert.Equal(t, Tup2(None[int](), None[string]()), UnZipOption(None[Tuple2[int, string]]()))
}

func TestOptionZipOption(t *testing.T) {
	assert.Equal(t, Some(Tup2(5, "123")), ZipOption(Some(5), Some("123")))
	assert.Equal(t, None[Tuple2[int, string]](), ZipOption(Some(5), None[string]()))
	assert.Equal(t, None[Tuple2[int, string]](), ZipOption(None[int](), Some("123")))
	assert.Equal(t, None[Tuple2[int, string]](), ZipOption(None[int](), None[string]()))
}
