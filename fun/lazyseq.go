package fun

//-------Iterator------------

type Iterator[A any] interface {
	Next() Option[A]
	Copy() Iterator[A]
	Reset()
}

//-------seqIterator----------

type seqIterator[A any] struct {
	seq          Seq[A]
	currentIndex int
}

func (iterator *seqIterator[A]) Next() Option[A] {
	if iterator.currentIndex < iterator.seq.Length() {
		current := iterator.seq[iterator.currentIndex]
		iterator.currentIndex = iterator.currentIndex + 1
		return Some(current)
	} else {
		return None[A]()
	}
}

func (iterator *seqIterator[A]) Copy() Iterator[A] {
	return &seqIterator[A]{iterator.seq, 0}
}

func (iterator *seqIterator[A]) Reset() {
	iterator.currentIndex = 0
}

//-------filterIterator----------

type filterIterator[A any] struct {
	inputIterator Iterator[A]
	filterF       func(A) bool
}

func (iterator *filterIterator[A]) Next() Option[A] {
	for {
		next := iterator.inputIterator.Next()
		switch {
		case next.IsDefined() && iterator.filterF(next.Get()):
			return next
		case next.IsDefined() && !iterator.filterF(next.Get()):
			continue
		default:
			return None[A]()
		}
	}
}

func (iterator *filterIterator[A]) Copy() Iterator[A] {
	return &filterIterator[A]{iterator.inputIterator.Copy(), iterator.filterF}
}

func (iterator *filterIterator[A]) Reset() {
	iterator.inputIterator.Reset()
}

//-------filterIndexIterator----------

type filterIndexIterator[A any] struct {
	inputIterator Iterator[A]
	filterIndexF  func(int) bool
	currentIndex  int
}

func (iterator *filterIndexIterator[A]) Next() Option[A] {
	if iterator.filterIndexF(iterator.currentIndex) {
		iterator.currentIndex = iterator.currentIndex + 1
		return iterator.inputIterator.Next()
	} else {
		return None[A]()
	}
}

func (iterator *filterIndexIterator[A]) Copy() Iterator[A] {
	return &filterIndexIterator[A]{iterator.inputIterator.Copy(), iterator.filterIndexF, 0}
}

func (iterator *filterIndexIterator[A]) Reset() {
	iterator.inputIterator.Reset()
	iterator.currentIndex = 0
}

//-------mapIterator----------

type mapIterator[A, B any] struct {
	inputIterator Iterator[A]
	mapF          func(A) B
}

func (iterator *mapIterator[A, B]) Next() Option[B] {
	return MapOption(iterator.inputIterator.Next(), iterator.mapF)
}

func (iterator *mapIterator[A, B]) Copy() Iterator[B] {
	return &mapIterator[A, B]{iterator.inputIterator.Copy(), iterator.mapF}
}

func (iterator *mapIterator[A, B]) Reset() {
	iterator.inputIterator.Reset()
}

//-------flatMapIterator----------

type flatMapIterator[A, B any] struct {
	inputIterator Iterator[A]
	flatMapF      func(A) Iterator[B]
	fmIterator    Iterator[B]
}

func (iterator *flatMapIterator[A, B]) setNewFmIteratorAndMove() Option[B] {
	next := iterator.inputIterator.Next()
	if next.IsDefined() {
		iterator.fmIterator = iterator.flatMapF(next.Get())
		return iterator.fmIterator.Next()
	}
	return None[B]()
}

func (iterator *flatMapIterator[A, B]) Next() Option[B] {
	if iterator.fmIterator == nil {
		return iterator.setNewFmIteratorAndMove()
	} else {
		nextFm := iterator.fmIterator.Next()
		if nextFm.IsEmpty() {
			return iterator.setNewFmIteratorAndMove()
		} else {
			return nextFm
		}
	}
}

func (iterator *flatMapIterator[A, B]) Copy() Iterator[B] {
	return &flatMapIterator[A, B]{iterator.inputIterator.Copy(), iterator.flatMapF, nil}
}

func (iterator *flatMapIterator[A, B]) Reset() {
	iterator.inputIterator.Reset()
	iterator.fmIterator = nil
}

//-------combined2Iterator----------

