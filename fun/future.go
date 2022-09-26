package fun

import "sync"

// A Future represents a value which may or may not currently be available, but will be available at some point.
type Future[A any] struct {
	value *Option[A]
	ch    chan A
	mtx   *sync.Mutex
}

// Creates a new Future by applying a function (which itself returns Future) to the result of this Future.
func (future *Future[A]) FlatMap(f func(A) Future[A]) Future[A] {
	return FutureValue(func() A {
		inner := f(future.Result())
		return inner.Result()
	})
}

// Creates a new Future from the provided function that will asynchronously produce a value A at some point.
func FutureValue[A any](f func() A) Future[A] {
	future := Future[A]{new(Option[A]), make(chan A), new(sync.Mutex)}

	go func() {
		r := f()
		future.ch <- r
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

// Creates a new Future by applying a function to the result of this Future.
func (future *Future[A]) Map(f func(A) A) Future[A] {
	return FutureValue(func() A {
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
	future.mtx.Lock()
	if !future.IsCompleted() {
		result := Some(<-future.ch)
		*future.value = result
	}
	future.mtx.Unlock()
	return future.value.Get()
}
