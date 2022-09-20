package fun

import (
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

func TestDistinct(t *testing.T) {
	assert.Equal(t, Seq[int]{1, 2, 3}, Seq[int]{1, 1, 2, 3, 3, 3}.Distinct())
	assert.Equal(t, Seq[int]{1, 2, 3}, Seq[int]{1, 2, 3}.Distinct())
	assert.Equal(t, Seq[int]{}, Seq[int]{}.Distinct())
	assert.Nil(t, nilSeq[int]().Distinct())
}