type combined2Iterator[A, B, C any] struct {
	inputIterator1 Iterator[A]
	inputIterator2 Iterator[B]

	combineF func(A, B) C
}

func (iterator *combined2Iterator[A, B, C]) Next() Option[C] {
	return ApplyOption2(iterator.inputIterator1.Next(), iterator.inputIterator2.Next(), func(a A, b B) Option[C] {
		return Some(iterator.combineF(a, b))
	})
}

func (iterator *combined2Iterator[A, B, C]) Copy() Iterator[C] {
	return &combined2Iterator[A, B, C]{iterator.inputIterator1.Copy(), iterator.inputIterator2.Copy(), iterator.combineF}
}

func (iterator *combined2Iterator[A, B, C]) Reset() {
	iterator.inputIterator1.Reset()
	iterator.inputIterator2.Reset()
}

//-------LazySeq--------------

// Lazy Sequence evaluates elements only when they are needed (unlike the regular Sequence that does it eagerly).
// It has the same functions as Sequence but many of them are "Lazy" and would not cause any processing until that is needed.
type LazySeq[A any] struct {
	Iterator      Iterator[A]
	KnownCapacity int
	NilUnderlying bool
}

// Returns true if this Lazy Sequence contains an element that is equal (as determined by ==) to elem, false otherwise.
// [Materializing action: iterates over the underlying Sequence]
func ContainsInLazySeq[A comparable](lazySeq LazySeq[A], elem A) bool {
	return lazySeq.Exists(func(a A) bool { return a == elem })
}

// Copies this Lazy Sequence and returns a copy. The copy would share the same underlying Sequence but a different iterator.
// [Lazy action: doesn't iterate over the underlying Sequence, just returns a new Lazy Sequence]
func (lazySeq LazySeq[A]) Copy() LazySeq[A] {
	return LazySeq[A]{lazySeq.Iterator.Copy(), lazySeq.KnownCapacity, lazySeq.NilUnderlying}
}

// Returns false if this Lazy Sequence is empty or nil, otherwise true if the given predicate p holds for some of the elements of this Lazy Sequence, otherwise false.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) Exists(p func(A) bool) bool {
	return resetAndReturn(lazySeq, lazySeq.Find(p).IsDefined())
}

// Returns a new Lazy Sequence consisting of all elements of this Lazy Sequence that satisfy the given predicate p. The order of the elements is preserved.
// [Lazy action: doesn't iterate over the underlying Sequence, just returns a new Lazy Sequence]
func (lazySeq LazySeq[A]) Filter(p func(A) bool) LazySeq[A] {
	newIterator := filterIterator[A]{lazySeq.Iterator.Copy(), p}
	return LazySeq[A]{&newIterator, lazySeq.KnownCapacity, lazySeq.NilUnderlying}
}

// Returns a new Lazy Sequence consisting of all elements of this Lazy Sequence that do not satisfy the given predicate p. The order of the elements is preserved.
// [Lazy action: doesn't iterate over the underlying Sequence, just returns a new Lazy Sequence]
func (lazySeq LazySeq[A]) FilterNot(p func(A) bool) LazySeq[A] {
	return lazySeq.Filter(func(a A) bool { return !p(a) })
}

// Finds the first element of the Lazy Sequence satisfying a predicate, if any. Returns an option value containing the first element in the Lazy Sequence that satisfies p, or None if none exists.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) Find(p func(A) bool) Option[A] {
	for next := lazySeq.Iterator.Next(); next.IsDefined(); next = lazySeq.Iterator.Next() {
		if next.Exists(p) {
			return resetAndReturn(lazySeq, next)
		}
	}
	return resetAndReturn(lazySeq, None[A]())
}

// Returns a new Lazy Sequence resulting from applying the given function (f: A => LazySeq[A]) to each element of this Lazy Sequence and then flattening results back to LazySeq[A]. The original Lazy Sequence type A doesn't change.
// [Lazy action: doesn't iterate over the underlying Sequence, just returns a new Lazy Sequence]
func (lazySeq LazySeq[A]) FlatMap(f func(A) LazySeq[A]) LazySeq[A] {
	return FlatMapLazySeq(lazySeq, f)
}

