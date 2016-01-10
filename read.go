package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		panic(err)
	}
    fmt.Println(string(b))
}
