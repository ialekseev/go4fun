package fun

import "sync"

// A Future represents a value which may or may not currently be available, but will be available at some point.
type Future[A any] struct {
	value *Option[A]
	ch    chan A
	mtx   *sync.Mutex
}

// Returns a new Future by applying a function (which itself returns Future) to the result of this Future (while keeping the same Future value's type A).
func (future *Future[A]) FlatMap(f func(A) Future[A]) Future[A] {
	return FlatMapFuture(future, f)
}

// Returns a new Future by applying a function (which itself returns Future) to the result of this Future (potentially, changing the Future value's type A => B).
func FlatMapFuture[A, B any](future *Future[A], f func(A) Future[B]) Future[B] {
	return FutureValue(func() B {
		inner := f(future.Result())
		return inner.Result()
	})
}

// Creates a new Future from the provided function that will asynchronously produce a value A at some point.
func FutureValue[A any](f func() A) Future[A] {
	future := Future[A]{new(Option[A]), make(chan A), new(sync.Mutex)}

	go func() {
		future.ch <- f()
	}()

	go func() {
		future.Result()
	}()

	return future
}

// Returns whether the Future has already been completed with a value.
func (future *Future[A]) IsCompleted() bool {
	return future.value.IsDefined()
}

// Returns a new Future by applying a function to the result of this Future (while keeping the same Future value's type A).
func (future *Future[A]) Map(f func(A) A) Future[A] {
	return MapFuture(future, f)
}

// Returns a new Future by applying a function to the result of this Future (potentially, changing the Future value's type A => B).
func MapFuture[A, B any](future *Future[A], f func(A) B) Future[B] {
	return FutureValue(func() B {
		return f(future.Result())
	})
}

// When this Future is completed apply the provided function on its value.
func (future *Future[A]) OnComplete(f func(A)) {
	go func() {
		f(future.Result())
	}()
}

// Await and return the result (of type A) of this Future.
func (future *Future[A]) Result() A {
	if !future.IsCompleted() {
		future.mtx.Lock()
		if !future.IsCompleted() {
			*future.value = Some(<-future.ch)
		}
		future.mtx.Unlock()
	}
	return future.value.Get()
}

// Returns the current value of this Future. If the Future was not completed the returned value will be None. If the Future was completed the value will be Some(value).
func (future *Future[A]) Value() Option[A] {
	return *future.value
}
