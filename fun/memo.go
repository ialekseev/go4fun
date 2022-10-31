package fun

// Wraps a provided function (with 1 argument) into a memoized function of the same signature.
// It would cache results of function calls and return back a cached result for the same input, if requested again.
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

// Wraps a provided function (with 2 arguments) into a memoized function of the same signature.
// It would cache results of function calls and return back a cached result for the same inputs, if requested again.
func Memo2[A, B comparable, C any](f func(A, B) C) func(A, B) C {
	return UnTupled2(Memo1(Tupled2(f)))
}

// Wraps a provided function (with 3 arguments) into a memoized function of the same signature.
// It would cache results of function calls and return back a cached result for the same inputs, if requested again.
func Memo3[A, B, C comparable, D any](f func(A, B, C) D) func(A, B, C) D {
	return UnTupled3(Memo1(Tupled3(f)))
}
