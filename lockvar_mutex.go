package goprotected

import (
	"sync"
)

// MutexVar mutex-protected var.
type MutexVar[T any] struct {
	variable T
	mutex    sync.Mutex
}

// New constructor.
// Wraps the variable with a mutex-protected wrapper.
func New[T any](variable T) *MutexVar[T] {
	return &MutexVar[T]{
		variable: variable,
	}
}

// Use method. Allow access the mutex-protected var in a thread-safe callback.
func (v *MutexVar[T]) Use(fn func(variable T)) {
	v.mutex.Lock()

	fn(v.variable)

	v.mutex.Unlock()
}
