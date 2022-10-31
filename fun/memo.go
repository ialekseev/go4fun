package fun

func Memo1[A comparable, B any](f func(A) B) func(A) B {
	m := make(map[A]B)

	return func(a A) B {
		result, ok := m[a]
		if ok {
			return result
		} else {
			evaluated := f(a)
			m[a] = evaluated
			return evaluated
		}
	}
}

func Memo2[A, B comparable, C any](f func(A, B) C) func(A, B) C {
	return UnTupled2(Memo1(Tupled2(f)))
}

func Memo3[A, B, C comparable, D any](f func(A, B, C) D) func(A, B, C) D {
	return UnTupled3(Memo1(Tupled3(f)))
}
