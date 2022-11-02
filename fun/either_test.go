package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEither_FlatMap(t *testing.T) {
	assert.Equal(t, Right[int]("route60"), Right[int]("60").FlatMap(func(r string) Either[int, string] { return Right[int]("route" + r) }))
	assert.Equal(t, Left[int, string](-1), Right[int]("60").FlatMap(func(r string) Either[int, string] { return Left[int, string](-1) }))
	assert.Equal(t, Left[int, string](-1), Left[int, string](-1).FlatMap(func(r string) Either[int, string] { return Right[int]("route" + r) }))
}

func TestEither_FlatMapEither(t *testing.T) {
	assert.Equal(t, Right[bool]("route60"), FlatMapEither(Right[bool](60), func(r int) Either[bool, string] { return Right[bool]("route" + fmt.Sprint(r)) }))
	assert.Equal(t, Left[bool, string](false), FlatMapEither(Right[bool](60), func(r int) Either[bool, string] { return Left[bool, string](false) }))
	assert.Equal(t, Left[bool, string](false), FlatMapEither(Left[bool, int](false), func(r int) Either[bool, string] { return Right[bool]("route" + fmt.Sprint(r)) }))
}

func TestEither_IsLeft(t *testing.T) {
	assert.True(t, Left[int, string](-1).IsLeft())
	assert.False(t, Right[int]("john lennon").IsLeft())
}

func TestEither_IsRight(t *testing.T) {
	assert.True(t, Right[int]("john lennon").IsRight())
	assert.False(t, Left[int, string](-1).IsRight())
}

func TestEither_Left(t *testing.T) {
	l := Left[int, string](-1)
	assert.True(t, l.a.IsDefined())
	assert.True(t, l.b.IsEmpty())
	assert.Equal(t, -1, l.a.Get())
}

func TestEither_LeftOption(t *testing.T) {
	assert.Equal(t, Some(-1), Left[int, string](-1).LeftOption())
	assert.Equal(t, None[int](), Right[int]("john lennon").LeftOption())
}

func TestEither_Map(t *testing.T) {
	assert.Equal(t, Right[int]("route60"), Right[int]("60").Map(func(r string) string { return "route" + r }))
	assert.Equal(t, Left[int, string](-1), Left[int, string](-1).Map(func(r string) string { return "route" + r }))
}

func TestEither_MapEither(t *testing.T) {
	assert.Equal(t, Right[bool]("route60"), MapEither(Right[bool](60), func(r int) string { return "route" + fmt.Sprint(r) }))
	assert.Equal(t, Left[bool, string](false), MapEither(Left[bool, int](false), func(r int) string { return "route" + fmt.Sprint(r) }))
}

func TestEither_Right(t *testing.T) {
	r := Right[int]("john lennon")
	assert.True(t, r.a.IsEmpty())
	assert.True(t, r.b.IsDefined())
	assert.Equal(t, "john lennon", r.b.Get())
}

func TestEither_RightOption(t *testing.T) {
	assert.Equal(t, Some("john lennon"), Right[int]("john lennon").RightOption())
	assert.Equal(t, None[string](), Left[int, string](-1).RightOption())
}

func TestEither_String(t *testing.T) {
	assert.Equal(t, "Right(john lennon)", Right[int]("john lennon").String())
	assert.Equal(t, "Left(-1)", Left[int, string](-1).String())
}

func TestEither_ToOption(t *testing.T) {
	assert.Equal(t, Some("john lennon"), Right[int]("john lennon").ToOption())
	assert.Equal(t, None[string](), Left[int, string](-1).ToOption())
}
