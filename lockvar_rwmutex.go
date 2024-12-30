package goprotected

import (
	"sync"
)

// RWMutexVar read-write mutex-protected var.
type RWMutexVar[T any] struct {
	variable T
	mutex    sync.RWMutex
}

// New constructor.
// Wraps the variable with a read-write mutex-protected wrapper.
func NewRW[T any](variable T) *RWMutexVar[T] {
	return &RWMutexVar[T]{
		variable: variable,
	}
}

// Use method. Allow access the read-write  mutex-protected var in a thread-safe callback for writing.
func (v *RWMutexVar[T]) Use(fn func(variable T)) {
	v.mutex.Lock()

	fn(v.variable)

	v.mutex.Unlock()
}

// RUse method. Allow access the read-write mutex-protected var in a thread-safe callback for reading.
func (v *RWMutexVar[T]) RUse(fn func(variable T)) {
	v.mutex.RLock()

	fn(v.variable)

	v.mutex.RUnlock()
}
