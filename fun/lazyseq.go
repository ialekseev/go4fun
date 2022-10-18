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
	return &seqIterator[A]{iterator.seq, iterator.currentIndex}
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
	return &flatMapIterator[A, B]{iterator.inputIterator.Copy(), iterator.flatMapF, iterator.fmIterator}
}

func (iterator *flatMapIterator[A, B]) Reset() {
	iterator.inputIterator.Reset()
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

type LazySeq[A any] struct {
	Iterator      Iterator[A]
	KnownCapacity int
	NilUnderlying bool
}

func (lazySeq LazySeq[A]) Copy() LazySeq[A] {
	return LazySeq[A]{lazySeq.Iterator.Copy(), lazySeq.KnownCapacity, lazySeq.NilUnderlying}
}

func (lazySeq LazySeq[A]) Exists(p func(A) bool) bool {
	return resetAndReturn(lazySeq, lazySeq.Find(p).IsDefined())
}

func (lazySeq LazySeq[A]) Filter(p func(A) bool) LazySeq[A] {
	newIterator := filterIterator[A]{lazySeq.Iterator.Copy(), p}
	return LazySeq[A]{&newIterator, lazySeq.KnownCapacity, lazySeq.NilUnderlying}
}

func (lazySeq LazySeq[A]) FilterNot(p func(A) bool) LazySeq[A] {
	return lazySeq.Filter(func(a A) bool { return !p(a) })
}

func (lazySeq LazySeq[A]) Find(p func(A) bool) Option[A] {
	for next := lazySeq.Iterator.Next(); next.IsDefined(); next = lazySeq.Iterator.Next() {
		if next.Exists(p) {
			return resetAndReturn(lazySeq, next)
		}
	}
	return resetAndReturn(lazySeq, None[A]())
}

func (lazySeq LazySeq[A]) FlatMap(f func(A) LazySeq[A]) LazySeq[A] {
	return FlatMapLazySeq(lazySeq, f)
}

func FlatMapLazySeq[A, B any](lazySeq LazySeq[A], f func(A) LazySeq[B]) LazySeq[B] {
	fI := func(a A) Iterator[B] {
		return f(a).Iterator
	}
	newIterator := flatMapIterator[A, B]{lazySeq.Iterator.Copy(), fI, nil}
	return LazySeq[B]{&newIterator, lazySeq.KnownCapacity, lazySeq.NilUnderlying}
}

func (lazySeq LazySeq[A]) Fold(z A, op func(A, A) A) A {
	return resetAndReturn(lazySeq, FoldLazySeq(lazySeq, z, op))
}

func FoldLazySeq[A, B any](lazySeq LazySeq[A], z B, op func(B, A) B) B {
	r := z
	for next := lazySeq.Iterator.Next(); next.IsDefined(); next = lazySeq.Iterator.Next() {
		r = op(r, next.Get())
	}
	return resetAndReturn(lazySeq, r)
}

func (lazySeq LazySeq[A]) ForAll(p func(A) bool) bool {
	return !lazySeq.Exists(func(a A) bool { return !p(a) })
}

func (lazySeq LazySeq[A]) Foreach(f func(A)) {
	for next := lazySeq.Iterator.Next(); next.IsDefined(); next = lazySeq.Iterator.Next() {
		f(next.Get())
	}
	lazySeq.Iterator.Reset()
}

func (lazySeq LazySeq[A]) Head() A {
	return resetAndReturn(lazySeq, lazySeq.Iterator.Next().GetOrElse(*new(A)))
}

func (lazySeq LazySeq[A]) HeadOption() Option[A] {
	return resetAndReturn(lazySeq, lazySeq.Iterator.Next())
}

func (lazySeq LazySeq[A]) IsEmpty() bool {
	return resetAndReturn(lazySeq, lazySeq.Iterator.Next().IsEmpty())
}

func LazySeqFromSeq[A any](seq Seq[A]) LazySeq[A] {
	return LazySeq[A]{&seqIterator[A]{seq, 0}, cap(seq), seq == nil}
}

func (lazySeq LazySeq[A]) Length() int {
	count := 0
	for next := lazySeq.Iterator.Next(); next.IsDefined(); next = lazySeq.Iterator.Next() {
		count = count + 1
	}
	return resetAndReturn(lazySeq, count)
}

func (lazySeq LazySeq[A]) Map(f func(A) A) LazySeq[A] {
	return MapLazySeq(lazySeq, f)
}

func MapLazySeq[A, B any](lazySeq LazySeq[A], f func(A) B) LazySeq[B] {
	newIterator := mapIterator[A, B]{lazySeq.Iterator.Copy(), f}
	return LazySeq[B]{&newIterator, lazySeq.KnownCapacity, lazySeq.NilUnderlying}
}

func MaxInLazySeq[A Ordered](lazySeq LazySeq[A]) A {
	return resetAndReturn(lazySeq, lazySeq.Reduce(func(a1, a2 A) A {
		if a1 > a2 {
			return a1
		} else {
			return a2
		}
	}))
}

func MinInLazySeq[A Ordered](lazySeq LazySeq[A]) A {
	return resetAndReturn(lazySeq, lazySeq.Reduce(func(a1, a2 A) A {
		if a1 < a2 {
			return a1
		} else {
			return a2
		}
	}))
}

func (lazySeq LazySeq[A]) NonEmpty() bool {
	return !lazySeq.IsEmpty()
}

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

func UnZipLazySeq[A, B any](lazySeq LazySeq[Tuple2[A, B]]) Tuple2[LazySeq[A], LazySeq[B]] {
	return Tuple2[LazySeq[A], LazySeq[B]]{
		MapLazySeq(lazySeq.Copy(), func(t Tuple2[A, B]) A { return t.a }),
		MapLazySeq(lazySeq.Copy(), func(t Tuple2[A, B]) B { return t.b })}
}

func ZipLazySeq[A, B any](lazySeq LazySeq[A], another LazySeq[B]) LazySeq[Tuple2[A, B]] {
	newIterator := combined2Iterator[A, B, Tuple2[A, B]]{lazySeq.Iterator.Copy(), another.Iterator.Copy(), Tup2[A, B]}
	return LazySeq[Tuple2[A, B]]{&newIterator, lazySeq.KnownCapacity, lazySeq.NilUnderlying || another.NilUnderlying}
}
