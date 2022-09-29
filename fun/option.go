package fun

import (
	"fmt"
)

type Option[A any] struct {
	value   A
	defined bool
}

// Returns a Some containing a result of applying a function f(A)=>B to Option[A]'s value if the Option is nonempty. Otherwise returns None.
// An Alias for Map function.
func ApplyOption1[A, B any](option Option[A], f func(A) B) Option[B] {
	return MapOption(option, f)
}

// Returns a Some containing a result of applying a binary function f(A,B)=>C to Option[A] & Option[B]'s values if both Options are nonempty. Otherwise returns None.
func ApplyOption2[A, B, C any](optionA Option[A], optionB Option[B], f func(A, B) C) Option[C] {
	bc := MapOption(optionA, Curry2(f))
	return FlatMapOption(bc, func(bc func(B) C) Option[C] {
		return MapOption(optionB, func(b B) C {
			return bc(b)
		})
	})
}

// Returns a Some containing a result of applying a function of 3 arguments f(A,B,C)=>D to Option[A] & Option[B] & Option[C]'s values if all 3 Options are nonempty. Otherwise returns None.
func ApplyOption3[A, B, C, D any](optionA Option[A], optionB Option[B], optionC Option[C], f func(A, B, C) D) Option[D] {
	bcd := MapOption(optionA, Curry3(f))
	return FlatMapOption(bcd, func(bcd func(B) func(C) D) Option[D] {
		return FlatMapOption(optionB, func(b B) Option[D] {
			return MapOption(optionC, func(c C) D {
				return bcd(b)(c)
			})
		})
	})
}

// Returns true if this Option has an element that is equal (as determined by ==) to elem, false otherwise.
func ContainsInOption[A comparable](option Option[A], elem A) bool {
	return option.defined && option.value == elem
}

// Returns true if this Option is nonempty and the predicate p returns true when applied to this Option's value. Otherwise, returns false.
func (option Option[A]) Exists(p func(A) bool) bool {
	return option.defined && p(option.value)
}

// Returns this Option if it is nonempty and applying the predicate p to this Option's value returns true. Otherwise, returns None.
func (option Option[A]) Filter(p func(A) bool) Option[A] {
	if option.defined && p(option.value) {
		return option
	} else {
		return None[A]()
	}
}

// Returns this Option if it is nonempty and applying the predicate p to this Option's value returns false. Otherwise, returns None.
func (option Option[A]) FilterNot(p func(A) bool) Option[A] {
	if option.defined && !p(option.value) {
		return option
	} else {
		return None[A]()
	}
}

// Returns the result of applying f to this Option's value if this Option is nonempty (without changing Option value's type A).
// Returns None if this Option is empty. Different from map in that f is expected to return an Option (which could be None).
func (option Option[A]) FlatMap(f func(A) Option[A]) Option[A] {
	return FlatMapOption(option, f)
}

// Returns the result of applying f to this Option's value if this Option is nonempty (potentially, changing Option value's type A => B).
// Returns None if this Option is empty. Different from map in that f is expected to return an Option (which could be None).
func FlatMapOption[A, B any](option Option[A], f func(A) Option[B]) Option[B] {
	if option.defined {
		return f(option.value)
	} else {
		return None[B]()
	}
}

// Returns the nested Option value if this Option is nonempty.
func FlattenOption[A any](option Option[Option[A]]) Option[A] {
	if option.defined {
		return option.value
	} else {
		return None[A]()
	}
}

// Returns the result of applying f to this Option's value if the Option is nonempty. Otherwise, returns defaultValue.
// Resulting value's type A is the same as Option value's type A.
func (option Option[A]) Fold(defaultValue A, f func(A) A) A {
	return FoldOption(option, defaultValue, f)
}

// Returns the result of applying f to this Option's value if the Option is nonempty. Otherwise, returns defaultValue.
// Resulting value's type B could, potentially, be different from the Option value's type A.
func FoldOption[A, B any](option Option[A], defaultValue B, f func(A) B) B {
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

// Applies the given procedure f to the Option's value, if it is nonempty. Otherwise, does nothing.
func (option Option[A]) Foreach(f func(A)) {
	if option.defined {
		f(option.value)
	}
}

// Returns the Option's value if the Option is nonempty, otherwise returns type A's default value.
func (option Option[A]) Get() A {
	if option.defined {
		return option.value
	} else {
		return *new(A)
	}
}

// Returns the Option's value if the Option is nonempty, otherwise returns defaultValue.
func (option Option[A]) GetOrElse(defaultValue A) A {
	if option.defined {
		return option.value
	} else {
		return defaultValue
	}
}

// Returns true if the Option is nonempty (has a value).
func (option Option[A]) IsDefined() bool {
	return option.defined
}

// Returns true if the Option is empty (doesn't have a value).
func (option Option[A]) IsEmpty() bool {
	return !option.defined
}

// Returns a Some containing the result of applying f to this Option's value if this Option is nonempty (without changing Option value's type A). Otherwise return None.
func (option Option[A]) Map(f func(A) A) Option[A] {
	return MapOption(option, f)
}

// Returns a Some containing the result of applying f to this Option's value if this Option is nonempty (potentially, changing Option value's type A => B). Otherwise return None.
func MapOption[A, B any](option Option[A], f func(A) B) Option[B] {
	if option.defined {
		return Some(f(option.value))
	} else {
		return None[B]()
	}
}

// None[A] represents a non-existing value of type A.
func None[A any]() Option[A] {
	return Option[A]{}
}

// Returns true if the Option is nonempty (has a value).
func (option Option[A]) NonEmpty() bool {
	return option.defined
}

// Returns this Option if it is nonempty, otherwise returns an alternative Option.
func (option Option[A]) OrElse(alternative Option[A]) Option[A] {
	if option.defined {
		return option
	} else {
		return alternative
	}
}

// Wrap a value A in Some[A]. Some[A] represents an existing value of type A.
func Some[A any](value A) Option[A] {
	return Option[A]{value, true}
}

// A String representation of Option. E.g. Some(5) or None
func (option Option[A]) String() string {
	if option.defined {
		return fmt.Sprintf("Some(%v)", option.value)
	}
	return "None"
}

// Returns a Sequence containing the Option's value if it is nonempty, or nil if the Option is empty.
func (option Option[A]) ToSeq() Seq[A] {
	if option.defined {
		return Seq[A]{option.value}
	} else {
		return nil
	}
}

// Converts an Option of a Tuple into a Tuple of 2 Options.
func UnZipOption[A, B any](pair Option[Tuple2[A, B]]) Tuple2[Option[A], Option[B]] {
	if pair.defined {
		return Tup2(Some(pair.value.a), Some(pair.value.b))
	} else {
		return Tup2(None[A](), None[B]())
	}
}

// Returns a Some formed from this Option and another Option by combining the corresponding elements in a Tuple. If either of the two Options is empty, None is returned.
func ZipOption[A, B any](option Option[A], another Option[B]) Option[Tuple2[A, B]] {
	if option.defined && another.defined {
		return Some(Tup2(option.value, another.value))
	} else {
		return None[Tuple2[A, B]]()
	}
}
