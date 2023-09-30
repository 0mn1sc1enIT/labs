package main

import "fmt"

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Prev *DoubleListNode
}

func NewDloubleListNode() *DoubleListNode {
	return &DoubleListNode{}
}

func NewDoubleListNodeWithValue(x int) *DoubleListNode {
	return &DoubleListNode{Val: x}
}

func NewDoubleListNodeWithValueAndNext(x int, next *DoubleListNode, prev *DoubleListNode) *DoubleListNode {
	return &DoubleListNode{Val: x, Next: next, Prev: prev}
}

type MyDoubleLinkedList struct {
	Head *DoubleListNode
	Size int
}

func DoubleLinkedList() MyDoubleLinkedList {
	return MyDoubleLinkedList{}
}

func (list *MyDoubleLinkedList) InsertNode(val int) {
	newNode := &DoubleListNode{Val: val}
	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (list *MyDoubleLinkedList) PrintD() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	linkedList := DoubleLinkedList()

	for _, val := range arr {
		linkedList.InsertNode(val)
	}
	linkedList.PrintD()
}
