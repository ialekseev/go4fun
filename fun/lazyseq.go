package fun

//-------iterator------------

type iterator[A any] interface {
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
	inputIterator iterator[A]
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
	inputIterator iterator[A]
	mapF          func(A) A
}

func (iterator mapIterator[A]) next() Option[A] {
	return iterator.inputIterator.next().Map(iterator.mapF)
}

//-------LazySeq--------------

type LazySeq[A any] struct {
	iterator      iterator[A]
	knownCapacity int
}

func (lazySeq LazySeq[A]) Filter(f func(A) bool) LazySeq[A] {
	lazySeq.iterator = filterIterator[A]{lazySeq.iterator, f}
	return lazySeq
}

func (lazySeq LazySeq[A]) FlatMap(f func(A) LazySeq[A]) LazySeq[A] {
	panic("Not Implemented")
}

func (lazySeq LazySeq[A]) FilterNot(f func(A) bool) LazySeq[A] {
	return lazySeq.Filter(func(a A) bool { return !f(a) })
}

func LazySeqFromSeq[A any](seq Seq[A]) LazySeq[A] {
	return LazySeq[A]{seqIterator[A]{seq, new(int)}, cap(seq)}
}

func (lazySeq LazySeq[A]) Map(f func(A) A) LazySeq[A] {
	lazySeq.iterator = mapIterator[A]{lazySeq.iterator, f}
	return lazySeq
}

func (lazySeq LazySeq[A]) Next() Option[A] {
	return lazySeq.iterator.next()
}

func (lazySeq LazySeq[A]) Strict() Seq[A] {
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
