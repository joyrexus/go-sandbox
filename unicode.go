package main

import (
	"fmt"
)

func main() {
	c := string(0x4eac)
	fmt.Println(c)
	fmt.Println(string(20140))
	fmt.Println("\xe4\xba\xac")
	fmt.Println("\u4eac")
	fmt.Printf("% x\n", c)
	fmt.Println("京" == c)

	r := []rune(c)
	fmt.Printf("%x\n", r)
	fmt.Printf("%016b\n", r[0]) // 0100111010101100
	//         E4            BA             AC
	//       0100       11 1010        10 1100
	// [1110]0100 + [10]11 1010  + [10]10 1100
	fmt.Printf("%016b\n", '京') // 0100111010101100

	fmt.Println('H', string(72), "\x48")
	fmt.Printf("\\x%x\n", 'H')
}
