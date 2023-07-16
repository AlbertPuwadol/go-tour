package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	var list *List[int]
	for i := 10; i > 0; i-- {
		node := &List[int]{val: i}
		node.next = list
		list = node
	}
	for node := list; node != nil; node = node.next {
		fmt.Println(node.val)
	}
}
