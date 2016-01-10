package main

import "fmt"

func main() {
	a := []byte("a")
	fmt.Println(string(a)) // a

	b := append(a, "b"...)
	fmt.Println(string(b)) // ab

	c := append(a, "c"...)
	fmt.Println(string(b)) // ac <- !!!
	fmt.Println(string(c)) // ac
}

