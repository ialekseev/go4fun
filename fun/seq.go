package fun

type Seq[A any] []A

// Appends an element to the Sequence. An alias for built-in append function.
func (seq Seq[A]) Append(elem A) Seq[A] {
	return append(seq, elem)
}

// Returns true if this sequence contains an element that is equal (as determined by ==) to elem, false otherwise.
func ContainsInSeq[A comparable](seq Seq[A], elem A) bool {
	for _, e := range seq {
		if e == elem {
			return true
		}
	}
	return false
}

// Returns a new Sequence from this Sequence without any duplicate elements.
func Distinct[A comparable](seq Seq[A]) Seq[A] {
	if seq == nil {
		return nil
	}
	m := make(map[A]struct{}, seq.Length())
	r := EmptySeq[A](seq.Length())
	for _, e := range seq {
		_, found := m[e]
		if !found {
			r = r.Append(e)
		}
		m[e] = struct{}{}
	}
	return r
}

// Creates a new empty Sequence of provided capacity. An underlying slice would have 0 length.
func EmptySeq[A any](capacity int) Seq[A] {
	return make(Seq[A], 0, capacity)
}

// Returns false if this Sequence is empty or nil, otherwise true if the given predicate p holds for some of the elements of this Sequence, otherwise false
func (seq Seq[A]) Exists(p func(A) bool) bool {
	return seq.Find(p).IsDefined()
}

// Returns a new Sequence consisting of all elements of this Sequence that satisfy the given predicate p. The order of the elements is preserved.
func (seq Seq[A]) Filter(p func(A) bool) Seq[A] {
	if seq == nil {
		return nil
	}
	r := EmptySeq[A](seq.Length())
	for _, e := range seq {
		if p(e) {
			r = r.Append(e)
		}
	}
	return r
}

