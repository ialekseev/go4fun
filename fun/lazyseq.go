package fun

//-------Iterator------------

type Iterator[A any] interface {
	Next() Option[A]
}

//-------seqIterator----------

type seqIterator[A any] struct {
	seq          Seq[A]
	currentIndex *int
}

func (iterator seqIterator[A]) Next() Option[A] {
	if *iterator.currentIndex < iterator.seq.Length() {
		current := iterator.seq[*iterator.currentIndex]
		*iterator.currentIndex = *iterator.currentIndex + 1
		return Some(current)
	} else {
		return None[A]()
	}
}

//-------filterIterator----------

type filterIterator[A any] struct {
	inputIterator Iterator[A]
	filterF       func(A) bool
}

func (iterator filterIterator[A]) Next() Option[A] {
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

//-------mapIterator----------

type mapIterator[A, B any] struct {
	inputIterator Iterator[A]
	mapF          func(A) B
}

func (iterator mapIterator[A, B]) Next() Option[B] {
	return MapOption(iterator.inputIterator.Next(), iterator.mapF)
}

//-------flatMapIterator----------

type flatMapIterator[A, B any] struct {
	inputIterator Iterator[A]
	flatMapF      func(A) Iterator[B]
	fmIterator    *Iterator[B]
}

func (iterator flatMapIterator[A, B]) setNewFmIteratorAndMove() Option[B] {
	next := iterator.inputIterator.Next()
	if next.IsDefined() {
		*iterator.fmIterator = iterator.flatMapF(next.Get())
		return (*iterator.fmIterator).Next()
	}
	return None[B]()
}

func (iterator flatMapIterator[A, B]) Next() Option[B] {
	if *iterator.fmIterator == nil {
		return iterator.setNewFmIteratorAndMove()
	} else {
		nextFm := (*iterator.fmIterator).Next()
		if nextFm.IsEmpty() {
			return iterator.setNewFmIteratorAndMove()
		} else {
			return nextFm
		}
	}
}

//-------LazySeq--------------

type LazySeq[A any] struct {
	Iterator      Iterator[A]
	KnownCapacity int
	NilUnderlying bool
}

func (lazySeq LazySeq[A]) Filter(f func(A) bool) LazySeq[A] {
	lazySeq.Iterator = filterIterator[A]{lazySeq.Iterator, f}
	return lazySeq
}

func (lazySeq LazySeq[A]) FlatMap(f func(A) LazySeq[A]) LazySeq[A] {
	return FlatMapLazySeq(lazySeq, f)
}

func FlatMapLazySeq[A, B any](lazySeq LazySeq[A], f func(A) LazySeq[B]) LazySeq[B] {
	fI := func(a A) Iterator[B] {
		return f(a).Iterator
	}
	newIterator := flatMapIterator[A, B]{lazySeq.Iterator, fI, new(Iterator[B])}
	return LazySeq[B]{newIterator, lazySeq.KnownCapacity, lazySeq.NilUnderlying}
}

func (lazySeq LazySeq[A]) FilterNot(f func(A) bool) LazySeq[A] {
	return lazySeq.Filter(func(a A) bool { return !f(a) })
}

func LazySeqFromSeq[A any](seq Seq[A]) LazySeq[A] {
	return LazySeq[A]{seqIterator[A]{seq, new(int)}, cap(seq), seq == nil}
}

func (lazySeq LazySeq[A]) Map(f func(A) A) LazySeq[A] {
	return MapLazySeq(lazySeq, f)
}

func MapLazySeq[A, B any](lazySeq LazySeq[A], f func(A) B) LazySeq[B] {
	newIterator := mapIterator[A, B]{lazySeq.Iterator, f}
	return LazySeq[B]{newIterator, lazySeq.KnownCapacity, lazySeq.NilUnderlying}
}

func (lazySeq LazySeq[A]) Next() Option[A] {
	return lazySeq.Iterator.Next()
}

func (lazySeq LazySeq[A]) Strict() Seq[A] {
	if lazySeq.NilUnderlying {
		return nil
	}

	result := make(Seq[A], 0, lazySeq.KnownCapacity)

	for {
		if next := lazySeq.Next(); next.IsDefined() {
			result = result.Append(next.Get())
		} else {
			break
		}
	}
	return result
}
