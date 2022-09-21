package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func nilSeq[T comparable]() Seq[T] {
	var n Seq[T] = nil
	return n
}

func TestSeqAppend(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 2, 3}, Seq[int]{1, 2}.Append(3))
	assert.Equal(t, Seq[int]{1}, Seq[int]{}.Append(1))
	assert.Equal(t, Seq[int]{1}, nilSeq[int]().Append(1))
}

func TestSeqContains(t *testing.T) {
	assert.True(t, Seq[int]{1, 2, 5}.Contains(2))
	assert.False(t, Seq[int]{1, 2, 5}.Contains(6))
	assert.False(t, Seq[int]{}.Contains(2))
	assert.False(t, nilSeq[int]().Contains(2))
}

func TestSeqDistinct(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 2, 3}, Seq[int]{1, 1, 2, 3, 3, 3}.Distinct())
	assert.Equal(t, Seq[int]{1, 2, 3}, Seq[int]{1, 2, 3}.Distinct())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Distinct())
	assert.Nil(t, nilSeq[int]().Distinct())
}

func TestSeqEmptySeq(t *testing.T) {
	assert.Equal(t, 0, EmptySeq[int](5).Length())
}

func TestSeqExists(t *testing.T) {
	assert.True(t, Seq[int]{2, 4, 5}.Exists(func(a int) bool { return a > 4 }))
	assert.False(t, Seq[int]{2, 4, 5}.Exists(func(a int) bool { return a > 5 }))
	assert.False(t, Seq[int]{}.Exists(func(a int) bool { return a > 0 }))
	assert.False(t, nilSeq[int]().Exists(func(a int) bool { return a > 0 }))
}

func TestSeqFilter(t *testing.T) {
	assert.Equal(t, Seq[int]{2, 4, 6}, Seq[int]{2, 3, 4, 5, 6}.Filter(func(a int) bool { return a%2 == 0 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{2, 3, 4, 5, 6}.Filter(func(a int) bool { return a > 6 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Filter(func(a int) bool { return a > 0 }))
	assert.Nil(t, nilSeq[int]().Filter(func(a int) bool { return a > 0 }))
}

func TestSeqFilterNot(t *testing.T) {
	assert.Equal(t, Seq[int]{3, 5}, Seq[int]{2, 3, 4, 5, 6}.FilterNot(func(a int) bool { return a%2 == 0 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{2, 3, 4, 5, 6}.FilterNot(func(a int) bool { return a >= 2 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{}.FilterNot(func(a int) bool { return a > 0 }))
	assert.Nil(t, nilSeq[int]().FilterNot(func(a int) bool { return a > 0 }))
}

func TestSeqFind(t *testing.T) {
	assert.Equal(t, Some(3), Seq[int]{1, 2, 3, 4, 5}.Find(func(a int) bool { return a > 2 }))
	assert.Equal(t, None[int](), Seq[int]{1, 2, 3, 4, 5}.Find(func(a int) bool { return a > 5 }))
	assert.Equal(t, None[int](), Seq[int]{}.Find(func(a int) bool { return a > 0 }))
	assert.Equal(t, None[int](), nilSeq[int]().Find(func(a int) bool { return a > 0 }))
}

func TestSeqFlatMap(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 1, 2, 2, 3, 3}, Seq[int]{1, 2, 3}.FlatMap(func(a int) Seq[int] { return Seq[int]{a, a} }))
	assert.Equal(t, Seq[int]{}, Seq[int]{}.FlatMap(func(a int) Seq[int] { return Seq[int]{a, a} }))
	assert.Nil(t, nilSeq[int]().FlatMap(func(a int) Seq[int] { return Seq[int]{a, a} }))
}

func TestSeqFlatMatSeq(t *testing.T) {
	assert.Equal(t, Seq[string]{"1", "1", "2", "2", "3", "3"}, FlatMapSeq(Seq[int]{1, 2, 3}, func(a int) Seq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)} }))
	assert.Equal(t, Seq[string]{}, FlatMapSeq(Seq[int]{}, func(a int) Seq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)} }))
	assert.Nil(t, FlatMapSeq(nilSeq[int](), func(a int) Seq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)} }))
}

func TestSeqFlattenSeq(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 2, 3, 4, 5}, FlattenSeq([]Seq[int]{{1, 2, 3}, {4, 5}}))
	assert.Equal(t, Seq[int]{}, FlattenSeq([]Seq[int]{}))

	var n []Seq[int] = nil
	assert.Nil(t, FlattenSeq(n))
}

