package main

import "fmt"

func inc(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x + y + z
		}
	}
}

func main() {
	tenPlus := inc(10)
	zeroPlus := inc(0)
	tenPlusTwoPlus := tenPlus(2)
	zeroPlusTwoPlus := zeroPlus(2)
	tenPlusTenPlus := tenPlus(10)
	fmt.Println(tenPlusTwoPlus(5))
	fmt.Println(zeroPlusTwoPlus(5))
	fmt.Println(tenPlusTenPlus(5))
}
