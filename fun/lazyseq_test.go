package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLazySeq_ContainsInLazySeq(t *testing.T) {
	assert.True(t, ContainsInLazySeq(Seq[int]{1, 2, 5}.Lazy(), 2))
	assert.False(t, ContainsInLazySeq(Seq[int]{1, 2, 5}.Lazy(), 6))
	assert.False(t, ContainsInLazySeq(Seq[int]{}.Lazy(), 2))
	assert.False(t, ContainsInLazySeq(nilSeq[int]().Lazy(), 2))
}

func TestLazySeq_Copy(t *testing.T) {
	//given
	lazySeq := Seq[int]{1, 2, 3, 4}.Lazy()
	//when
	copy := lazySeq.Copy()
	//then
	assert.Equal(t, lazySeq.KnownCapacity, copy.KnownCapacity)
	assert.Equal(t, lazySeq.NilUnderlying, copy.NilUnderlying)

	assert.Equal(t, Seq[int]{2, 4}, lazySeq.Filter(func(a int) bool { return a%2 == 0 }).Strict())
	assert.Equal(t, Seq[int]{1, 2, 3, 4}, copy.Strict())
}

func TestLazySeq_Exists(t *testing.T) {
	assert.True(t, Seq[int]{2, 4, 5}.Lazy().Exists(func(a int) bool { return a > 4 }))
	assert.False(t, Seq[int]{2, 4, 5}.Lazy().Exists(func(a int) bool { return a > 5 }))
	assert.False(t, Seq[int]{}.Lazy().Exists(func(a int) bool { return a > 0 }))
	assert.False(t, nilSeq[int]().Lazy().Exists(func(a int) bool { return a > 0 }))
}

func TestLazySeq_Filter(t *testing.T) {
	assert.Equal(t, Seq[int]{2, 4, 6}, Seq[int]{2, 3, 4, 5, 6}.Lazy().Filter(func(a int) bool { return a%2 == 0 }).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{2, 3, 4, 5, 6}.Lazy().Filter(func(a int) bool { return a > 6 }).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Lazy().Filter(func(a int) bool { return a > 0 }).Strict())
	assert.Nil(t, nilSeq[int]().Lazy().Filter(func(a int) bool { return a > 0 }).Strict())
}

func TestLazySeq_FilterNot(t *testing.T) {
	assert.Equal(t, Seq[int]{3, 5}, Seq[int]{2, 3, 4, 5, 6}.Lazy().FilterNot(func(a int) bool { return a%2 == 0 }).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{2, 3, 4, 5, 6}.Lazy().FilterNot(func(a int) bool { return a >= 2 }).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Lazy().FilterNot(func(a int) bool { return a > 0 }).Strict())
	assert.Nil(t, nilSeq[int]().Lazy().FilterNot(func(a int) bool { return a > 0 }).Strict())
}

func TestLazySeq_Find(t *testing.T) {
	assert.Equal(t, Some(3), Seq[int]{1, 2, 3, 4, 5}.Lazy().Find(func(a int) bool { return a > 2 }))
	assert.Equal(t, None[int](), Seq[int]{1, 2, 3, 4, 5}.Lazy().Find(func(a int) bool { return a > 5 }))
	assert.Equal(t, None[int](), Seq[int]{}.Lazy().Find(func(a int) bool { return a > 0 }))
	assert.Equal(t, None[int](), nilSeq[int]().Lazy().Find(func(a int) bool { return a > 0 }))
}

func TestLazySeq_FlatMap(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 1, 2, 2, 3, 3}, Seq[int]{1, 2, 3}.Lazy().FlatMap(func(a int) LazySeq[int] { return Seq[int]{a, a}.Lazy() }).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Lazy().FlatMap(func(a int) LazySeq[int] { return Seq[int]{a, a}.Lazy() }).Strict())
	assert.Nil(t, nilSeq[int]().Lazy().FlatMap(func(a int) LazySeq[int] { return Seq[int]{a, a}.Lazy() }).Strict())
}

