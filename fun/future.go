package fun

import "sync"

type Future[A comparable] struct {
	ch    chan A
	value Option[A]
	mtx   *sync.Mutex
}

func NewFuture[A comparable](f func() A) Future[A] {
	future := Future[A]{make(chan A), None[A](), new(sync.Mutex)}

	go func() {
		r := f()
		future.ch <- r
		close(future.ch)
	}()

	return future
}

func (future *Future[A]) IsCompleted() bool {
	return future.value.IsDefined()
}

// Await and return a result (of type A) of this Future.
func (future *Future[A]) Result() A {
	future.mtx.Lock()
	if !future.IsCompleted() {
		future.value = Some(<-future.ch)
	}
	future.mtx.Unlock()
	return future.value.value
}

// When this Future is completed apply the provided function on its value.
func (future *Future[A]) OnComplete(f func(A)) {
	go func() {
		f(future.Result())
	}()
}

func (future *Future[A]) Map(f func(A) A) Future[A] {
	return NewFuture(func() A {
		return f(future.Result())
	})
}

func (future *Future[A]) FlatMap(f func(A) Future[A]) Future[A] {
	return NewFuture(func() A {
		inner := f(future.Result())
		return inner.Result()
	})
}
