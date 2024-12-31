package semaphore_test

import (
	"context"
	"fmt"
	"time"

	"github.com/peczenyj/go-protected/semaphore"
)

func ExampleNew() {
	protectedVar := semaphore.New(new(int), 1)

	ctx := context.Background()

	protectedVar.Use(ctx, 1, func(i *int) {
		*i++
	})

	protectedVar.TryUse(1, func(i *int) {
		*i++
	})

	protectedVar.Do(ctx, func(i *int) {
		fmt.Println("got i=", *i)
	})

	// Output:
	// got i= 2
}

func ExampleProtectedVar_TryUse() {
	protectedVar := semaphore.New(new(int), 1)

	success := protectedVar.TryUse(10, func(i *int) {
		fmt.Println("got i=", *i) // will never be called
	})

	fmt.Printf("TryUse returns %t", success)

	// Output:
	// TryUse returns false
}

func ExampleProtectedVar_Use() {
	protectedVar := semaphore.New(new(int), 1)

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	err := protectedVar.Use(ctx, 10, func(i *int) {
		fmt.Println("got i=", *i) // will never be called
	})

	fmt.Printf("Use returns %v", err)

	// Output:
	// Use returns context deadline exceeded
}
