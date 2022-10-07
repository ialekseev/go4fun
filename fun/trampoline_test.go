package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrampolineDoneTrampolining(t *testing.T) {
	//when
	tramp := DoneTrampolining(5)

	//then
	assert.Nil(t, tramp.call)
	assert.True(t, tramp.done)
	assert.Equal(t, 5, tramp.result)
}

func TestTrampolineMoreTrampolining(t *testing.T) {
	//when
	tramp := MoreTrampolining(func() Trampoline[int] { return DoneTrampolining(5) })

	//then
	assert.NotNil(t, tramp.call)
	assert.False(t, tramp.done)
	assert.Equal(t, 0, tramp.result)
}

func TestTrampolineRun(t *testing.T) {
	//given
	tramp := MoreTrampolining(func() Trampoline[int] { return DoneTrampolining(5) })

	//when
	r := tramp.Run()

	//then
	assert.Equal(t, 5, r)
}
