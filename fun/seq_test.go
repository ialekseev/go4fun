package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func nilSeq[T any]() Seq[T] {
	var n Seq[T] = nil
	return n
}

func TestSeq_Append(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 2, 3}, Seq[int]{1, 2}.Append(3))
	assert.Equal(t, Seq[int]{1}, Seq[int]{}.Append(1))
	assert.Equal(t, Seq[int]{1}, nilSeq[int]().Append(1))
}

func TestSeq_ContainsInSeq(t *testing.T) {
	assert.True(t, ContainsInSeq(Seq[int]{1, 2, 5}, 2))
	assert.False(t, ContainsInSeq(Seq[int]{1, 2, 5}, 6))
	assert.False(t, ContainsInSeq(Seq[int]{}, 2))
	assert.False(t, ContainsInSeq(nilSeq[int](), 2))
}

func TestSeq_Distinct(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 2, 3}, Distinct(Seq[int]{1, 1, 2, 3, 3, 3}))
	assert.Equal(t, Seq[int]{1, 2, 3}, Distinct(Seq[int]{1, 2, 3}))
	assert.Equal(t, Seq[int]{}, Distinct(Seq[int]{}))
	assert.Nil(t, Distinct(nilSeq[int]()))
}

func TestSeq_EmptySeq(t *testing.T) {
	assert.Equal(t, 0, EmptySeq[int](5).Length())
}

func TestSeq_Exists(t *testing.T) {
	assert.True(t, Seq[int]{2, 4, 5}.Exists(func(a int) bool { return a > 4 }))
	assert.False(t, Seq[int]{2, 4, 5}.Exists(func(a int) bool { return a > 5 }))
	assert.False(t, Seq[int]{}.Exists(func(a int) bool { return a > 0 }))
	assert.False(t, nilSeq[int]().Exists(func(a int) bool { return a > 0 }))
}

func TestSeq_Filter(t *testing.T) {
	assert.Equal(t, Seq[int]{2, 4, 6}, Seq[int]{2, 3, 4, 5, 6}.Filter(func(a int) bool { return a%2 == 0 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{2, 3, 4, 5, 6}.Filter(func(a int) bool { return a > 6 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Filter(func(a int) bool { return a > 0 }))
	assert.Nil(t, nilSeq[int]().Filter(func(a int) bool { return a > 0 }))
}

func TestSeq_FilterNot(t *testing.T) {
	assert.Equal(t, Seq[int]{3, 5}, Seq[int]{2, 3, 4, 5, 6}.FilterNot(func(a int) bool { return a%2 == 0 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{2, 3, 4, 5, 6}.FilterNot(func(a int) bool { return a >= 2 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{}.FilterNot(func(a int) bool { return a > 0 }))
	assert.Nil(t, nilSeq[int]().FilterNot(func(a int) bool { return a > 0 }))
}

func TestSeq_Find(t *testing.T) {
	assert.Equal(t, Some(3), Seq[int]{1, 2, 3, 4, 5}.Find(func(a int) bool { return a > 2 }))
	assert.Equal(t, None[int](), Seq[int]{1, 2, 3, 4, 5}.Find(func(a int) bool { return a > 5 }))
	assert.Equal(t, None[int](), Seq[int]{}.Find(func(a int) bool { return a > 0 }))
	assert.Equal(t, None[int](), nilSeq[int]().Find(func(a int) bool { return a > 0 }))
}

func TestSeq_FlatMap(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 1, 2, 2, 3, 3}, Seq[int]{1, 2, 3}.FlatMap(func(a int) Seq[int] { return Seq[int]{a, a} }))
	assert.Equal(t, Seq[int]{}, Seq[int]{}.FlatMap(func(a int) Seq[int] { return Seq[int]{a, a} }))
	assert.Nil(t, nilSeq[int]().FlatMap(func(a int) Seq[int] { return Seq[int]{a, a} }))
}

func TestSeq_FlatMatSeq(t *testing.T) {
	assert.Equal(t, Seq[string]{"1", "1", "2", "2", "3", "3"}, FlatMapSeq(Seq[int]{1, 2, 3}, func(a int) Seq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)} }))
	assert.Equal(t, Seq[string]{}, FlatMapSeq(Seq[int]{}, func(a int) Seq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)} }))
	assert.Nil(t, FlatMapSeq(nilSeq[int](), func(a int) Seq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)} }))
}

