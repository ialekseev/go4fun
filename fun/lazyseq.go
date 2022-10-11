package fun

type LazySeq[A any] struct {
	slice        []A
	currentIndex *int
	next         func(slice []A, currentIndex *int) Option[A]
}

func (seq LazySeq[A]) Filter(f func(A) bool) LazySeq[A] {
	seq.next = func(slice []A, currentIndex *int) Option[A] {
		for {
			current := seq.next(slice, currentIndex)
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
	return seq
}

func LazySeqFromSlice[A any](slice []A) LazySeq[A] {
	nextF := func(slice []A, currentIndex *int) Option[A] {
		return next(slice, currentIndex)
	}
	return LazySeq[A]{slice, new(int), nextF}
}

func (seq LazySeq[A]) Map(f func(A) A) LazySeq[A] {
	seq.next = func(slice []A, currentIndex *int) Option[A] {
		return seq.next(slice, currentIndex).Map(func(a A) A { return f(a) })
	}
	return seq
}

func (seq LazySeq[A]) Next() Option[A] {
	return next(seq.slice, seq.currentIndex)
}

func next[A any](slice []A, currentIndex *int) Option[A] {
	if *currentIndex < len(slice) {
		current := slice[*currentIndex]
		*currentIndex = *currentIndex + 1
		return Some(current)
	} else {
		return None[A]()
	}
}
