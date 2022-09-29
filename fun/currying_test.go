package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurryingCurry2(t *testing.T) {
	//given
	f := func(a int, b bool) string {
		return fmt.Sprint(a) + fmt.Sprint(b)
	}
	//when
	cf := Curry2(f)
	//then
	assert.Equal(t, f(1, true), cf(1)(true))
}

func TestCurryingCurry3(t *testing.T) {
	//given
	f := func(a int, b bool, c float64) string {
		return fmt.Sprint(a) + fmt.Sprint(b) + fmt.Sprint(c)
	}
	//when
	curriedF := Curry3(f)
	//then
	assert.Equal(t, f(1, true, 5.5), curriedF(1)(true)(5.5))
}

func TestCurryingUnCurry2(t *testing.T) {
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

func TestCurryingUnCurry3(t *testing.T) {
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
