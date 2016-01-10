package main

import "fmt"

/*
	&    bitwise AND            integers
	|    bitwise OR             integers
	^    bitwise XOR            integers
	&^   bit clear (AND NOT)    integers
*/
func main() {
	fmt.Println(65/64, 65%64)
	bit := uint(62)
	fmt.Printf("%064b\n", 1 << bit) // 0000001
	fmt.Printf("%07b\n", 1 << 0) // 0000001
	fmt.Printf("%07b\n", 1 << 1) // 0000010
	fmt.Printf("%07b\n", 1 << 2) // 0000100

	/*
		n := uint(10)
		fmt.Printf("%07b\n", n) // 0001010
		m := uint(123)
			fmt.Printf("%07b\n", m)   // 1111011
			fmt.Printf("%07b\n", n&m) // 0001010
			fmt.Printf("%07b\n", n|m) // 1111011

		    // Use bitwise OR | to get the bits that are in 1 OR 2
		    // 1     = 00000001
		    // 2     = 00000010
		    // 1 | 2 = 00000011 = 3
		    fmt.Println(1 | 2)

		    // Use bitwise OR | to get the bits that are in 1 OR 5
		    // 1     = 00000001
		    // 5     = 00000101
		    // 1 | 5 = 00000101 = 5
		    fmt.Println(1 | 5)

		    // Use bitwise XOR ^ to get the bits that are in 3 OR 6 BUT NOT BOTH
		    // 3     = 00000011
		    // 6     = 00000110
		    // 3 ^ 6 = 00000101 = 5
		    fmt.Println(3 ^ 6)

		    // Use bitwise AND & to get the bits that are in 3 AND 6
		    // 3     = 00000011
		    // 6     = 00000110
		    // 3 & 6 = 00000010 = 2
		    fmt.Println(3 & 6)

		    // Use bit clear AND NOT &^ to get the bits that are in 3 AND NOT 6 (order matters)
		    // 3      = 00000011
		    // 6      = 00000110
		    // 3 &^ 6 = 00000001 = 1
		    fmt.Printf("%08b\n", 3)
		    fmt.Printf("%08b\n", 6)
		    fmt.Printf("%08b\n", 3 &^ 6)

	*/
}