func TestSeq_FlattenSeq(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 2, 3, 4, 5}, FlattenSeq([]Seq[int]{{1, 2, 3}, {4, 5}}))
	assert.Equal(t, Seq[int]{}, FlattenSeq([]Seq[int]{}))

	var n []Seq[int] = nil
	assert.Nil(t, FlattenSeq(n))
}

func TestSeq_Fold(t *testing.T) {
	assert.Equal(t, 6, Seq[int]{1, 2, 3}.Fold(0, func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 0, Seq[int]{}.Fold(0, func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 0, nilSeq[int]().Fold(0, func(a1, a2 int) int { return a1 + a2 }))

	assert.Equal(t, "hi!", Seq[string]{"h", "i", "!"}.Fold("", func(a1, a2 string) string { return a1 + a2 }))
	assert.Equal(t, "", Seq[string]{}.Fold("", func(a1, a2 string) string { return a1 + a2 }))
	assert.Equal(t, "", nilSeq[string]().Fold("", func(a1, a2 string) string { return a1 + a2 }))
}

func TestSeq_FoldSeq(t *testing.T) {
	assert.Equal(t, "0123", FoldSeq(Seq[int]{1, 2, 3}, "0", func(b string, a int) string { return b + fmt.Sprint(a) }))
	assert.Equal(t, "0", FoldSeq(Seq[int]{}, "0", func(b string, a int) string { return b + fmt.Sprint(a) }))
	assert.Equal(t, "0", FoldSeq(nilSeq[int](), "0", func(b string, a int) string { return b + fmt.Sprint(a) }))
}

func TestSeq_ForAll(t *testing.T) {
	assert.True(t, Seq[int]{2, 4, 6, 8}.ForAll(func(a int) bool { return a%2 == 0 }))
	assert.True(t, Seq[int]{}.ForAll(func(a int) bool { return a%2 == 0 }))
	assert.True(t, nilSeq[int]().ForAll(func(a int) bool { return a%2 == 0 }))
	assert.False(t, Seq[int]{3, 5, 6, 8}.ForAll(func(a int) bool { return a%2 == 0 }))
}

func TestSeq_Foreach(t *testing.T) {
	//given
	seq := EmptySeq[int](5)
	//when
	Seq[int]{1, 2, 3}.Foreach(func(a int) { seq = seq.Append(a) })
	//then
	assert.Equal(t, Seq[int]{1, 2, 3}, seq)
}

func TestSeq_Head(t *testing.T) {
	assert.Equal(t, "abc", Seq[string]{"abc", "def", "ghi"}.Head())
	assert.Equal(t, "", Seq[string]{}.Head())
	assert.Equal(t, 0, Seq[int]{}.Head())
	assert.Equal(t, 0, nilSeq[int]().Head())
}

func TestSeq_HeadOption(t *testing.T) {
	assert.Equal(t, Some("abc"), Seq[string]{"abc", "def", "ghi"}.HeadOption())
	assert.Equal(t, None[string](), Seq[string]{}.HeadOption())
	assert.Equal(t, None[string](), nilSeq[string]().HeadOption())
}

func TestSeq_IsEmpty(t *testing.T) {
	assert.True(t, Seq[int]{}.IsEmpty())
	assert.True(t, nilSeq[int]().IsEmpty())
	assert.False(t, Seq[int]{1, 2, 3}.IsEmpty())
	assert.False(t, Seq[int]{1}.IsEmpty())
}

func TestSeq_Lazy(t *testing.T) {
	//given
	seq := Seq[int]{2, 4, 5}

	//when
	lazy := seq.Lazy()

	//then
	assert.NotNil(t, lazy.Iterator)
	assert.Equal(t, 3, lazy.KnownCapacity)
}

func TestSeq_Length(t *testing.T) {
	assert.Equal(t, 3, Seq[int]{1, 2, 3}.Length())
	assert.Equal(t, 0, Seq[int]{}.Length())
	assert.Equal(t, 0, nilSeq[int]().Length())
}

func TestSeq_Map(t *testing.T) {
	assert.Equal(t, Seq[string]{"a!", "b!", "c!"}, Seq[string]{"a", "b", "c"}.Map(func(a string) string { return a + "!" }))
	assert.Equal(t, Seq[string]{}, Seq[string]{}.Map(func(a string) string { return a + "!" }))
	assert.Nil(t, nilSeq[string]().Map(func(a string) string { return a + "!" }))
}

func TestSeq_MapSeq(t *testing.T) {
	assert.Equal(t, Seq[string]{"1", "2", "3"}, MapSeq(Seq[int]{1, 2, 3}, func(a int) string { return fmt.Sprint(a) }))
	assert.Equal(t, Seq[string]{}, MapSeq(Seq[int]{}, func(a int) string { return fmt.Sprint(a) }))
	assert.Nil(t, MapSeq(nilSeq[int](), func(a int) string { return fmt.Sprint(a) }))
}

func TestSeq_MaxInSeq(t *testing.T) {
	assert.Equal(t, 7, MaxInSeq(Seq[int]{-1, 4, 7, 3, -4, 0, 2}))
	assert.Equal(t, -2, MaxInSeq(Seq[int]{-2}))
	assert.Equal(t, 0, MaxInSeq(Seq[int]{}))
	assert.Equal(t, 0, MaxInSeq(nilSeq[int]()))
}

func TestSeq_MinInSeq(t *testing.T) {
	assert.Equal(t, -7, MinInSeq(Seq[int]{-1, 4, -7, 3, -4, 0, 2}))
	assert.Equal(t, 4, MinInSeq(Seq[int]{4}))
	assert.Equal(t, 0, MinInSeq(Seq[int]{}))
	assert.Equal(t, 0, MinInSeq(nilSeq[int]()))
}

func TestSeq_NonEmpty(t *testing.T) {
	assert.True(t, Seq[int]{1, 2, 3}.NonEmpty())
	assert.True(t, Seq[int]{1}.NonEmpty())
	assert.False(t, Seq[int]{}.NonEmpty())
	assert.False(t, nilSeq[int]().NonEmpty())
}

func TestSeq_Reduce(t *testing.T) {
	assert.Equal(t, 10, Seq[int]{1, 2, 3, 4}.Reduce(func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 1, Seq[int]{1}.Reduce(func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 0, Seq[int]{}.Reduce(func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 0, nilSeq[int]().Reduce(func(a1, a2 int) int { return a1 + a2 }))
}

func TestSeq_Take(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 2}, Seq[int]{1, 2, 3, 4}.Take(2))
	assert.Equal(t, Seq[int]{1, 2, 3, 4}, Seq[int]{1, 2, 3, 4}.Take(10))
	assert.Equal(t, Seq[int]{}, Seq[int]{1, 2, 3, 4}.Take(0))
	assert.Equal(t, Seq[int]{}, Seq[int]{1, 2, 3, 4}.Take(-5))
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Take(2))
	assert.Nil(t, nilSeq[int]().Take(2))
}

func TestSeq_UnZipSeq(t *testing.T) {
	assert.Equal(t, Tup2(Seq[int]{1, 2, 3}, Seq[string]{"a", "b", "c"}), UnZipSeq(Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b"), Tup2(3, "c")}))
	assert.Equal(t, Tup2(Seq[int]{}, Seq[string]{}), UnZipSeq(Seq[Tuple2[int, string]]{}))

	r := UnZipSeq(nilSeq[Tuple2[int, string]]())
	assert.Nil(t, r.a)
	assert.Nil(t, r.b)
}

func TestSeq_ZipSeq(t *testing.T) {
	assert.Equal(t, Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b"), Tup2(3, "c")}, ZipSeq(Seq[int]{1, 2, 3}, Seq[string]{"a", "b", "c"}))
	assert.Equal(t, Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b")}, ZipSeq(Seq[int]{1, 2}, Seq[string]{"a", "b", "c"}))
	assert.Equal(t, Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b")}, ZipSeq(Seq[int]{1, 2, 3}, Seq[string]{"a", "b"}))

	assert.Equal(t, Seq[Tuple2[int, string]]{}, ZipSeq(Seq[int]{1, 2, 3}, Seq[string]{}))
	assert.Equal(t, Seq[Tuple2[int, string]]{}, ZipSeq(Seq[int]{}, Seq[string]{"a", "b", "c"}))
	assert.Equal(t, Seq[Tuple2[int, string]]{}, ZipSeq(Seq[int]{}, Seq[string]{}))

	assert.Nil(t, ZipSeq(Seq[int]{1, 2, 3}, nilSeq[string]()))
	assert.Nil(t, ZipSeq(nilSeq[int](), Seq[string]{"a", "b", "c"}))
	assert.Nil(t, ZipSeq(nilSeq[int](), nilSeq[string]()))
}
