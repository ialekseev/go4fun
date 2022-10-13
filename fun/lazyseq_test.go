package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLazySeqFilter(t *testing.T) {
	assert.Equal(t, Seq[int]{2, 4, 6}, Seq[int]{2, 3, 4, 5, 6}.Lazy().Filter(func(a int) bool { return a%2 == 0 }).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{2, 3, 4, 5, 6}.Lazy().Filter(func(a int) bool { return a > 6 }).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Lazy().Filter(func(a int) bool { return a > 0 }).Strict())
	assert.Nil(t, nilSeq[int]().Lazy().Filter(func(a int) bool { return a > 0 }).Strict())
}

func TestLazySeqFilterNot(t *testing.T) {
	assert.Equal(t, Seq[int]{3, 5}, Seq[int]{2, 3, 4, 5, 6}.Lazy().FilterNot(func(a int) bool { return a%2 == 0 }).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{2, 3, 4, 5, 6}.Lazy().FilterNot(func(a int) bool { return a >= 2 }).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Lazy().FilterNot(func(a int) bool { return a > 0 }).Strict())
	assert.Nil(t, nilSeq[int]().Lazy().FilterNot(func(a int) bool { return a > 0 }).Strict())
}

func TestLazySeqFlatMap(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 1, 2, 2, 3, 3}, Seq[int]{1, 2, 3}.Lazy().FlatMap(func(a int) LazySeq[int] { return Seq[int]{a, a}.Lazy() }).Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Lazy().FlatMap(func(a int) LazySeq[int] { return Seq[int]{a, a}.Lazy() }).Strict())
	assert.Nil(t, nilSeq[int]().Lazy().FlatMap(func(a int) LazySeq[int] { return Seq[int]{a, a}.Lazy() }).Strict())
}

func TestLazySeqFlatMatLazySeq(t *testing.T) {
	assert.Equal(t, Seq[string]{"1", "1", "2", "2", "3", "3"}, FlatMapLazySeq(Seq[int]{1, 2, 3}.Lazy(), func(a int) LazySeq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)}.Lazy() }).Strict())
	assert.Equal(t, Seq[string]{}, FlatMapLazySeq(Seq[int]{}.Lazy(), func(a int) LazySeq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)}.Lazy() }).Strict())
	assert.Nil(t, FlatMapLazySeq(nilSeq[int]().Lazy(), func(a int) LazySeq[string] { return Seq[string]{fmt.Sprint(a), fmt.Sprint(a)}.Lazy() }).Strict())
}

func TestLazySeqLazySeqFromSeq(t *testing.T) {
	//given
	seq := Seq[int]{2, 4, 5}

	//when
	lazy := LazySeqFromSeq(seq)

	//then
	assert.NotNil(t, lazy.iterator)
	assert.Equal(t, 3, lazy.knownCapacity)
}

func TestLazySeqMap(t *testing.T) {
	assert.Equal(t, Seq[string]{"a!", "b!", "c!"}, Seq[string]{"a", "b", "c"}.Lazy().Map(func(a string) string { return a + "!" }).Strict())
	assert.Equal(t, Seq[string]{}, Seq[string]{}.Lazy().Map(func(a string) string { return a + "!" }).Strict())
	assert.Nil(t, nilSeq[string]().Lazy().Map(func(a string) string { return a + "!" }).Strict())
}

func TestLazySeqMapLazySeq(t *testing.T) {
	assert.Equal(t, Seq[string]{"1", "2", "3"}, MapLazySeq(Seq[int]{1, 2, 3}.Lazy(), func(a int) string { return fmt.Sprint(a) }).Strict())
	assert.Equal(t, Seq[string]{}, MapLazySeq(Seq[int]{}.Lazy(), func(a int) string { return fmt.Sprint(a) }).Strict())
	assert.Nil(t, MapLazySeq(nilSeq[int]().Lazy(), func(a int) string { return fmt.Sprint(a) }).Strict())
}

func TestLazySeqNext(t *testing.T) {
	seq := Seq[int]{2, 4, 6, 8}.Lazy()
	assert.Equal(t, Some(2), seq.Next())
	assert.Equal(t, Some(4), seq.Next())
	assert.Equal(t, Some(6), seq.Next())
	assert.Equal(t, Some(8), seq.Next())
	assert.Equal(t, None[int](), seq.Next())
	assert.Equal(t, None[int](), seq.Next())

	seq = Seq[int]{}.Lazy()
	assert.Equal(t, None[int](), seq.Next())

	seq = nilSeq[int]().Lazy()
	assert.Equal(t, None[int](), seq.Next())
}

func TestLazySeqStrict(t *testing.T) {
	assert.Equal(t, Seq[int]{2, 4, 6, 8}, Seq[int]{2, 4, 6, 8}.Lazy().Strict())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Lazy().Strict())
	assert.Nil(t, nilSeq[int]().Lazy().Strict())
}
