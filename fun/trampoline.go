package fun

type Trampoline[A any] struct {
	call   func() Trampoline[A]
	done   bool
	result A
}

func DoneTrampolining[A any](a A) Trampoline[A] {
	return Trampoline[A]{done: true, result: a}
}

func MoreTrampolining[A any](more func() Trampoline[A]) Trampoline[A] {
	return Trampoline[A]{call: more, done: false}
}

func (t Trampoline[A]) Run() A {
	next := t
	for {
		if next.done {
			return next.result
		}
		next = next.call()
	}
}
