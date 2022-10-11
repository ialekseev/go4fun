package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLazySeqMap(t *testing.T) {

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
