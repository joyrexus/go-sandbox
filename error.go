package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("%s", errors.New("Hi!"))
}
