package fun

//-------Iterator------------

type Iterator[A any] interface {
	hasMore() bool
	next() A
}

//-------SeqIterator----------

type SeqIterator[A any] struct {
	seq          Seq[A]
	currentIndex *int
}

func (iterator SeqIterator[A]) hasMore() bool {
	return *iterator.currentIndex < iterator.seq.Length()
}

func (iterator SeqIterator[A]) next() A {
	current := iterator.seq[*iterator.currentIndex]
	*iterator.currentIndex = *iterator.currentIndex + 1
	return current
}

//-------LazySeq--------------

type LazySeq[A any] struct {
	iterator      Iterator[A]
	next          func(Iterator[A]) Option[A]
	knownCapacity int
}

func (lazySeq LazySeq[A]) Filter(f func(A) bool) LazySeq[A] {
	nextF := lazySeq.next
	lazySeq.next = func(iterator Iterator[A]) Option[A] {
		for {
			current := nextF(iterator)
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

func (lazySeq LazySeq[A]) FilterNot(f func(A) bool) LazySeq[A] {
	return lazySeq.Filter(func(a A) bool { return !f(a) })
}

func LazySeqFromSeq[A any](seq Seq[A]) LazySeq[A] {
	lazySeq := LazySeq[A]{SeqIterator[A]{seq, new(int)}, nil, cap(seq)}

	lazySeq.next = func(iterator Iterator[A]) Option[A] {
		if iterator.hasMore() {
			return Some(iterator.next())
		} else {
			return None[A]()
		}
	}
	return lazySeq
}

func (lazySeq LazySeq[A]) Map(f func(A) A) LazySeq[A] {
	nextF := lazySeq.next
	lazySeq.next = func(iterator Iterator[A]) Option[A] {
		return nextF(iterator).Map(f)
	}
	return lazySeq
}

func (lazySeq LazySeq[A]) Next() Option[A] {
	return lazySeq.next(lazySeq.iterator)
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
