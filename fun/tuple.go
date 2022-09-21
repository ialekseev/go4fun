package fun

import "fmt"

type Tuple1[A any] struct {
	a A
}

type Tuple2[A, B any] struct {
	a A
	b B
}

type Tuple3[A, B, C any] struct {
	a A
	b B
	c C
}

func Tup1[A any](a A) Tuple1[A] {
	return Tuple1[A]{a}
}

func Tup2[A, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{a, b}
}

func Tup3[A, B, C any](a A, b B, c C) Tuple3[A, B, C] {
	return Tuple3[A, B, C]{a, b, c}
}

func (tuple Tuple1[A]) String() string {
	return fmt.Sprintf("(%v)", tuple.a)
}

func (tuple Tuple2[A, B]) String() string {
	return fmt.Sprintf("(%v,%v)", tuple.a, tuple.b)
}

func (tuple Tuple3[A, B, C]) String() string {
	return fmt.Sprintf("(%v,%v,%v)", tuple.a, tuple.b, tuple.c)
}
