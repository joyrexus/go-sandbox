package main

import (
	"os"
	"fmt"
)

func main() {
	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		fmt.Println("OH NOES!")
	}
	if true {
		fmt.Println("True!")
	}
	if !false {
		fmt.Println("True too!")
	}
}
