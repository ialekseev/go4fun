package fun

//-------iterator------------

type iterator[A any] interface {
	hasMore() bool
	next() A
}

//-------seqIterator----------

type seqIterator[A any] struct {
	seq          Seq[A]
	currentIndex *int
}

func (iterator seqIterator[A]) hasMore() bool {
	return *iterator.currentIndex < iterator.seq.Length()
}

func (iterator seqIterator[A]) next() A {
	current := iterator.seq[*iterator.currentIndex]
	*iterator.currentIndex = *iterator.currentIndex + 1
	return current
}

//-------LazySeq--------------

type LazySeq[A any] struct {
	iterator      iterator[A]
	next          func() Option[A]
	knownCapacity int
}

func (lazySeq LazySeq[A]) Filter(f func(A) bool) LazySeq[A] {
	nextF := lazySeq.next
	lazySeq.next = func() Option[A] {
		for {
			current := nextF()
			switch {
			case current.IsDefined() && f(current.Get()):
				return current
			case current.IsDefined() && !f(current.Get()):
				continue
			default:
				return None[A]()
			}
		}
	}
	return lazySeq
}

func (lazySeq LazySeq[A]) FlatMap(f func(A) LazySeq[A]) LazySeq[A] {
	panic("Not Implemented")
}

func (lazySeq LazySeq[A]) FilterNot(f func(A) bool) LazySeq[A] {
	return lazySeq.Filter(func(a A) bool { return !f(a) })
}

func LazySeqFromSeq[A any](seq Seq[A]) LazySeq[A] {
	lazySeq := LazySeq[A]{seqIterator[A]{seq, new(int)}, nil, cap(seq)}

	lazySeq.next = func() Option[A] {
		if lazySeq.iterator.hasMore() {
			return Some(lazySeq.iterator.next())
		} else {
			return None[A]()
		}
	}
	return lazySeq
}

func (lazySeq LazySeq[A]) Map(f func(A) A) LazySeq[A] {
	nextF := lazySeq.next
	lazySeq.next = func() Option[A] {
		return nextF().Map(f)
	}
	return lazySeq
}

func (lazySeq LazySeq[A]) Next() Option[A] {
	return lazySeq.next()
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
