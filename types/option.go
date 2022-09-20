package types

import (
	"fmt"
)

type Option[A comparable] struct {
	value   A
	defined bool
}

// Wrap a value A in Some[A]. Some[A] represents an existing value of type A.
func Some[A comparable](value A) Option[A] {
	return Option[A]{value, true}
}

// None[A] represents a non-existing value of type A.
func None[A comparable]() Option[A] {
	return Option[A]{}
}

// Returns true if this Option has an element that is equal (as determined by ==) to elem, false otherwise.
func (option Option[A]) Contains(elem A) bool {
	return option.defined && option.value == elem
}

// Returns true if this Option is nonempty and the predicate p returns true when applied to this Option's value. Otherwise, returns false.
func (option Option[A]) Exists(p func(A) bool) bool {
	return option.defined && p(option.value)
}

// Returns this Option if it is nonempty and applying the predicate p to this Option's value returns true. Otherwise, return None.
func (option Option[A]) Filter(p func(A) bool) Option[A] {
	if option.defined && p(option.value) {
		return option
	} else {
		return None[A]()
	}
}

// Returns this Option if it is nonempty and applying the predicate p to this Option's value returns false. Otherwise, return None.
func (option Option[A]) FilterNot(p func(A) bool) Option[A] {
	if option.defined && !p(option.value) {
		return option
	} else {
		return None[A]()
	}
}

// Returns the result of applying f to this Option's value if this Option is nonempty (without changing Option value's type A).
// Returns None if this Option is empty. Slightly different from map in that f is expected to return an Option (which could be None).
func (option Option[A]) FlatMap(f func(A) Option[A]) Option[A] {
	if option.defined {
		return f(option.value)
	} else {
		return None[A]()
	}
}

// Returns the result of applying f to this Option's value if this Option is nonempty (potentially, changing Option value's type A => B).
// Returns None if this Option is empty. Slightly different from map in that f is expected to return an Option (which could be None).
func FlatMapOption[A, B comparable](option Option[A], f func(A) Option[B]) Option[B] {
	if option.defined {
		return f(option.value)
	} else {
		return None[B]()
	}
}

// Returns the nested Option value if this Option is nonempty.
func FlattenOption[A comparable](option Option[Option[A]]) Option[A] {
	if option.defined {
		return option.value
	} else {
		return None[A]()
	}
}

// Returns the result of applying f to this Option's value if the Option is nonempty. Otherwise, returns defaultValue.
// Resulting value's type A is the same as Option value's type A.
func (option Option[A]) Fold(defaultValue A, f func(A) A) A {
	if option.defined {
		return f(option.value)
	} else {
		return defaultValue
	}
}

// Returns the result of applying f to this Option's value if the Option is nonempty. Otherwise, returns defaultValue.
// Resulting value's type B could, potentially, be different from the Option value's type A.
func FoldOption[A, B comparable](option Option[A], defaultValue B, f func(A) B) B {
	if option.defined {
		return f(option.value)
	} else {
		return defaultValue
	}
}

// Returns true if this Option is empty or the predicate p returns true when applied to this Option's value.
func (option Option[A]) ForAll(p func(A) bool) bool {
	return !option.defined || p(option.value)
}

// Apply the given procedure f to the Option's value, if it is nonempty. Otherwise, do nothing.
func (option Option[A]) Foreach(f func(A)) {
	panic("Not implemented")
}

// Evaluate and return alternate value if empty
func (option Option[A]) GetOrElse(defaultValue A) Option[A] {
	panic("Not implemented")
}

// Return value, panic if empty
func (option Option[A]) Get() A {
	panic("Not implemented")
}

// True if not empty
func (option Option[A]) IsDefined() bool {
	panic("Not implemented")
}

// True if empty
func (option Option[A]) IsEmpty() bool {
	panic("Not implemented")
}

// Returns a Some containing the result of applying f to this Option's value if this Option is nonempty. Otherwise return None.
func (option Option[A]) Map(f func(A) A) A {
	panic("Not implemented")
}

// Returns a Some containing the result of applying f to this Option's value if this Option is nonempty. Otherwise return None.
func MapOption[A, B comparable](option Option[A], f func(A) B) Option[B] {
	panic("Not implemented")
}

// True if not empty
func (option Option[A]) NonEmpty() bool {
	panic("Not implemented")
}

// Evaluate and return alternate optional value if empty
func (option Option[A]) OrElse(alternative Option[A]) Option[A] {
	panic("Not implemented")
}

func (option Option[T]) String() string {
	if option.defined {
		return fmt.Sprintf("Some(%v)", option.value)
	}
	return "None"
}

// Returns a Sequence containing the Option's value if it is nonempty, or the empty list if the Option is empty.
func (option Option[A]) ToSeq() Seq[A] {
	panic("Not implemented")
}

// Returns a slice containing the Option's value if it is nonempty, or the empty list if the Option is empty.
func (option Option[A]) ToSlice() []A {
	panic("Not implemented")
}

// Converts an Option of a pair into an Option of the first element and an Option of the second element.
func UnZipOption[A, B comparable](pair Option[Tuple[A, B]]) Tuple[Option[A], Option[B]] {
	panic("Not implemented")
}

// Returns a Some formed from this Option and another Option by combining the corresponding elements in a pair. If either of the two Options is empty, None is returned.
func Zip[A, B comparable](option Option[A], another Option[B]) Option[Tuple[A, B]] {
	panic("Not implemented")
}
