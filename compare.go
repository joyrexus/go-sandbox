package main

import "fmt"

func main() {
	var a, b int = 42, 42
    fmt.Println(a == b)

    var i, j interface{} = 42, 42 // a, b
    fmt.Println(i == j)

	var s, t struct{ i interface{} }
    s.i, t.i = a, b
    fmt.Println(s == t)

	// var q, r struct{ s []string }
    // fmt.Println(q == r)
	// invalid operation since a struct is comparable 
	// only if its fields are comparable
}

