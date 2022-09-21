package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTup1(t *testing.T) {
	assert.Equal(t, Tuple1[int]{1}, Tup1(1))
}

func TestTup2(t *testing.T) {
	assert.Equal(t, Tuple2[int, string]{1, "abc"}, Tup2(1, "abc"))
}

func TestTup3(t *testing.T) {
	assert.Equal(t, Tuple3[int, string, bool]{1, "abc", true}, Tup3(1, "abc", true))
}

func TestString1(t *testing.T) {
	assert.Equal(t, "(1)", Tup1(1).String())
}

func TestString2(t *testing.T) {
	assert.Equal(t, "(1,abc)", Tup2(1, "abc").String())
}

func TestString3(t *testing.T) {
	assert.Equal(t, "(1,abc,true)", Tup3(1, "abc", true).String())
}
