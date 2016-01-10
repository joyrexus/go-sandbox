package main

import (
	"fmt"
	"strings"
)

func main() {
	letters := []string{"A", "B", "C", "D"}
	fmt.Println(strings.Join(letters, ", "))
	fmt.Println("a" + "b" + "c")
}
