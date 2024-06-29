package goprotected_test

import (
	"fmt"

	goprotected "github.com/peczenyj/go-protected"
)

func ExampleNew() {
	lockVar := goprotected.New(new(int))

	lockVar.Use(func(i *int) {
		*i++
	})

	lockVar.Use(func(i *int) {
		fmt.Println("got i=", *i)
	})

	// Output:
	// got i= 1
}
