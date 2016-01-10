package main

import "fmt"

type Person interface {
    Greet() string
}

type Bob string

func (b *Bob) Greet() string {
	return "Hello!"
}

type Gerhard struct {
	Greeting string
}

func (g Gerhard) Greet() string {
	return g.Greeting
}
	
func main() {
	bob := new(Bob)
	gerhard := &Gerhard{"Guten Tag!"}
	people := []Person{
		bob,
		gerhard,
	}
	for _, p := range people {
		fmt.Println(p.Greet())
	}
}