// Returns a new Lazy Sequence resulting from applying the given function (f: A => LazySeq[B]) to each element of this Lazy Sequence and then flattening results to LazySeq[B]. The original Lazy Sequence type A could change to B.
// [Lazy action: doesn't iterate over the underlying Sequence, just returns a new Lazy Sequence]
func FlatMapLazySeq[A, B any](lazySeq LazySeq[A], f func(A) LazySeq[B]) LazySeq[B] {
	fI := func(a A) Iterator[B] {
		return f(a).Iterator
	}
	newIterator := flatMapIterator[A, B]{lazySeq.Iterator.Copy(), fI, nil}
	return LazySeq[B]{&newIterator, lazySeq.KnownCapacity, lazySeq.NilUnderlying}
}

// Applies a binary operator op to a start value z (of type A) and all Lazy Sequence elements (of type A), going left to right. The accumulation result also keeps the same type A.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) Fold(z A, op func(A, A) A) A {
	return resetAndReturn(lazySeq, FoldLazySeq(lazySeq, z, op))
}

// Applies a binary operator op to a start value z (of type B) and all Lazy Sequence elements (of type A), going left to right. The accumulation result is of type B.
// [Materializing action: iterates over the underlying Sequence]
func FoldLazySeq[A, B any](lazySeq LazySeq[A], z B, op func(B, A) B) B {
	r := z
	for next := lazySeq.Iterator.Next(); next.IsDefined(); next = lazySeq.Iterator.Next() {
		r = op(r, next.Get())
	}
	return resetAndReturn(lazySeq, r)
}

// Returns true if this Lazy Sequence is empty or nil or the given predicate p holds for all elements of this Lazy Sequence, otherwise false.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) ForAll(p func(A) bool) bool {
	return !lazySeq.Exists(func(a A) bool { return !p(a) })
}

// Applies a given procedure f to all elements of this Lazy Sequence.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) Foreach(f func(A)) {
	for next := lazySeq.Iterator.Next(); next.IsDefined(); next = lazySeq.Iterator.Next() {
		f(next.Get())
	}
	lazySeq.Iterator.Reset()
}

// Returns the first element of this Lazy Sequence. Returns type A's default value if the Lazy Sequence is empty or nil.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) Head() A {
	return resetAndReturn(lazySeq, lazySeq.Iterator.Next().GetOrElse(*new(A)))
}

// Returns the first element of this Lazy Sequence if it is nonempty, None if it is empty or nil.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) HeadOption() Option[A] {
	return resetAndReturn(lazySeq, lazySeq.Iterator.Next())
}

// Returns true if the Lazy Sequence contain no elements, false otherwise.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) IsEmpty() bool {
	return resetAndReturn(lazySeq, lazySeq.Iterator.Next().IsEmpty())
}

// Returns a new Lazy Sequence based on a provided underlying Sequence.
func LazySeqFromSeq[A any](seq Seq[A]) LazySeq[A] {
	return LazySeq[A]{&seqIterator[A]{seq, 0}, cap(seq), seq == nil}
}

// Returns the length of the Lazy Sequence.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) Length() int {
	count := 0
	for next := lazySeq.Iterator.Next(); next.IsDefined(); next = lazySeq.Iterator.Next() {
		count = count + 1
	}
	return resetAndReturn(lazySeq, count)
}

// Returns a new Lazy Sequence resulting from applying the given function f to each element of this Lazy Sequence and collecting the results (without changing type A of the Lazy Sequence elements).
// [Lazy action: doesn't iterate over the underlying Sequence, just returns a new Lazy Sequence]
func (lazySeq LazySeq[A]) Map(f func(A) A) LazySeq[A] {
	return MapLazySeq(lazySeq, f)
}

// Returns a new Lazy Sequence resulting from applying the given function f to each element of this Lazy Sequence and collecting the results (potentially, changing type A of the Sequence elements to B).
// [Lazy action: doesn't iterate over the underlying Sequence, just returns a new Lazy Sequence]
func MapLazySeq[A, B any](lazySeq LazySeq[A], f func(A) B) LazySeq[B] {
	newIterator := mapIterator[A, B]{lazySeq.Iterator.Copy(), f}
	return LazySeq[B]{&newIterator, lazySeq.KnownCapacity, lazySeq.NilUnderlying}
}

