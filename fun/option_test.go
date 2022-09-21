package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionSome(t *testing.T) {
	//when
	option := Some(5)
	//then
	assert.Equal(t, true, option.defined)
	assert.Equal(t, 5, option.value)
}

func TestOptionNone(t *testing.T) {
	//when
	option := None[string]()
	//then
	assert.Equal(t, false, option.defined)
	assert.Equal(t, "", option.value)
}

func TestOptionContains(t *testing.T) {
	assert.True(t, Some(5).Contains(5))
	assert.False(t, Some(5).Contains(6))
	assert.False(t, None[int]().Contains(5))
}

func TestOptionExists(t *testing.T) {
	assert.True(t, Some(5).Exists(func(v int) bool { return v < 6 }))
	assert.False(t, Some(5).Exists(func(v int) bool { return v < 5 }))
	assert.False(t, None[int]().Exists(func(v int) bool { return v == 5 }))
}

func TestOptionFilter(t *testing.T) {
	assert.Equal(t, Some(5), Some(5).Filter(func(v int) bool { return v < 6 }))
	assert.Equal(t, None[int](), Some(5).Filter(func(v int) bool { return v < 5 }))
	assert.Equal(t, None[int](), None[int]().Filter(func(v int) bool { return v == 5 }))
}

func TestOptionFilterNot(t *testing.T) {
	assert.Equal(t, Some(5), Some(5).FilterNot(func(v int) bool { return v < 5 }))
	assert.Equal(t, None[int](), Some(5).FilterNot(func(v int) bool { return v < 6 }))
	assert.Equal(t, None[int](), None[int]().FilterNot(func(v int) bool { return v == 5 }))
}

func TestOptionFlatMap(t *testing.T) {
	assert.Equal(t, Some("abcdef"), Some("abc").FlatMap(func(s string) Option[string] { return Some(s + "def") }))
	assert.Equal(t, None[string](), Some("abc").FlatMap(func(s string) Option[string] { return None[string]() }))
	assert.Equal(t, None[string](), None[string]().FlatMap(func(s string) Option[string] { return Some(s + "def") }))
}

func TestOptionFlatMapOption(t *testing.T) {
	assert.Equal(t, Some("123"), FlatMapOption(Some(123), func(s int) Option[string] { return Some("123") }))
	assert.Equal(t, None[string](), FlatMapOption(Some(123), func(s int) Option[string] { return None[string]() }))
	assert.Equal(t, None[string](), FlatMapOption(None[int](), func(s int) Option[string] { return Some("123") }))
}

func TestOptionFlatten(t *testing.T) {
	assert.Equal(t, Some(5), FlattenOption(Some(Some(5))))
	assert.Equal(t, None[int](), FlattenOption(Some(None[int]())))
}

func TestOptionFold(t *testing.T) {
	assert.Equal(t, 6, Some(5).Fold(-1, func(v int) int { return v + 1 }))
	assert.Equal(t, -1, None[int]().Fold(-1, func(v int) int { return v + 1 }))
}

func TestOptionFoldOption(t *testing.T) {
	assert.Equal(t, "5A", FoldOption(Some(5), "", func(v int) string { return fmt.Sprint(v) + "A" }))
	assert.Equal(t, "", FoldOption(None[int](), "", func(v int) string { return fmt.Sprint(v) + "A" }))
}

func TestOptionForAll(t *testing.T) {
	assert.True(t, Some("abc").ForAll(func(s string) bool { return s == "abc" }))
	assert.True(t, None[string]().ForAll(func(s string) bool { return s == "abc" }))
	assert.False(t, Some("abc").ForAll(func(s string) bool { return s == "def" }))
}

func TestOptionForeach(t *testing.T) {
	//given
	e := 0
	//when
	Some(5).Foreach(func(v int) { e = v })
	//then
	assert.Equal(t, 5, e)
}

func TestOptionGet(t *testing.T) {
	assert.Equal(t, "abc", Some("abc").Get())
	assert.Panics(t, func() { None[string]().Get() })
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
	assert.Equal(t, Some(6), Some(5).Map(func(v int) int { return v + 1 }))
	assert.Equal(t, None[int](), None[int]().Map(func(v int) int { return v + 1 }))
}

func TestOptionMapOption(t *testing.T) {
	assert.Equal(t, Some("5A"), MapOption(Some(5), func(v int) string { return fmt.Sprint(v) + "A" }))
	assert.Equal(t, None[string](), MapOption(None[int](), func(v int) string { return fmt.Sprint(v) + "A" }))
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

func TestOptionString(t *testing.T) {
	assert.Equal(t, "Some(5)", Some(5).String())
	assert.Equal(t, "None", None[int]().String())
}

func TestOptionToSeq(t *testing.T) {
	assert.Equal(t, Seq[int]{5}, Some(5).ToSeq())
	assert.Nil(t, None[int]().ToSeq())
}

func TestOptionUnZipOption(t *testing.T) {
	assert.Equal(t, NewTuple2(Some(5), Some("abc")), UnZipOption(Some(NewTuple2(5, "abc"))))
	assert.Equal(t, NewTuple2(None[int](), None[string]()), UnZipOption(None[Tuple2[int, string]]()))
}

func TestOptionZipOption(t *testing.T) {
	assert.Equal(t, Some(NewTuple2(5, "123")), ZipOption(Some(5), Some("123")))
	assert.Equal(t, None[Tuple2[int, string]](), ZipOption(Some(5), None[string]()))
	assert.Equal(t, None[Tuple2[int, string]](), ZipOption(None[int](), Some("123")))
	assert.Equal(t, None[Tuple2[int, string]](), ZipOption(None[int](), None[string]()))
}
