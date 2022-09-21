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

func TestSeqExists(t *testing.T) {
	assert.True(t, Seq[int]{2, 4, 5}.Exists(func(v int) bool { return v > 4 }))
	assert.False(t, Seq[int]{2, 4, 5}.Exists(func(v int) bool { return v > 5 }))
	assert.False(t, Seq[int]{}.Exists(func(v int) bool { return v > 0 }))
	assert.False(t, nilSeq[int]().Exists(func(v int) bool { return v > 0 }))
}

func TestSeqFilter(t *testing.T) {
	assert.Equal(t, Seq[int]{2, 4, 6}, Seq[int]{2, 3, 4, 5, 6}.Filter(func(v int) bool { return v%2 == 0 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{2, 3, 4, 5, 6}.Filter(func(v int) bool { return v > 6 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Filter(func(v int) bool { return v > 0 }))
	assert.Nil(t, nilSeq[int]().Filter(func(v int) bool { return v > 0 }))
}

func TestSeqFilterNot(t *testing.T) {
	assert.Equal(t, Seq[int]{3, 5}, Seq[int]{2, 3, 4, 5, 6}.FilterNot(func(v int) bool { return v%2 == 0 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{2, 3, 4, 5, 6}.FilterNot(func(v int) bool { return v >= 2 }))
	assert.Equal(t, Seq[int]{}, Seq[int]{}.FilterNot(func(v int) bool { return v > 0 }))
	assert.Nil(t, nilSeq[int]().FilterNot(func(v int) bool { return v > 0 }))
}

func TestSeqFind(t *testing.T) {
	assert.Equal(t, Some(3), Seq[int]{1, 2, 3, 4, 5}.Find(func(v int) bool { return v > 2 }))
	assert.Equal(t, None[int](), Seq[int]{1, 2, 3, 4, 5}.Find(func(v int) bool { return v > 5 }))
	assert.Equal(t, None[int](), Seq[int]{}.Find(func(v int) bool { return v > 0 }))
	assert.Equal(t, None[int](), nilSeq[int]().Find(func(v int) bool { return v > 0 }))
}

func TestSeqFlatMap(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 1, 2, 2, 3, 3}, Seq[int]{1, 2, 3}.FlatMap(func(v int) Seq[int] { return Seq[int]{v, v} }))
	assert.Equal(t, Seq[int]{}, Seq[int]{}.FlatMap(func(v int) Seq[int] { return Seq[int]{v, v} }))
	assert.Nil(t, nilSeq[int]().FlatMap(func(v int) Seq[int] { return Seq[int]{v, v} }))
}

func TestSeqFlatMatSeq(t *testing.T) {
	assert.Equal(t, Seq[string]{"1", "1", "2", "2", "3", "3"}, FlatMapSeq(Seq[int]{1, 2, 3}, func(v int) Seq[string] { return Seq[string]{fmt.Sprint(v), fmt.Sprint(v)} }))
	assert.Equal(t, Seq[string]{}, FlatMapSeq(Seq[int]{}, func(v int) Seq[string] { return Seq[string]{fmt.Sprint(v), fmt.Sprint(v)} }))
	assert.Nil(t, FlatMapSeq(nilSeq[int](), func(v int) Seq[string] { return Seq[string]{fmt.Sprint(v), fmt.Sprint(v)} }))
}
