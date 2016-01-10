package main

import (
	"fmt"
	"github.com/joyrexus/stringutil"
)

func Hi(name string) string {
	return fmt.Sprintf("hello, %v\n", name)
}

func main() {
	fmt.Println("hello, world!")
	fmt.Print(Hi("Jason"))
	fmt.Println(stringutil.Reverse("Nice!"))
}
