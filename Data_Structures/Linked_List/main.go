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

func (ls *LinkedList) Size() int {
	return ls.elementCount
}

func (ls *LinkedList) Empty() bool {
	return ls.elementCount == 0
}

func (ls *LinkedList) pushFront(value int) {
	// newNode := &Node{value: value}

	// if ls.head == nil {
	// 	ls.head = newNode
	// 	ls.elementCount++
	// 	return
	// }

	// current := ls.head
	// for current.next != nil {
	// 	current = current.next
	// }

	// current.next = newNode
	// ls.elementCount++

	//===============================
    //===============================
	// Stack variant
	newNode := &Node{value: value}
	
	if ls.head == nil {
		ls.head = newNode
		ls.elementCount++
		return
	}

	current := ls.head
	ls.head = newNode
	newNode.next = current
	ls.elementCount++
}

func (ls *LinkedList) popFront() {
	if ls.head == nil {
		return
	}

	ls.head = ls.head.next
	ls.elementCount--
}

func (ls *LinkedList) printList() []int {
	var elems []int
	for i := ls.head; i != nil; i = i.next {
		elems = append(elems, i.value)
	}
	return elems
}

func main() {
	ls := &LinkedList{}
	ls.pushFront(11)
	ls.pushFront(54)
	ls.pushFront(33)
	fmt.Println("List:", ls.printList())
	fmt.Println("Size:", ls.Size())
	printList(ls)
	ls.popFront()
	fmt.Println("List:", ls.printList())
	fmt.Println("Size:", ls.Size())
	printList(ls)
}

func printList(l *LinkedList) {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d ", curr.value)
		curr = curr.next
	}
	fmt.Println()
}