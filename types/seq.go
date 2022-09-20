package types

type Seq[A any] []A

// Tests whether this sequence contains a given value as an element.
func (seq Seq[A]) Contains(elem A) bool {
	panic("Not implemented")
}

// Builds a new sequence from this sequence without any duplicate elements.
func (seq Seq[T]) Distinct(slice []T) []T {
	panic("Not implemented")
}

// Tests whether a predicate holds for at least one element of this sequence.
func (option Seq[A]) Exists(f func(A) bool) bool {
	panic("Not implemented")
}

// Selects all elements of this sequence which satisfy a predicate.
func (seq Seq[T]) Filter(f func(T) bool) []T {
	panic("Not implemented")
}

// Selects all elements of this sequence which do not satisfy a predicate.
func (seq Seq[T]) FilterNot(f func(T) bool) []T {
	panic("Not implemented")
}

// Finds the first element of the sequence satisfying a predicate, if any.
func (seq Seq[T]) Find(f func(T) bool) (T, bool) {
	for _, e := range seq {
		if f(e) {
			return e, true
		}
	}
	return *new(T), false
}

// Builds a new sequence by applying a function to all elements of this sequence and using the elements of the resulting sequences.
func (seq Seq[T]) FlatMap(f func(T) []T) []T {
	panic("Not implemented")
}

func FlatMapSeq[A, B any](seq Seq[A], f func(A) Seq[B]) Seq[B] {
	panic("Not implemented")
}

// Converts this sequence of sequences into a sequence formed by the elements of these sequences.
func FlattenSeq[A any](seq Seq[Seq[A]]) Seq[A] {
	panic("Not implemented")
}

// Folds the elements of this sequence using the specified associative binary operator.
func (seq Seq[A]) Fold(defaultValue A, f func(A) A) A {
	panic("Not implemented")
}

// Applies a binary operator to a start value and all elements of this sequence.
func foldSeq[A, B any](seq Seq[A], defaultValue B, f func(A) B) B {
	panic("Not implemented")
}

// Tests whether a predicate holds for all elements of this sequence.
func (seq Seq[A]) ForAll(f func(A) bool) bool {
	panic("Not implemented")
}

// Applies a given procedure f to all elements of this sequence.
func (seq Seq[A]) Foreach(f func(A)) {
	panic("Not implemented")
}

// Selects the first element of this sequence.
func (seq Seq[A]) Head() A {
	panic("Not implemented")
}

// Optionally selects the first element.
func (seq Seq[A]) HeadOption() Option[A] {
	panic("Not implemented")
}

// True if this sequence is empty
func (seq Seq[A]) IsEmpty() bool {
	panic("Not implemented")
}

// Builds a new sequence by applying a function to all elements of this sequence.
func (seq Seq[A]) Map(f func(A) A) A {
	panic("Not implemented")
}

// Builds a new sequence by applying a function to all elements of this sequence.
func MapSeq[A, B any](seq Seq[A], f func(A) B) Seq[B] {
	panic("Not implemented")
}

// True if this sequence is not empty.
func (seq Seq[A]) NonEmpty() bool {
	panic("Not implemented")
}

// Returns an underlying slice of this sequence
func (seq Seq[A]) ToSlice() []A {
	panic("Not implemented")
}

// Converts this sequence of pairs into two sequences of the first and second half of each pair.
func UnZipSeq[A, B any](pair Seq[Tuple[A, B]]) Tuple[Seq[A], Seq[B]] {
	panic("Not implemented")
}

// Returns a sequence formed from this sequence and another sequence by combining corresponding elements in pairs.
func ZipSeq[A, B any](seq Seq[A], another Seq[B]) Seq[Tuple[A, B]] {
	panic("Not implemented")
}
