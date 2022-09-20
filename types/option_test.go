package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	//when
	option := Some(5)
	//then
	assert.Equal(t, true, option.defined)
	assert.Equal(t, 5, option.value)
}

func TestNone(t *testing.T) {
	//when
	option := None[string]()
	//then
	assert.Equal(t, false, option.defined)
	assert.Equal(t, "", option.value)
}

func TestContains(t *testing.T) {
	assert.True(t, Some(5).Contains(5))
	assert.False(t, Some(5).Contains(6))
	assert.False(t, None[int]().Contains(5))
}

func TestExists(t *testing.T) {
	assert.True(t, Some(5).Exists(func(v int) bool { return v < 6 }))
	assert.False(t, Some(5).Exists(func(v int) bool { return v < 5 }))
	assert.False(t, None[int]().Exists(func(v int) bool { return v == 5 }))
}

func TestFilter(t *testing.T) {
	assert.Equal(t, Some(5), Some(5).Filter(func(v int) bool { return v < 6 }))
	assert.Equal(t, None[int](), Some(5).Filter(func(v int) bool { return v < 5 }))
	assert.Equal(t, None[int](), None[int]().Filter(func(v int) bool { return v == 5 }))
}

func TestFilterNot(t *testing.T) {
	assert.Equal(t, Some(5), Some(5).FilterNot(func(v int) bool { return v < 5 }))
	assert.Equal(t, None[int](), Some(5).FilterNot(func(v int) bool { return v < 6 }))
	assert.Equal(t, None[int](), None[int]().FilterNot(func(v int) bool { return v == 5 }))
}

func TestFlatMap(t *testing.T) {
	assert.Equal(t, Some("abcdef"), Some("abc").FlatMap(func(s string) Option[string] { return Some(s + "def") }))
	assert.Equal(t, None[string](), Some("abc").FlatMap(func(s string) Option[string] { return None[string]() }))
	assert.Equal(t, None[string](), None[string]().FlatMap(func(s string) Option[string] { return Some(s + "def") }))
}

func TestFlatMapOption(t *testing.T) {
	assert.Equal(t, Some("123"), FlatMapOption(Some(123), func(s int) Option[string] { return Some("123") }))
	assert.Equal(t, None[string](), FlatMapOption(Some(123), func(s int) Option[string] { return None[string]() }))
	assert.Equal(t, None[string](), FlatMapOption(None[int](), func(s int) Option[string] { return Some("123") }))
}

func TestFlatten(t *testing.T) {
	assert.Equal(t, Some(5), FlattenOption(Some(Some(5))))
	assert.Equal(t, None[int](), FlattenOption(Some(None[int]())))
}

func TestFold(t *testing.T) {
	assert.Equal(t, 6, Some(5).Fold(-1, func(v int) int { return v + 1 }))
	assert.Equal(t, -1, None[int]().Fold(-1, func(v int) int { return v + 1 }))
}

func TestFoldOption(t *testing.T) {
	assert.Equal(t, "5", FoldOption(Some(5), "", func(v int) string { return "5" }))
	assert.Equal(t, "", FoldOption(None[int](), "", func(v int) string { return "5" }))
}

func TestForAll(t *testing.T) {
	assert.True(t, Some("abc").ForAll(func(s string) bool { return s == "abc" }))
	assert.True(t, None[string]().ForAll(func(s string) bool { return s == "abc" }))
	assert.False(t, Some("abc").ForAll(func(s string) bool { return s == "def" }))
}