func TestLazySeq_FlatMapLazySeq(t *testing.T) {
	assert.Equal(t, Seq[string]{"1", "1", "2", "2", "3", "3"}, FlatMapLazySeq(Seq[int]{1, 2, 3}.Lazy(), func(a int) LazySeq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)}.Lazy() }).Strict())
	assert.Equal(t, Seq[string]{}, FlatMapLazySeq(Seq[int]{}.Lazy(), func(a int) LazySeq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)}.Lazy() }).Strict())
	assert.Nil(t, FlatMapLazySeq(nilSeq[int]().Lazy(), func(a int) LazySeq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)}.Lazy() }).Strict())
}

func TestLazySeq_Fold(t *testing.T) {
	assert.Equal(t, 6, Seq[int]{1, 2, 3}.Lazy().Fold(0, func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 0, Seq[int]{}.Lazy().Fold(0, func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 0, nilSeq[int]().Lazy().Fold(0, func(a1, a2 int) int { return a1 + a2 }))

	assert.Equal(t, "hi!", Seq[string]{"h", "i", "!"}.Lazy().Fold("", func(a1, a2 string) string { return a1 + a2 }))
	assert.Equal(t, "", Seq[string]{}.Lazy().Fold("", func(a1, a2 string) string { return a1 + a2 }))
	assert.Equal(t, "", nilSeq[string]().Lazy().Fold("", func(a1, a2 string) string { return a1 + a2 }))
}

func TestLazySeq_FoldLazySeq(t *testing.T) {
	assert.Equal(t, "0123", FoldLazySeq(Seq[int]{1, 2, 3}.Lazy(), "0", func(b string, a int) string { return b + fmt.Sprint(a) }))
	assert.Equal(t, "0", FoldLazySeq(Seq[int]{}.Lazy(), "0", func(b string, a int) string { return b + fmt.Sprint(a) }))
	assert.Equal(t, "0", FoldLazySeq(nilSeq[int]().Lazy(), "0", func(b string, a int) string { return b + fmt.Sprint(a) }))
}

func TestLazySeq_ForAll(t *testing.T) {
	assert.True(t, Seq[int]{2, 4, 6, 8}.Lazy().ForAll(func(a int) bool { return a%2 == 0 }))
	assert.True(t, Seq[int]{}.Lazy().ForAll(func(a int) bool { return a%2 == 0 }))
	assert.True(t, nilSeq[int]().Lazy().ForAll(func(a int) bool { return a%2 == 0 }))
	assert.False(t, Seq[int]{3, 5, 6, 8}.Lazy().ForAll(func(a int) bool { return a%2 == 0 }))
}

func TestLazySeq_Foreach(t *testing.T) {
	//given
	seq := EmptySeq[int](5)
	//when
	Seq[int]{1, 2, 3}.Lazy().Foreach(func(a int) { seq = seq.Append(a) })
	//then
	assert.Equal(t, Seq[int]{1, 2, 3}, seq)
}

func TestLazySeq_Head(t *testing.T) {
	assert.Equal(t, "abc", Seq[string]{"abc", "def", "ghi"}.Lazy().Head())
	assert.Equal(t, "", Seq[string]{}.Lazy().Head())
	assert.Equal(t, 0, Seq[int]{}.Lazy().Head())
	assert.Equal(t, 0, nilSeq[int]().Lazy().Head())
}

func TestLazySeq_HeadOption(t *testing.T) {
	assert.Equal(t, Some("abc"), Seq[string]{"abc", "def", "ghi"}.Lazy().HeadOption())
	assert.Equal(t, None[string](), Seq[string]{}.Lazy().HeadOption())
	assert.Equal(t, None[string](), nilSeq[string]().Lazy().HeadOption())
}

func TestLazySeq_IsEmpty(t *testing.T) {
	assert.True(t, Seq[int]{}.Lazy().IsEmpty())
	assert.True(t, nilSeq[int]().Lazy().IsEmpty())
	assert.False(t, Seq[int]{1, 2, 3}.Lazy().IsEmpty())
	assert.False(t, Seq[int]{1}.Lazy().IsEmpty())
}

func TestLazySeq_LazySeqFromSeq(t *testing.T) {
	//given
	seq := Seq[int]{2, 4, 5}

	//when
	lazy := LazySeqFromSeq(seq)

	//then
	assert.NotNil(t, lazy.Iterator)
	assert.Equal(t, 3, lazy.KnownCapacity)
}

func TestLazySeq_Length(t *testing.T) {
	assert.Equal(t, 3, Seq[int]{1, 2, 3}.Lazy().Length())
	assert.Equal(t, 0, Seq[int]{}.Lazy().Length())
	assert.Equal(t, 0, nilSeq[int]().Lazy().Length())
}

func TestLazySeq_Map(t *testing.T) {
	assert.Equal(t, Seq[string]{"a!", "b!", "c!"}, Seq[string]{"a", "b", "c"}.Lazy().Map(func(a string) string { return a + "!" }).Strict())
	assert.Equal(t, Seq[string]{}, Seq[string]{}.Lazy().Map(func(a string) string { return a + "!" }).Strict())
	assert.Nil(t, nilSeq[string]().Lazy().Map(func(a string) string { return a + "!" }).Strict())
}

func TestLazySeq_MapLazySeq(t *testing.T) {
	assert.Equal(t, Seq[string]{"1", "2", "3"}, MapLazySeq(Seq[int]{1, 2, 3}.Lazy(), func(a int) string { return fmt.Sprint(a) }).Strict())
	assert.Equal(t, Seq[string]{}, MapLazySeq(Seq[int]{}.Lazy(), func(a int) string { return fmt.Sprint(a) }).Strict())
	assert.Nil(t, MapLazySeq(nilSeq[int]().Lazy(), func(a int) string { return fmt.Sprint(a) }).Strict())
}

func TestLazySeq_MaxInLazySeq(t *testing.T) {
	assert.Equal(t, 7, MaxInLazySeq(Seq[int]{-1, 4, 7, 3, -4, 0, 2}.Lazy()))
	assert.Equal(t, -2, MaxInLazySeq(Seq[int]{-2}.Lazy()))
	assert.Equal(t, 0, MaxInLazySeq(Seq[int]{}.Lazy()))
	assert.Equal(t, 0, MaxInLazySeq(nilSeq[int]().Lazy()))
}

func TestLazySeq_MinInLazySeq(t *testing.T) {
	assert.Equal(t, -7, MinInLazySeq(Seq[int]{-1, 4, -7, 3, -4, 0, 2}.Lazy()))
	assert.Equal(t, 4, MinInLazySeq(Seq[int]{4}.Lazy()))
	assert.Equal(t, 0, MinInLazySeq(Seq[int]{}.Lazy()))
	assert.Equal(t, 0, MinInLazySeq(nilSeq[int]().Lazy()))
}

func TestLazySeq_Next(t *testing.T) {
	seq := Seq[int]{2, 4, 6, 8}.Lazy()
	assert.Equal(t, Some(2), seq.Iterator.Next())
	assert.Equal(t, Some(4), seq.Iterator.Next())
	assert.Equal(t, Some(6), seq.Iterator.Next())
	assert.Equal(t, Some(8), seq.Iterator.Next())
	assert.Equal(t, None[int](), seq.Iterator.Next())
	assert.Equal(t, None[int](), seq.Iterator.Next())

	seq = Seq[int]{}.Lazy()
	assert.Equal(t, None[int](), seq.Iterator.Next())

	seq = nilSeq[int]().Lazy()
	assert.Equal(t, None[int](), seq.Iterator.Next())
}

func TestLazySeq_NonEmpty(t *testing.T) {
	assert.True(t, Seq[int]{1, 2, 3}.Lazy().NonEmpty())
	assert.True(t, Seq[int]{1}.Lazy().NonEmpty())
	assert.False(t, Seq[int]{}.Lazy().NonEmpty())
	assert.False(t, nilSeq[int]().Lazy().NonEmpty())
}

func TestLazySeq_Reduce(t *testing.T) {
	assert.Equal(t, 10, Seq[int]{1, 2, 3, 4}.Lazy().Reduce(func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 1, Seq[int]{1}.Lazy().Reduce(func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 0, Seq[int]{}.Lazy().Reduce(func(a1, a2 int) int { return a1 + a2 }))
	assert.Equal(t, 0, nilSeq[int]().Lazy().Reduce(func(a1, a2 int) int { return a1 + a2 }))
}

func TestLazySeq_Strict(t *testing.T) {
	assert.Equal(t, Seq[int]{2, 4, 6, 8}, Seq[int]{2, 4, 6, 8}.Lazy().Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Lazy().Strict())
	assert.Nil(t, nilSeq[int]().Lazy().Strict())
}

func TestLazySeq_Take(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 2}, Seq[int]{1, 2, 3, 4}.Lazy().Take(2).Strict())
	assert.Equal(t, Seq[int]{1, 2, 3, 4}, Seq[int]{1, 2, 3, 4}.Lazy().Take(10).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{1, 2, 3, 4}.Lazy().Take(0).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{1, 2, 3, 4}.Lazy().Take(-5).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Lazy().Take(2).Strict())
	assert.Nil(t, nilSeq[int]().Lazy().Take(2).Strict())
}

func TestLazySeq_ToSeq(t *testing.T) {
	assert.Equal(t, Seq[int]{2, 4, 6, 8}, Seq[int]{2, 4, 6, 8}.Lazy().ToSeq())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Lazy().ToSeq())
	assert.Nil(t, nilSeq[int]().Lazy().Strict().Lazy().ToSeq())
}

func TestLazySeq_UnZipLazySeq(t *testing.T) {
	r := UnZipLazySeq(Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b"), Tup2(3, "c")}.Lazy())
	assert.Equal(t, Tup2(Seq[int]{1, 2, 3}, Seq[string]{"a", "b", "c"}), Tup2(r.a.Strict(), r.b.Strict()))

	r = UnZipLazySeq(Seq[Tuple2[int, string]]{}.Lazy())
	assert.Equal(t, Tup2(Seq[int]{}, Seq[string]{}), Tup2(r.a.Strict(), r.b.Strict()))

	r = UnZipLazySeq(nilSeq[Tuple2[int, string]]().Lazy())
	assert.Nil(t, r.a.Strict())
	assert.Nil(t, r.b.Strict())
}

func TestLazySeq_ZipLazySeq(t *testing.T) {
	assert.Equal(t, Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b"), Tup2(3, "c")}, ZipLazySeq(Seq[int]{1, 2, 3}.Lazy(), Seq[string]{"a", "b", "c"}.Lazy()).Strict())
	assert.Equal(t, Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b")}, ZipLazySeq(Seq[int]{1, 2}.Lazy(), Seq[string]{"a", "b", "c"}.Lazy()).Strict())
	assert.Equal(t, Seq[Tuple2[int, string]]{Tup2(1, "a"), Tup2(2, "b")}, ZipLazySeq(Seq[int]{1, 2, 3}.Lazy(), Seq[string]{"a", "b"}.Lazy()).Strict())

	assert.Equal(t, Seq[Tuple2[int, string]]{}, ZipLazySeq(Seq[int]{1, 2, 3}.Lazy(), Seq[string]{}.Lazy()).Strict())
	assert.Equal(t, Seq[Tuple2[int, string]]{}, ZipLazySeq(Seq[int]{}.Lazy(), Seq[string]{"a", "b", "c"}.Lazy()).Strict())
	assert.Equal(t, Seq[Tuple2[int, string]]{}, ZipLazySeq(Seq[int]{}.Lazy(), Seq[string]{}.Lazy()).Strict())

	assert.Nil(t, ZipLazySeq(Seq[int]{1, 2, 3}.Lazy(), nilSeq[string]().Lazy()).Strict())
	assert.Nil(t, ZipLazySeq(nilSeq[int]().Lazy(), Seq[string]{"a", "b", "c"}.Lazy()).Strict())
	assert.Nil(t, ZipLazySeq(nilSeq[int]().Lazy(), nilSeq[string]().Lazy()).Strict())
}
