// println.go
package main

import (
	"fmt"
)

func ExamplePrinln() {
	// As with typical tests, examples are functions that reside in a package's _test.go files.
	// Unlike normal test functions, though, example functions take no arguments and begin with the word Example instead of Test.
	// $ go test -v

	fmt.Println("hello")
	// Output: hello
}
