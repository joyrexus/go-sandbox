package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%V, %V", NewPerson(), foo())
}

type Person struct {
	Name string
}

func NewPerson() (peep *Person) {
	// return &Person{"Bill"}
	// return peep // nil-valued
	return // nil-valued
}

func foo() *bytes.Buffer {
	return nil // &bytes.Buffer{} // new(bytes.Buffer)
}
