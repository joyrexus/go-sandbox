package main

import (
	"fmt"
)

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list of integers.
func (l *IntList) Sum() int {
	if l.Tail == nil {
		return 0
	}
	return l.Value + l.Tail.Sum()
}

func main() {
	n := 4
	list := &IntList{n, nil}

	for i := n; i > 1; i-- {
		list = &IntList{i, list}
	}

	// list := &IntList{4, &IntList{3, &IntList{}}}
	fmt.Println(list.Sum())
}
