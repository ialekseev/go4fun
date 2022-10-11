package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLazySeqNext(t *testing.T) {
	//given
	seq := Seq[int]([]int{2, 4, 6, 8}).Lazy()

	//then
	assert.Equal(t, Some(2), seq.Next())
	assert.Equal(t, Some(4), seq.Next())
	assert.Equal(t, Some(6), seq.Next())
	assert.Equal(t, Some(8), seq.Next())
	assert.Equal(t, None[int](), seq.Next())
}
