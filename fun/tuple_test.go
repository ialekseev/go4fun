package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTuple1(t *testing.T) {
	assert.Equal(t, Tuple1[int]{1}, NewTuple1(1))
}

func TestNewTuple2(t *testing.T) {
	assert.Equal(t, Tuple2[int, string]{1, "abc"}, NewTuple2(1, "abc"))
}

func TestNewTuple3(t *testing.T) {
	assert.Equal(t, Tuple3[int, string, bool]{1, "abc", true}, NewTuple3(1, "abc", true))
}

func TestString1(t *testing.T) {
	assert.Equal(t, "(1)", NewTuple1(1).String())
}

func TestString2(t *testing.T) {
	assert.Equal(t, "(1,abc)", NewTuple2(1, "abc").String())
}

func TestString3(t *testing.T) {
	assert.Equal(t, "(1,abc,true)", NewTuple3(1, "abc", true).String())
}