func TestSeqFold(t *testing.T) {
	assert.Equal(t, 6, Seq[int]{1, 2, 3}.Fold(0, func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 0, Seq[int]{}.Fold(0, func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 0, nilSeq[int]().Fold(0, func(a1, a2 int) int { return a1 + a2 }))

	assert.Equal(t, "hi!", Seq[string]{"h", "i", "!"}.Fold("", func(a1, a2 string) string { return a1 + a2 }))
	assert.Equal(t, "", Seq[string]{}.Fold("", func(a1, a2 string) string { return a1 + a2 }))
	assert.Equal(t, "", nilSeq[string]().Fold("", func(a1, a2 string) string { return a1 + a2 }))
}

func TestSeqFoldSeq(t *testing.T) {
	assert.Equal(t, "0123", FoldSeq(Seq[int]{1, 2, 3}, "0", func(b string, a int) string { return b + fmt.Sprint(a) }))
	assert.Equal(t, "0", FoldSeq(Seq[int]{}, "0", func(b string, a int) string { return b + fmt.Sprint(a) }))
	assert.Equal(t, "0", FoldSeq(nilSeq[int](), "0", func(b string, a int) string { return b + fmt.Sprint(a) }))
}

func TestSeqForAll(t *testing.T) {
	assert.True(t, Seq[int]{2, 4, 6, 8}.ForAll(func(a int) bool { return a%2 == 0 }))
	assert.True(t, Seq[int]{}.ForAll(func(a int) bool { return a%2 == 0 }))
	assert.True(t, nilSeq[int]().ForAll(func(a int) bool { return a%2 == 0 }))
	assert.False(t, Seq[int]{3, 5, 6, 8}.ForAll(func(a int) bool { return a%2 == 0 }))
}

func TestSeqForeach(t *testing.T) {
	//given
	seq := EmptySeq[int](5)
	//when
	Seq[int]{1, 2, 3}.Foreach(func(a int) { seq = seq.Append(a) })
	//then
	assert.Equal(t, Seq[int]{1, 2, 3}, seq)
}

func TestSeqHead(t *testing.T) {
	assert.Equal(t, "abc", Seq[string]{"abc", "def", "ghi"}.Head())
	assert.Panics(t, func() { Seq[string]{}.Head() })
	assert.Panics(t, func() { nilSeq[string]().Head() })
}

func TestSeqHeadOption(t *testing.T) {
	assert.Equal(t, Some("abc"), Seq[string]{"abc", "def", "ghi"}.HeadOption())
	assert.Equal(t, None[string](), Seq[string]{}.HeadOption())
	assert.Equal(t, None[string](), nilSeq[string]().HeadOption())
}

func TestSeqIsEmpty(t *testing.T) {
	assert.True(t, Seq[int]{}.IsEmpty())
	assert.True(t, nilSeq[int]().IsEmpty())
	assert.False(t, Seq[int]{1, 2, 3}.IsEmpty())
	assert.False(t, Seq[int]{1}.IsEmpty())
}

func TestSeqLength(t *testing.T) {
	assert.Equal(t, 3, Seq[int]{1, 2, 3}.Length())
	assert.Equal(t, 0, Seq[int]{}.Length())
	assert.Equal(t, 0, nilSeq[int]().Length())
}

func TestSeqMap(t *testing.T) {
	assert.Equal(t, Seq[string]{"a!", "b!", "c!"}, Seq[string]{"a", "b", "c"}.Map(func(a string) string { return a + "!" }))
	assert.Equal(t, Seq[string]{}, Seq[string]{}.Map(func(a string) string { return a + "!" }))
	assert.Nil(t, nilSeq[string]().Map(func(a string) string { return a + "!" }))
}

func TestSeqMapSeq(t *testing.T) {
	assert.Equal(t, Seq[string]{"1", "2", "3"}, MapSeq(Seq[int]{1, 2, 3}, func(a int) string { return fmt.Sprint(a) }))
	assert.Equal(t, Seq[string]{}, MapSeq(Seq[int]{}, func(a int) string { return fmt.Sprint(a) }))
	assert.Nil(t, MapSeq(nilSeq[int](), func(a int) string { return fmt.Sprint(a) }))
}
