package goprotected_test

import (
	"fmt"

	goprotected "github.com/peczenyj/go-protected"
)

func ExampleNew() {
	protectedVar := goprotected.New(new(int))

	protectedVar.Use(func(i *int) {
		*i++
	})

	protectedVar.Use(func(i *int) {
		fmt.Println("got i=", *i)
	})

	// Output:
	// got i= 1
}

func ExampleNewRW() {
	rwProtectedVar := goprotected.NewRW(new(int))

	rwProtectedVar.Use(func(i *int) {
		*i += 5
	})

	rwProtectedVar.RUse(func(i *int) {
		fmt.Println("[read-only] got i=", *i)
	})

	// Output:
	// [read-only] got i= 5
}
