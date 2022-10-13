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

type mapIterator[A any] struct {
	inputIterator Iterator[A]
	mapF          func(A) A
}

func (iterator mapIterator[A]) next() Option[A] {
	return iterator.inputIterator.next().Map(iterator.mapF)
}

//-------flatMapIterator----------

type flatMapIterator[A any] struct {
	inputIterator Iterator[A]
	flatMapF      func(A) Iterator[A]
	fmIterator    *Iterator[A]
}

func (iterator flatMapIterator[A]) setNewFmIteratorAndMove() Option[A] {
	next := iterator.inputIterator.next()
	if next.IsDefined() {
		*iterator.fmIterator = iterator.flatMapF(next.Get())
		return (*iterator.fmIterator).next()
	}
	return None[A]()
}

func (iterator flatMapIterator[A]) next() Option[A] {
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
	fI := func(a A) Iterator[A] {
		return f(a).iterator
	}
	lazySeq.iterator = flatMapIterator[A]{lazySeq.iterator, fI, new(Iterator[A])}
	return lazySeq
}

func (lazySeq LazySeq[A]) FilterNot(f func(A) bool) LazySeq[A] {
	return lazySeq.Filter(func(a A) bool { return !f(a) })
}

func LazySeqFromSeq[A any](seq Seq[A]) LazySeq[A] {
	return LazySeq[A]{seqIterator[A]{seq, new(int)}, cap(seq), seq == nil}
}

func (lazySeq LazySeq[A]) Map(f func(A) A) LazySeq[A] {
	lazySeq.iterator = mapIterator[A]{lazySeq.iterator, f}
	return lazySeq
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
