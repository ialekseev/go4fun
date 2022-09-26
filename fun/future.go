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

func (future Future[A]) IsCompleted() bool {
	return future.value.IsDefined()
}

func (future Future[A]) read() A {
	var v A
	future.mtx.Lock()
	if !future.IsCompleted() {
		v = <-future.ch
	} else {
		v = future.value.Get()
	}
	future.mtx.Unlock()
	return v
}

func (future Future[A]) OnComplete(f func(A)) {
	go func() {
		f(future.read())
	}()
}
