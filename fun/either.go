// Work In Progress
package fun

type Either[L, R comparable] struct {
	left    Option[L]
	right   Option[R]
	isRight bool
}

func Left[L, R comparable](l L) Either[L, R] {
	return Either[L, R]{Some(l), None[R](), false}
}

func Right[L, R comparable](r R) Either[L, R] {
	return Either[L, R]{None[L](), Some(r), true}
}

func (either Either[L, R]) IsLeft() bool {
	return !either.isRight
}

func (either Either[L, R]) IsRight() bool {
	return either.isRight
}

func (either Either[L, R]) ToOption() Option[R] {
	return either.right
}

func (either Either[L, R]) Right() Option[R] {
	return either.right
}

func (either Either[L, R]) Left() Option[L] {
	return either.left
}

func (either Either[L, R]) Map(f func(R) R) Either[L, R] {
	if either.isRight {
		return Right[L](MapOption(either.right, f).Get())
	} else {
		return either
	}
}
