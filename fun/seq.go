package fun

type Seq[A comparable] []A

// Returns true if this sequence contains an element that is equal (as determined by ==) to elem, false otherwise.
func (seq Seq[A]) Contains(elem A) bool {
	for _, e := range seq {
		if e == elem {
			return true
		}
	}
	return false
}

// Returns a new Sequence from this Sequence without any duplicate elements.
func (seq Seq[A]) Distinct() Seq[A] {
	if seq == nil {
		return nil
	}
	m := make(map[A]struct{}, len(seq))
	r := make(Seq[A], 0, len(seq))
	for _, e := range seq {
		_, found := m[e]
		if !found {
			r = append(r, e)
		}
		m[e] = struct{}{}
	}
	return r
}

// Returns false if this Sequence is empty or nil, otherwise true if the given predicate p holds for some of the elements of this Sequence, otherwise false
func (seq Seq[A]) Exists(p func(A) bool) bool {
	for _, e := range seq {
		if p(e) {
			return true
		}
	}
	return false
}

// Returns a new Sequence consisting of all elements of this Sequence that satisfy the given predicate p. The order of the elements is preserved.
func (seq Seq[A]) Filter(p func(A) bool) Seq[A] {
	if seq == nil {
		return nil
	}
	r := make(Seq[A], 0, len(seq))
	for _, e := range seq {
		if p(e) {
			r = append(r, e)
		}
	}
	return r
}

// Returns a new Sequence consisting of all elements of this Sequence that do not satisfy the given predicate p. The order of the elements is preserved.
func (seq Seq[A]) FilterNot(p func(A) bool) Seq[A] {
	if seq == nil {
		return nil
	}
	r := make(Seq[A], 0, len(seq))
	for _, e := range seq {
		if !p(e) {
			r = append(r, e)
		}
	}
	return r
}

// Finds the first element of the Sequence satisfying a predicate, if any. Returns an option value containing the first element in the Sequence that satisfies p, or None if none exists.
func (seq Seq[A]) Find(p func(A) bool) Option[A] {
	for _, e := range seq {
		if p(e) {
			return Some(e)
		}
	}
	return None[A]()
}

// Returns a new Sequence resulting from applying the given function (f: A => Seq[A]) to each element of this Sequence and then flattening results back to Seq[A]. The original Sequence type A doesn't change.
func (seq Seq[A]) FlatMap(f func(A) Seq[A]) Seq[A] {
	return FlatMapSeq(seq, f)
}

// Returns a new Sequence resulting from applying the given function (f: A => Seq[B]) to each element of this Sequence and then flattening results to Seq[B]. The original Sequence type A could change to B.
func FlatMapSeq[A, B comparable](seq Seq[A], f func(A) Seq[B]) Seq[B] {
	if seq == nil {
		return nil
	}
	r := make(Seq[B], 0, len(seq))
	for _, e := range seq {
		subSeq := f(e)
		for _, e1 := range subSeq {
			r = append(r, e1)
		}
	}
	return r
}

// Converts this slice of sequences into a sequence formed by the elements of these sequences.
func FlattenSeq[A comparable](seq []Seq[A]) Seq[A] {
	panic("Not implemented")
}

// Folds the elements of this sequence using the specified associative binary operator.
func (seq Seq[A]) Fold(defaultValue A, f func(A) A) A {
	panic("Not implemented")
}

// Applies a binary operator to a start value and all elements of this sequence.
func foldSeq[A, B comparable](seq Seq[A], defaultValue B, f func(A) B) B {
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
func MapSeq[A, B comparable](seq Seq[A], f func(A) B) Seq[B] {
	panic("Not implemented")
}

// True if this sequence is not empty.
func (seq Seq[A]) NonEmpty() bool {
	panic("Not implemented")
}

// Converts this sequence of pairs into two sequences of the first and second half of each pair.
func UnZipSeq[A, B comparable](pair Seq[Tuple2[A, B]]) Tuple2[Seq[A], Seq[B]] {
	panic("Not implemented")
}

// Returns a sequence formed from this sequence and another sequence by combining corresponding elements in pairs.
func ZipSeq[A, B comparable](seq Seq[A], another Seq[B]) Seq[Tuple2[A, B]] {
	panic("Not implemented")
}
