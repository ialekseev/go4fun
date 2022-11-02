package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTuple1_Tup1(t *testing.T) {
	assert.Equal(t, Tuple1[int]{1}, Tup1(1))
}

func TestTuple2_Tup2(t *testing.T) {
	assert.Equal(t, Tuple2[int, string]{1, "abc"}, Tup2(1, "abc"))
}

func TestTuple3_Tup3(t *testing.T) {
	assert.Equal(t, Tuple3[int, string, bool]{1, "abc", true}, Tup3(1, "abc", true))
}

func TestTuple1_String(t *testing.T) {
	assert.Equal(t, "(1)", Tup1(1).String())
}

func TestTuple2_String(t *testing.T) {
	assert.Equal(t, "(1,abc)", Tup2(1, "abc").String())
}

func TestTuple3_String(t *testing.T) {
	assert.Equal(t, "(1,abc,true)", Tup3(1, "abc", true).String())
}
