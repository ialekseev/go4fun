package types

type Tuple[A, B any] struct {
	a A
	b B
}

func NewTuple[A, B any](a A, b B) Tuple[A, B] {
	return Tuple[A, B]{a, b}
}
