package semaphore_test

import (
	"context"
	"fmt"

	"github.com/peczenyj/go-protected/semaphore"
)

func ExampleNew() {
	protectedVar := semaphore.New(new(int), 1)

	ctx := context.Background()

	protectedVar.Use(ctx, 1, func(i *int) {
		*i++
	})

	protectedVar.Do(ctx, func(i *int) {
		fmt.Println("got i=", *i)
	})

	// Output:
	// got i= 1
}
