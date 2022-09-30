package fun

// Takes 2 functions f(A)=>B and g(B)=>C and returns a composed function h(A)=>C.
func Compose2[A, B, C any](f func(A) B, g func(B) C) func(A) C {
	return func(a A) C { return g(f(a)) }
}

// Takes 3 functions f(A)=>B, g(B)=>C and h(C)=>D and returns a composed function j(A)=>D.
func Compose3[A, B, C, D any](f func(A) B, g func(B) C, h func(C) D) func(A) D {
	return func(a A) D { return h(g(f(a))) }
}
