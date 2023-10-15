package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func NewDoubleListNodeWithValue(x int) *Node {
	return &Node{Val: x}
}

type Queue struct {
	First *Node
	Last  *Node
	n     int
}

func queue() *Queue {
	return &Queue{}
}

func (q *Queue) Empty() bool {
	return q.First == nil
}

func (q *Queue) Enqueue(val int) {
	newNode := NewDoubleListNodeWithValue(val)
	q.n++
	if q.First == nil {
		q.First = newNode
		q.Last = newNode
		return
	}

	newNode.Prev = q.Last
	q.Last.Next = newNode
	q.Last = newNode
}

func (q *Queue) Dequeue() int {
	if q.Empty() {
		return -1
	}

	el := q.First.Val
	q.First = q.First.Next
	if q.First != nil {
		q.First.Prev = nil
	}
	q.n--
	return el
}

func (q *Queue) front() int {
	if q.Empty() {
		return -1
	}
	return q.First.Val
}

func (q *Queue) back() int {
	if q.Empty() {
		return -1
	}
	return q.Last.Val
}

func (q *Queue) Print() {
	var c int
	current := q.First
	for current != nil {
		fmt.Printf("%d ", current.Val)
		c++
		if c < q.n {
			fmt.Printf("<-> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func main() {
	arr := []int{4, 5, 7, 8, 2, 6}
	q := queue()

	for _, val := range arr {
		q.Enqueue(val)
		q.Print()
	}
	fmt.Printf("front:%d, back:%d\n", q.front(), q.back())
	for i := 0; i < len(arr); i++ {
		fmt.Println(q.Dequeue())
		q.Print()
		fmt.Printf("front:%d, back:%d\n", q.front(), q.back())
	}
}