// Returns a new Sequence consisting of all elements of this Sequence that do not satisfy the given predicate p. The order of the elements is preserved.
func (seq Seq[A]) FilterNot(p func(A) bool) Seq[A] {
	if seq == nil {
		return nil
	}
	r := EmptySeq[A](seq.Length())
	for _, e := range seq {
		if !p(e) {
			r = r.Append(e)
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
func FlatMapSeq[A, B any](seq Seq[A], f func(A) Seq[B]) Seq[B] {
	if seq == nil {
		return nil
	}
	r := EmptySeq[B](seq.Length())
	for _, e := range seq {
		subSeq := f(e)
		for _, e1 := range subSeq {
			r = r.Append(e1)
		}
	}
	return r
}

// Converts this slice of Sequences into a Sequence formed by the elements of these Sequences (returns a Sequence resulting from concatenating all element Sequences).
func FlattenSeq[A any](seq []Seq[A]) Seq[A] {
	if seq == nil {
		return nil
	}
	r := EmptySeq[A](len(seq))
	for _, subSeq := range seq {
		for _, e1 := range subSeq {
			r = r.Append(e1)
		}
	}
	return r
}

// Applies a binary operator op to a start value z (of type A) and all Sequence elements (of type A), going left to right. The accumulation result also keeps the same type A.
func (seq Seq[A]) Fold(z A, op func(A, A) A) A {
	return FoldSeq(seq, z, op)
}

// Applies a binary operator op to a start value z (of type B) and all Sequence elements (of type A), going left to right. The accumulation result is of type B.
func FoldSeq[A, B any](seq Seq[A], z B, op func(B, A) B) B {
	r := z
	for _, e := range seq {
		r = op(r, e)
	}
	return r
}

// Returns true if this Sequence is empty or nil or the given predicate p holds for all elements of this Sequence, otherwise false.
func (seq Seq[A]) ForAll(p func(A) bool) bool {
	return !seq.Exists(func(a A) bool { return !p(a) })
}

// Applies a given procedure f to all elements of this Sequence.
func (seq Seq[A]) Foreach(f func(A)) {
	for _, e := range seq {
		f(e)
	}
}

// Returns the first element of this Sequence. Returns type A's default value if the Sequence is empty or nil.
func (seq Seq[A]) Head() A {
	if seq.NonEmpty() {
		return seq[0]
	}
	return *new(A)
}

// Returns the first element of this Sequence if it is nonempty, None if it is empty or nil.
func (seq Seq[A]) HeadOption() Option[A] {
	if seq.NonEmpty() {
		return Some(seq[0])
	} else {
		return None[A]()
	}
}

// Returns true if the Sequence contain no elements, false otherwise.
func (seq Seq[A]) IsEmpty() bool {
	return seq.Length() == 0
}

// Converts this Sequence to Lazy Sequence
func (seq Seq[A]) Lazy() LazySeq[A] {
	return LazySeqFromSeq(seq)
}

// Returns the length of the Sequence. An alias for built-in len function.
func (seq Seq[A]) Length() int {
	return len(seq)
}

// Returns a new Sequence resulting from applying the given function f to each element of this Sequence and collecting the results (without changing type A of the Sequence elements).
func (seq Seq[A]) Map(f func(A) A) Seq[A] {
	return MapSeq(seq, f)
}

// Returns a new Sequence resulting from applying the given function f to each element of this Sequence and collecting the results (potentially, changing type A of the Sequence elements to B).
func MapSeq[A, B any](seq Seq[A], f func(A) B) Seq[B] {
	if seq == nil {
		return nil
	}
	r := EmptySeq[B](seq.Length())
	for _, e := range seq {
		r = r.Append(f(e))
	}
	return r
}

// Returns the largest element of this Sequence. Or a default value of type A if the Sequence is empty or nil.
func MaxInSeq[A Ordered](seq Seq[A]) A {
	return seq.Reduce(func(a1, a2 A) A {
		if a1 > a2 {
			return a1
		} else {
			return a2
		}
	})
}

// Returns the smallest element of this Sequence. Or a default value of type A if the Sequence is empty or nil.
func MinInSeq[A Ordered](seq Seq[A]) A {
	return seq.Reduce(func(a1, a2 A) A {
		if a1 < a2 {
			return a1
		} else {
			return a2
		}
	})
}

// Returns true if the Sequence contains at least one element, false otherwise.
func (seq Seq[A]) NonEmpty() bool {
	return !seq.IsEmpty()
}

// Returns a result of applying reduce operator op between all the elements of the Sequence, going left to right. If the Sequence is empty or nil then a default value of type A is returned.
func (seq Seq[A]) Reduce(op func(A, A) A) A {
	if seq.IsEmpty() {
		return *new(A)
	}
	r := seq.Head()
	for i := 1; i < seq.Length(); i++ {
		r = op(r, seq[i])
	}
	return r
}

// Converts this Sequence of Tuples into a Tuple of two Sequences.
func UnZipSeq[A, B any](seq Seq[Tuple2[A, B]]) Tuple2[Seq[A], Seq[B]] {
	if seq == nil {
		return Tuple2[Seq[A], Seq[B]]{nil, nil}
	}
	seqA := EmptySeq[A](seq.Length())
	seqB := EmptySeq[B](seq.Length())
	for _, e := range seq {
		seqA = seqA.Append(e.a)
		seqB = seqB.Append(e.b)
	}
	return Tup2(seqA, seqB)
}

// Returns a new Sequence formed from this Sequence and another Sequence by combining corresponding elements in Tuples. If one of the two collections is longer than the other, its remaining elements are ignored.
func ZipSeq[A, B any](seq Seq[A], another Seq[B]) Seq[Tuple2[A, B]] {
	if seq == nil || another == nil {
		return nil
	}

	minLen := seq.Length()
	if another.Length() < minLen {
		minLen = another.Length()
	}
	r := EmptySeq[Tuple2[A, B]](minLen)
	for i := 0; i < minLen; i++ {
		r = r.Append(Tup2(seq[i], another[i]))
	}
	return r
}
