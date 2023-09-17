package main

import (
	"fmt"
)


type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
	elementCount int
}

func (ls LinkedList) Init(v int, n *Node) (int, *Node) {
	ls.head.value = v
	ls.head.next = n
	return v, n
}

func (ls LinkedList) Size() int {
	return ls.elementCount
}

func (ls LinkedList) Empty() bool {
	return ls.elementCount == 0
}

func (ls LinkedList) pushFront(value int) {
	ls.head.value, ls.head.next = ls.Init(value, ls.head.next)
	ls.elementCount++
}




func main() {
	fmt.Println()
}