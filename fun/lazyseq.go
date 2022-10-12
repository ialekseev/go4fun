package fun

type LazySeq[A any] struct {
	currentIndex              *int
	next                      func() Option[A]
	underlyingSeqCapacityHint int
}

func (lazySeq LazySeq[A]) Filter(f func(A) bool) LazySeq[A] {
	nextF := lazySeq.next
	lazySeq.next = func() Option[A] {
		for {
			current := nextF()
			switch {
			case current.IsDefined() && f(current.Get()):
				*lazySeq.currentIndex = *lazySeq.currentIndex + 1
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

// func (lazySeq LazySeq[A]) FlatMap(f func(A) LazySeq[A]) LazySeq[A] {
// 	nextF := lazySeq.next
// 	lazySeq.next = func() Option[A] {

// 		if n:=nextF(); n.IsDefined() {
// 			flatMapped := f(n.Get())

// 			flatMapped.n
// 		}

// 	}
// 	return lazySeq
// }

func LazySeqFromSeq[A any](seq Seq[A]) LazySeq[A] {
	lazySeq := LazySeq[A]{new(int), nil, cap(seq)}

	lazySeq.next = func() Option[A] {
		if *lazySeq.currentIndex < seq.Length() {
			current := seq[*lazySeq.currentIndex]
			*lazySeq.currentIndex = *lazySeq.currentIndex + 1
			return Some(current)
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
	result := make(Seq[A], 0, lazySeq.underlyingSeqCapacityHint)

	for {
		if next := lazySeq.Next(); next.IsDefined() {
			result = result.Append(next.Get())
		} else {
			break
		}
	}
	return result
}
