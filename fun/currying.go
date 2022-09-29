package fun

func Curry2[A, B, C any](f func(A, B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return f(a, b)
		}
	}
}

func Curry3[A, B, C, D any](f func(A, B, C) D) func(A) func(B) func(C) D {
	return func(a A) func(B) func(C) D {
		return func(b B) func(C) D {
			return func(c C) D {
				return f(a, b, c)
			}
		}
	}
}

func UnCurry2[A, B, C any](f func(A) func(B) C) func(A, B) C {
	return func(a A, b B) C {
		return f(a)(b)
	}
}

func UnCurry3[A, B, C, D any](f func(A) func(B) func(C) D) func(A, B, C) D {
	return func(a A, b B, c C) D {
		return f(a)(b)(c)
	}
}
