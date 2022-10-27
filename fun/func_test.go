package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncCompose2(t *testing.T) {
	//given
	f := func(a int) bool { return a != 0 }
	g := func(b bool) string { return fmt.Sprint(b) }

	//when
	h := Compose2(f, g)

	//then
	assert.Equal(t, g(f(1)), h(1))
}

func TestFuncCompose3(t *testing.T) {
	//given
	f := func(a int) string { return fmt.Sprint(a) }
	g := func(b string) bool { return b != "" }
	h := func(c bool) string { return fmt.Sprint(c) }

	//when
	j := Compose3(f, g, h)

	//then
	assert.Equal(t, h(g(f(1))), j(1))
}

func TestFuncCurry2(t *testing.T) {
	//given
	f := func(a int, b bool) string {
		return fmt.Sprint(a) + fmt.Sprint(b)
	}
	//when
	cf := Curry2(f)
	//then
	assert.Equal(t, f(1, true), cf(1)(true))
}

func TestFuncCurry3(t *testing.T) {
	//given
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + fmt.Sprint(b) + fmt.Sprint(c)
	}
	//when
	curriedF := Curry3(f)
	//then
	assert.Equal(t, f(1, true, 5.5), curriedF(1)(true)(5.5))
}

func TestFuncUnCurry2(t *testing.T) {
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

func TestFuncUnCurry3(t *testing.T) {
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
