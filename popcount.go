package main

import "fmt"

func main() {
	/*
	var pc [20]byte
	
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("pc[%d] = pc[%d/2] + byte(%d&1)\n", i, i, i)
		fmt.Printf("%d set bits for %d (%08b): %08b = %08b + %08b\n\n", pc[i], i, i, pc[i], pc[i/2], byte(i&1))
	}
	*/

	/*
	var x uint16
	x = 955
	fmt.Println(byte(x))
	fmt.Printf("%016b\n", x)
	fmt.Printf("%016b\n", x>>0)
	fmt.Printf("%016b\n", x>>8)
	*/

	var x uint64 = 655
	fmt.Printf("%012b\n", x)
	fmt.Printf("%012b\n", byte(x))
	fmt.Printf("%012b\n", byte(x>>8))
}