// Returns the largest element of this Lazy Sequence. Or a default value of type A if the Lazy Sequence is empty or nil.
// [Materializing action: iterates over the underlying Sequence]
func MaxInLazySeq[A Ordered](lazySeq LazySeq[A]) A {
	return resetAndReturn(lazySeq, lazySeq.Reduce(func(a1, a2 A) A {
		if a1 > a2 {
			return a1
		} else {
			return a2
		}
	}))
}

// Returns the smallest element of this Lazy Sequence. Or a default value of type A if the Lazy Sequence is empty or nil.
// [Materializing action: iterates over the underlying Sequence]
func MinInLazySeq[A Ordered](lazySeq LazySeq[A]) A {
	return resetAndReturn(lazySeq, lazySeq.Reduce(func(a1, a2 A) A {
		if a1 < a2 {
			return a1
		} else {
			return a2
		}
	}))
}

// Returns true if the Lazy Sequence contains at least one element, false otherwise.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) NonEmpty() bool {
	return !lazySeq.IsEmpty()
}

// Returns a result of applying reduce operator op between all the elements of the Lazy Sequence, going left to right. If the Lazy Sequence is empty or nil then a default value of type A is returned.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) Reduce(op func(A, A) A) A {
	next := lazySeq.Iterator.Next()
	if next.IsEmpty() {
		return *new(A)
	}

	r := next.Get()
	for next = lazySeq.Iterator.Next(); next.IsDefined(); next = lazySeq.Iterator.Next() {
		r = op(r, next.Get())
	}
	return resetAndReturn(lazySeq, r)
}

func resetAndReturn[A, B any](lazySeq LazySeq[A], result B) B {
	lazySeq.Iterator.Reset()
	return result
}

// Converts this Lazy Sequence into a materialized Sequence.
// [Materializing action: iterates over the underlying Sequence]
func (lazySeq LazySeq[A]) Strict() Seq[A] {
	if lazySeq.NilUnderlying {
		return nil
	}

	result := make(Seq[A], 0, lazySeq.KnownCapacity)
	for next := lazySeq.Iterator.Next(); next.IsDefined(); next = lazySeq.Iterator.Next() {
		result = result.Append(next.Get())
	}
	return resetAndReturn(lazySeq, result)
}

// Returns a new Lazy Sequence containing first n elements of this Lazy Sequence. Or the whole Lazy Sequence, if it has less than n elements. If n is negative, returns an empty Lazy Sequence.
// [Lazy action: doesn't iterate over the underlying Sequence, just returns a new Lazy Sequence]
func (lazySeq LazySeq[A]) Take(n int) LazySeq[A] {
	newIterator := filterIndexIterator[A]{lazySeq.Iterator.Copy(), func(i int) bool { return i < n }, 0}
	return LazySeq[A]{&newIterator, lazySeq.KnownCapacity, lazySeq.NilUnderlying}
}

// Converts this Lazy Sequence of Tuples into a Tuple of two Lazy Sequences.
// [Lazy action: doesn't iterate over the underlying Sequence, just returns new Lazy Sequences]
func UnZipLazySeq[A, B any](lazySeq LazySeq[Tuple2[A, B]]) Tuple2[LazySeq[A], LazySeq[B]] {
	return Tuple2[LazySeq[A], LazySeq[B]]{
		MapLazySeq(lazySeq.Copy(), func(t Tuple2[A, B]) A { return t.a }),
		MapLazySeq(lazySeq.Copy(), func(t Tuple2[A, B]) B { return t.b })}
}

// Returns a new Lazy Sequence formed from this Lazy Sequence and another Lazy Sequence by combining corresponding elements in Tuples. If one of the two Lazy Sequences is longer than the other, its remaining elements are ignored.
// [Lazy action: doesn't iterate over underlying Sequences, just returns a new combined Lazy Sequence]
func ZipLazySeq[A, B any](lazySeq LazySeq[A], another LazySeq[B]) LazySeq[Tuple2[A, B]] {
	newIterator := combined2Iterator[A, B, Tuple2[A, B]]{lazySeq.Iterator.Copy(), another.Iterator.Copy(), Tup2[A, B]}
	return LazySeq[Tuple2[A, B]]{&newIterator, lazySeq.KnownCapacity, lazySeq.NilUnderlying || another.NilUnderlying}
}
