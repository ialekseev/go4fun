package fun

// Partially applies a 2-argument function `f` to its 1st argument (A). Returns a function of 1 remaining argument (B).
func Apply2Partial_1[A, B, C any](f func(A, B) C, a A) func(B) C {
	return func(b B) C {
		return f(a, b)
	}
}

// Partially applies a 2-argument function `f` to its 2nd argument (B). Returns a function of 1 remaining argument (A).
func Apply2Partial_2[A, B, C any](f func(A, B) C, b B) func(A) C {
	return func(a A) C {
		return f(a, b)
	}
}

// Partially applies a 3-argument function `f` to its 1st argument (A). Returns a function of 2 remaining arguments (B, C).
func Apply3Partial_1[A, B, C, D any](f func(A, B, C) D, a A) func(B, C) D {
	return func(b B, c C) D {
		return f(a, b, c)
	}
}

// Partially applies a 3-argument function `f` to its 2nd argument (B). Returns a function of 2 remaining arguments (A, C).
func Apply3Partial_2[A, B, C, D any](f func(A, B, C) D, b B) func(A, C) D {
	return func(a A, c C) D {
		return f(a, b, c)
	}
}

// Partially applies a 3-argument function `f` to its 3rd argument (C). Returns a function of 2 remaining arguments (A, B).
func Apply3Partial_3[A, B, C, D any](f func(A, B, C) D, c C) func(A, B) D {
	return func(a A, b B) D {
		return f(a, b, c)
	}
}

// Partially applies a 3-argument function `f` to its 1st & 2nd arguments (A, B). Returns a function of 1 remaining argument (C).
func Apply3Partial_1_2[A, B, C, D any](f func(A, B, C) D, a A, b B) func(C) D {
	return func(c C) D {
		return f(a, b, c)
	}
}

// Partially applies a 3-argument function `f` to its 1st & 3rd arguments (A, C). Returns a function of 1 remaining argument (B).
func Apply3Partial_1_3[A, B, C, D any](f func(A, B, C) D, a A, c C) func(B) D {
	return func(b B) D {
		return f(a, b, c)
	}
}

// Partially applies a 3-argument function `f` to its 2nd & 3rd arguments (B, C). Returns a function of 1 remaining argument (A).
func Apply3Partial_2_3[A, B, C, D any](f func(A, B, C) D, b B, c C) func(A) D {
	return func(a A) D {
		return f(a, b, c)
	}
}

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
