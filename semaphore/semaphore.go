package semaphore

import (
	"context"

	"golang.org/x/sync/semaphore"
)

// ProtectedVar is a go protected var but uses a weighted semaphore.
type ProtectedVar[T any] struct {
	variable T
	weighted *semaphore.Weighted
}

// New creates a new weighted semaphore protected variable with the given
// maximum combined weight for concurrent access.
func New[T any](variable T, weight int64) *ProtectedVar[T] {
	return &ProtectedVar[T]{
		variable: variable,
		weighted: semaphore.NewWeighted(weight),
	}
}

// Use method. Allow access the weighted semaphore protected var in a thread-safe callback.
// The method acquires the semaphore with a weight of n, blocking until resources
// are available or ctx is done. On success, executes the callback and returns nil.
// On failure, returns ctx.Err() and leaves the semaphore unchanged.
func (v *ProtectedVar[T]) Use(ctx context.Context, n int64, callback func(T)) error { //nolint:varnamelen //not needed.
	err := v.weighted.Acquire(ctx, n)
	if err != nil {
		return err //nolint:wrapcheck //not needed.
	}

	defer v.weighted.Release(n)

	callback(v.variable)

	return nil
}

// Do is the equivalent of call Use method but with n equal to 1.
func (v *ProtectedVar[T]) Do(ctx context.Context, callback func(T)) error {
	return v.Use(ctx, 1, callback)
}

// TryUse acquires the semaphore with a weight of n without blocking.
// On success, executed the callback and returns true.
// On failure, returns false and leaves the semaphore unchanged.
func (v *ProtectedVar[T]) TryUse(n int64, callback func(T)) bool {
	if !v.weighted.TryAcquire(n) {
		return false
	}

	defer v.weighted.Release(n)

	callback(v.variable)

	return true
}
