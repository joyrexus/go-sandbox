package main

import (
	"fmt"
)

func main() {
	var (
		u uint8 = 5
		v uint8 = 4
		x uint8 = 3
		y uint8 = 2
	)

	fmt.Println("\nbitwise AND")
	fmt.Printf("%08b %s %08b = %08b\n", u, "&", u, u & u)
	fmt.Printf("%08b %s %08b = %08b\n", u, "&", v, u & v)
	fmt.Printf("%08b %s %08b = %08b\n", u, "&", x, u & x)
	fmt.Printf("%08b %s %08b = %08b\n", u, "&", y, u & y)

	fmt.Println("\nbitwise OR")
	fmt.Printf("%08b %s %08b = %08b\n", u, "|", u, u | u)
	fmt.Printf("%08b %s %08b = %08b\n", u, "|", v, u | v)
	fmt.Printf("%08b %s %08b = %08b\n", u, "|", x, u | x)
	fmt.Printf("%08b %s %08b = %08b\n", u, "|", y, u | y)

	fmt.Println("\nbitwise XOR")
	fmt.Printf("%08b %s %08b = %08b\n", u, "^", u, u ^ u)
	fmt.Printf("%08b %s %08b = %08b\n", u, "^", v, u ^ v)
	fmt.Printf("%08b %s %08b = %08b\n", u, "^", x, u ^ x)
	fmt.Printf("%08b %s %08b = %08b\n", u, "^", y, u ^ y)

	fmt.Println("\nbit clear (AND NOT)")
	fmt.Printf("%08b %s %08b = %08b\n", u, "&^", u, u &^ u)
	fmt.Printf("%08b %s %08b = %08b\n", u, "&^", v, u &^ v)
	fmt.Printf("%08b %s %08b = %08b\n", u, "&^", x, u &^ x)
	fmt.Printf("%08b %s %08b = %08b\n", u, "&^", y, u &^ y)

	fmt.Println("\nleft shift")
	fmt.Printf("%08b %s %08b = %08b\n", u, "<<", v, u << v)
	fmt.Printf("%08b %s %08b = %08b\n", u, "<<", x, u << x)
	fmt.Printf("%08b %s %08b = %08b\n", u, "<<", y, u << y)

	fmt.Println("\nright shift")
	fmt.Printf("%08b %s %08b = %08b\n", u, ">>", v, u >> v)
	fmt.Printf("%08b %s %08b = %08b\n", u, ">>", x, u >> x)
	fmt.Printf("%08b %s %08b = %08b\n", u, ">>", y, u >> y)

	fmt.Println("\nbitwise negation (or complement)")
	fmt.Printf("^%08b = %08b\n", u, ^u)
	fmt.Printf("^%08b = %08b\n", v, ^v)
	fmt.Printf("^%08b = %08b\n", x, ^x)
	fmt.Printf("^%08b = %08b\n", y, ^y)
}
