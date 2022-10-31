package fun

// Takes 2 functions f(A)=>B and g(B)=>C and returns a composed function h(A)=>C.
func Compose2[A, B, C any](f func(A) B, g func(B) C) func(A) C {
	return func(a A) C { return g(f(a)) }
}

// Takes 3 functions f(A)=>B, g(B)=>C and h(C)=>D and returns a composed function j(A)=>D.
func Compose3[A, B, C, D any](f func(A) B, g func(B) C, h func(C) D) func(A) D {
	return func(a A) D { return h(g(f(a))) }
}

// Transforms a function of 2 arguments into a chain of 2 1-argument functions.
func Curry2[A, B, C any](f func(A, B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return f(a, b)
		}
	}
}

// Transforms a function of 3 arguments into a chain of 3 1-argument functions.
func Curry3[A, B, C, D any](f func(A, B, C) D) func(A) func(B) func(C) D {
	return func(a A) func(B) func(C) D {
		return func(b B) func(C) D {
			return func(c C) D {
				return f(a, b, c)
			}
		}
	}
}

// Transforms a chain of 2 1-argument functions into a function of 2 arguments.
func UnCurry2[A, B, C any](f func(A) func(B) C) func(A, B) C {
	return func(a A, b B) C {
		return f(a)(b)
	}
}

// Transforms a chain of 3 1-argument functions into a function of 3 arguments.
func UnCurry3[A, B, C, D any](f func(A) func(B) func(C) D) func(A, B, C) D {
	return func(a A, b B, c C) D {
		return f(a)(b)(c)
	}
}

// Transforms a function of 2 arguments into a function of 1 tupled argument.
func Tupled2[A, B, C any](f func(A, B) C) func(Tuple2[A, B]) C {
	return func(t Tuple2[A, B]) C {
		return f(t.a, t.b)
	}
}

// Transforms a function of 3 arguments into a function of 1 tupled argument.
func Tupled3[A, B, C, D any](f func(A, B, C) D) func(Tuple3[A, B, C]) D {
	return func(t Tuple3[A, B, C]) D {
		return f(t.a, t.b, t.c)
	}
}

// Transforms a function of 1 tupled argument into a function of 2 arguments.
func UnTupled2[A, B, C any](f func(Tuple2[A, B]) C) func(A, B) C {
	return func(a A, b B) C {
		return f(Tup2(a, b))
	}
}

// Transforms a function of 1 tupled argument into a function of 3 arguments.
func UnTupled3[A, B, C, D any](f func(Tuple3[A, B, C]) D) func(A, B, C) D {
	return func(a A, b B, c C) D {
		return f(Tup3(a, b, c))
	}
}
