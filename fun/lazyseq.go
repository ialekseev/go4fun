package fun

type LazySeq[A any] struct {
	internal     Seq[A]
	currentIndex *int
	next         func(seq Seq[A], currentIndex *int) Option[A]
}

func (lazySeq LazySeq[A]) Filter(f func(A) bool) LazySeq[A] {
	nextF := lazySeq.next
	lazySeq.next = func(s Seq[A], currentIndex *int) Option[A] {
		for {
			current := nextF(s, currentIndex)
			switch {
			case current.IsDefined() && f(current.Get()):
				*currentIndex = *currentIndex + 1
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
	nextF := func(s Seq[A], currentIndex *int) Option[A] {
		if *currentIndex < seq.Length() {
			current := seq[*currentIndex]
			*currentIndex = *currentIndex + 1
			return Some(current)
		} else {
			return None[A]()
		}
	}
	return LazySeq[A]{seq, new(int), nextF}
}

func (lazySeq LazySeq[A]) Map(f func(A) A) LazySeq[A] {
	nextF := lazySeq.next
	lazySeq.next = func(s Seq[A], currentIndex *int) Option[A] {
		return nextF(s, currentIndex).Map(f)
	}
	return lazySeq
}

func (lazySeq LazySeq[A]) Next() Option[A] {
	return lazySeq.next(lazySeq.internal, lazySeq.currentIndex)
}

func (lazySeq LazySeq[A]) Strict() Seq[A] {
	if lazySeq.internal == nil {
		return nil
	}

	result := make(Seq[A], 0, lazySeq.internal.Length())

	for {
		if next := lazySeq.Next(); next.IsDefined() {
			result = result.Append(next.Get())
		} else {
			break
		}
	}
	return result
}
