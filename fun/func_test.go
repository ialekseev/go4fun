package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc_Apply2(t *testing.T) {
	//given
	f := func(a int, b bool) string {
		return fmt.Sprint(a) + fmt.Sprint(b)
	}
	//when
	r := Apply2(f, 10, true)
	//then
	assert.Equal(t, "10true", r)
}

func TestFunc_Apply2Partial_1(t *testing.T) {
	//given
	f := func(a int, b bool) string {
		return fmt.Sprint(a) + fmt.Sprint(b)
	}
	//when
	p := Apply2Partial_1(f, 10)
	//then
	assert.Equal(t, f(10, true), p(true))
}

func TestFunc_Apply2Partial_2(t *testing.T) {
	//given
	f := func(a int, b bool) string {
		return fmt.Sprint(a) + fmt.Sprint(b)
	}
	//when
	p := Apply2Partial_2(f, true)
	//then
	assert.Equal(t, f(10, true), p(10))
}

func TestFunc_Apply3(t *testing.T) {
	//given
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c)
	}
	//when
	r := Apply3(f, 10, true, 5.5)
	//then
	assert.Equal(t, "10 true 5.5", r)
}

func TestFunc_Apply3Partial_1(t *testing.T) {
	//given
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + fmt.Sprint(b) + fmt.Sprint(c)
	}
	//when
	p := Apply3Partial_1(f, 10)
	//then
	assert.Equal(t, f(10, true, 5.5), p(true, 5.5))
}

func TestFunc_Apply3Partial_2(t *testing.T) {
	//given
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + fmt.Sprint(b) + fmt.Sprint(c)
	}
	//when
	p := Apply3Partial_2(f, true)
	//then
	assert.Equal(t, f(10, true, 5.5), p(10, 5.5))
}

func TestFunc_Apply3Partial_3(t *testing.T) {
	//given
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + fmt.Sprint(b) + fmt.Sprint(c)
	}
	//when
	p := Apply3Partial_3(f, 5.5)
	//then
	assert.Equal(t, f(10, true, 5.5), p(10, true))
}

func TestFunc_Apply3Partial_1_2(t *testing.T) {
	//given
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + fmt.Sprint(b) + fmt.Sprint(c)
	}
	//when
	p := Apply3Partial_1_2(f, 10, true)
	//then
	assert.Equal(t, f(10, true, 5.5), p(5.5))
}

func TestFunc_Apply3Partial_1_3(t *testing.T) {
	//given
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + fmt.Sprint(b) + fmt.Sprint(c)
	}
	//when
	p := Apply3Partial_1_3(f, 10, 5.5)
	//then
	assert.Equal(t, f(10, true, 5.5), p(true))
}

func TestFunc_Apply3Partial_2_3(t *testing.T) {
	//given
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + fmt.Sprint(b) + fmt.Sprint(c)
	}
	//when
	p := Apply3Partial_2_3(f, true, 5.5)
	//then
	assert.Equal(t, f(10, true, 5.5), p(10))
}

func TestFunc_Compose2(t *testing.T) {
	//given
	f := func(a int) bool { return a != 0 }
	g := func(b bool) string { return fmt.Sprint(b) }
	//when
	h := Compose2(f, g)
	//then
	assert.Equal(t, g(f(1)), h(1))
}

func TestFunc_Compose3(t *testing.T) {
	//given
	f := func(a int) string { return fmt.Sprint(a) }
	g := func(b string) bool { return b != "" }
	h := func(c bool) string { return fmt.Sprint(c) }
	//when
	j := Compose3(f, g, h)
	//then
	assert.Equal(t, h(g(f(1))), j(1))
}

func TestFunc_Curry2(t *testing.T) {
	//given
	f := func(a int, b bool) string {
		return fmt.Sprint(a) + fmt.Sprint(b)
	}
	//when
	cf := Curry2(f)
	//then
	assert.Equal(t, f(1, true), cf(1)(true))
}

func TestFunc_Curry3(t *testing.T) {
	//given
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + fmt.Sprint(b) + fmt.Sprint(c)
	}
	//when
	curriedF := Curry3(f)
	//then
	assert.Equal(t, f(1, true, 5.5), curriedF(1)(true)(5.5))
}

func TestFunc_Tupled2(t *testing.T) {
	//given
	f := func(a int, b bool) string {
		return fmt.Sprint(a) + fmt.Sprint(b)
	}
	//when
	tupledF := Tupled2(f)
	//then
	assert.Equal(t, f(1, true), tupledF(Tup2(1, true)))
}

func TestFunc_Tupled3(t *testing.T) {
	//given
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + fmt.Sprint(b) + fmt.Sprint(c)
	}
	//when
	tupledF := Tupled3(f)
	//then
	assert.Equal(t, f(1, true, 5.5), tupledF(Tup3(1, true, 5.5)))
}

func TestFunc_UnCurry2(t *testing.T) {
	//given
	f := func(a int) func(bool) string {
		return func(b bool) string {
			return fmt.Sprint(a) + fmt.Sprint(b)
		}
	}
	//when
	unCurriedF := UnCurry2(f)
	//then
	assert.Equal(t, f(1)(true), unCurriedF(1, true))
}

func TestFunc_UnCurry3(t *testing.T) {
	//given
	f := func(a int) func(bool) func(float64) string {
		return func(b bool) func(float64) string {
			return func(c float64) string {
				return fmt.Sprint(a) + fmt.Sprint(b) + fmt.Sprint(c)
			}
		}
	}
	//when
	unCurriedF := UnCurry3(f)
	//then
	assert.Equal(t, f(1)(true)(5.5), unCurriedF(1, true, 5.5))
}

func TestFunc_UnTupled2(t *testing.T) {
	//given
	f := func(t Tuple2[int, bool]) string {
		return fmt.Sprint(t.a) + fmt.Sprint(t.b)
	}
	//when
	unTupledF := UnTupled2(f)
	//then
	assert.Equal(t, f(Tup2(1, true)), unTupledF(1, true))
}

func TestFunc_UnTupled3(t *testing.T) {
	//given
	f := func(t Tuple3[int, bool, float64]) string {
		return fmt.Sprint(t.a) + fmt.Sprint(t.b) + fmt.Sprint(t.c)
	}
	//when
	unTupledF := UnTupled3(f)
	//then
	assert.Equal(t, f(Tup3(1, true, 5.5)), unTupledF(1, true, 5.5))
}
