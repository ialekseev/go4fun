package fun

//-------Iterator------------

type Iterator[A any] interface {
	next() Option[A]
}

//-------seqIterator----------

type seqIterator[A any] struct {
	seq          Seq[A]
	currentIndex *int
}

func (iterator seqIterator[A]) next() Option[A] {
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

func (iterator filterIterator[A]) next() Option[A] {
	for {
		next := iterator.inputIterator.next()
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

func (iterator mapIterator[A, B]) next() Option[B] {
	return MapOption(iterator.inputIterator.next(), iterator.mapF)
}

//-------flatMapIterator----------

type flatMapIterator[A, B any] struct {
	inputIterator Iterator[A]
	flatMapF      func(A) Iterator[B]
	fmIterator    *Iterator[B]
}

func (iterator flatMapIterator[A, B]) setNewFmIteratorAndMove() Option[B] {
	next := iterator.inputIterator.next()
	if next.IsDefined() {
		*iterator.fmIterator = iterator.flatMapF(next.Get())
		return (*iterator.fmIterator).next()
	}
	return None[B]()
}

func (iterator flatMapIterator[A, B]) next() Option[B] {
	if *iterator.fmIterator == nil {
		return iterator.setNewFmIteratorAndMove()
	} else {
		nextFm := (*iterator.fmIterator).next()
		if nextFm.IsEmpty() {
			return iterator.setNewFmIteratorAndMove()
		} else {
			return nextFm
		}
	}
}

//-------LazySeq--------------

type LazySeq[A any] struct {
	iterator      Iterator[A]
	knownCapacity int
	nilUnderlying bool
}

func (lazySeq LazySeq[A]) Filter(f func(A) bool) LazySeq[A] {
	lazySeq.iterator = filterIterator[A]{lazySeq.iterator, f}
	return lazySeq
}

func (lazySeq LazySeq[A]) FlatMap(f func(A) LazySeq[A]) LazySeq[A] {
	return FlatMapLazySeq(lazySeq, f)
}

func FlatMapLazySeq[A, B any](lazySeq LazySeq[A], f func(A) LazySeq[B]) LazySeq[B] {
	fI := func(a A) Iterator[B] {
		return f(a).iterator
	}
	newIterator := flatMapIterator[A, B]{lazySeq.iterator, fI, new(Iterator[B])}
	return LazySeq[B]{newIterator, lazySeq.knownCapacity, lazySeq.nilUnderlying}
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
	newIterator := mapIterator[A, B]{lazySeq.iterator, f}
	return LazySeq[B]{newIterator, lazySeq.knownCapacity, lazySeq.nilUnderlying}
}

func (lazySeq LazySeq[A]) Next() Option[A] {
	return lazySeq.iterator.next()
}

func (lazySeq LazySeq[A]) Strict() Seq[A] {
	if lazySeq.nilUnderlying {
		return nil
	}

	result := make(Seq[A], 0, lazySeq.knownCapacity)

	for {
		if next := lazySeq.Next(); next.IsDefined() {
			result = result.Append(next.Get())
		} else {
			break
		}
	}
	return result
}
