package main

import (
	"fmt"
	"strings"
)

func main() {
	WordCount := func(s string) map[string]int {
		count := make(map[string]int)
		for _, w := range strings.Fields(s) {
			count[w] += 1
		}
		return count
	}

	fmt.Println(WordCount("ahem ahem hello world"))
}
