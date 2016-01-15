package main

import "fmt"

func main() {
	// operators of the same precedence level associate left to right
	x := 10 / 2 * 5 // (10 / 2) * 5
	fmt.Println(x)  // 25

	y := 2 * 5 / 10 // (2 * 5) / 10
	fmt.Println(y)	// 1
}
