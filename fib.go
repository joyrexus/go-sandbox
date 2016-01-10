package main

import "fmt"

func fib(n int) int {
	if n < 1 {
		return 1
	}
	x, y := 1, 1
	for i := 1; i <= n; i++ {
		x, y = y, x+y
	}
	return x
}

func fibr(n int) int {
	if n < 2 {
		return 1
	}
	return fibr(n-1) + fibr(n-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, fib(i), fibr(i))
	}
}
