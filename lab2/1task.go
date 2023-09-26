package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode() *ListNode {
	return &ListNode{}
}

func NewListNodeWithValue(x int) *ListNode {
	return &ListNode{Val: x}
}

func NewListNodeWithValueAndNext(x int, next *ListNode) *ListNode {
	return &ListNode{Val: x, Next: next}
}

type MyLinkedList struct {
	Head *ListNode
	Size int
}

func LinkedList() MyLinkedList {
	return MyLinkedList{}
}

func (list *MyLinkedList) GetSize() int {
	return list.Size
}

func (list *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{Val: val}
	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	list.Size++
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index > list.Size {
		fmt.Println("Index is larger than the size=")
		return
	}
	if index == 0 {
		list.Head = list.Head.Next
	} else {
		current := list.Head
		for i := 0; i < index-1; i++ {
			current = current.Next
		}
		current.Next = current.Next.Next
	}
	list.Size--
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		list.AddAtHead(val)
	}
	newNode := &ListNode{Val: val}
	current := list.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	list.Size++
}

func (list *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{Val: val}
	newNode.Next = list.Head
	list.Head = newNode
	list.Size++
}

func (list *MyLinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func (list *MyLinkedList) Reverse() {
	var prev *ListNode
	current := list.Head

	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}

	list.Head = prev
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	linkedList := LinkedList()

	for _, val := range arr {
		linkedList.AddAtTail(val)
	}
	linkedList.Print()

	fmt.Println("Deleting the last element...")
	linkedList.DeleteAtIndex(4)
	linkedList.Print()

	fmt.Println("Add at index 3")
	linkedList.AddAtIndex(3, 10)
	linkedList.Print()

	fmt.Println("Add at head 15")
	linkedList.AddAtHead(15)
	linkedList.Print()

	fmt.Println("Reversing linked list")
	linkedList.Reverse()
	linkedList.Print()
}
