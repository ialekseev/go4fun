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
