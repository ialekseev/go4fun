package fun

import (
	"fmt"
)

type Either[L, R any] Tuple2[Option[L], Option[R]]

// Returns a result of applying f to Right value if this Either is Right (without changing Right type R). Different from map in that f is expected to return an Either (which could be Left).
// Returns unchanged Left if this Either is Left.
func (either Either[L, R]) FlatMap(f func(R) Either[L, R]) Either[L, R] {
	return FlatMapEither(either, f)
}

// Returns a result of applying f to Right value if this Either is Right (potentially, changing Right type R => T)). Different from map in that f is expected to return an Either (which could be Left).
// Returns unchanged Left if this Either is Left.
func FlatMapEither[L, R, T any](either Either[L, R], f func(R) Either[L, T]) Either[L, T] {
	if either.IsRight() {
		return f(either.RightOption().Get())
	} else {
		return Left[L, T](either.LeftOption().Get())
	}
}

// Returns true if this Either is Left.
func (either Either[L, R]) IsLeft() bool {
	return either.a.IsDefined() && either.b.IsEmpty()
}

// Returns true if this Either is Right.
func (either Either[L, R]) IsRight() bool {
	return either.a.IsEmpty() && either.b.IsDefined()
}

// Creates a new Left Either.
func Left[L, R any](l L) Either[L, R] {
	return Either[L, R]{Some(l), None[R]()}
}

// Returns a Left Option from Either.
func (either Either[L, R]) LeftOption() Option[L] {
	return either.a
}

// Returns a result of applying f to Right value if this Either is Right (without changing Right type R).
// Returns unchanged Left if this Either is Left.
func (either Either[L, R]) Map(f func(R) R) Either[L, R] {
	return MapEither(either, f)
}

// Returns a result of applying f to Right value if this Either is Right (potentially, changing Right type R => T)).
// Returns unchanged Left if this Either is Left.
func MapEither[L, R, T any](either Either[L, R], f func(R) T) Either[L, T] {
	if either.IsRight() {
		return Right[L](f(either.RightOption().Get()))
	} else {
		return Left[L, T](either.LeftOption().Get())
	}
}

// Creates a new Right Either.
func Right[L, R any](r R) Either[L, R] {
	return Either[L, R]{None[L](), Some(r)}
}

// Returns a Right Option from Either.
func (either Either[L, R]) RightOption() Option[R] {
	return either.b
}

// A String representation of Either. E.g. Right(5) or Left(bad)
func (either Either[L, R]) String() string {
	if either.IsRight() {
		return fmt.Sprintf("Right(%v)", either.RightOption().Get())
	} else {
		return fmt.Sprintf("Left(%v)", either.LeftOption().Get())
	}
}

// Returns a Some containing the Right value if it exists or a None if this Either is Left.
func (either Either[L, R]) ToOption() Option[R] {
	return either.RightOption()
}
