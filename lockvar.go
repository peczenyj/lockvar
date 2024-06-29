package goprotected

import (
	"sync"
)

type MutexVar[T any] struct {
	variable T
	mutex    sync.Mutex
}

func New[T any](variable T) *MutexVar[T] {
	return &MutexVar[T]{
		variable: variable,
	}
}

func (v *MutexVar[T]) Use(fn func(variable T)) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	fn(v.variable)
}
