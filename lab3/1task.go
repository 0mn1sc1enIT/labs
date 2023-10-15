package main

import (
	"fmt"
)

type FCS struct {
	a []int
	n int
}

func FixedCapacityStack(capacity int) *FCS {
	return &FCS{
		a: make([]int, 0, capacity),
		n: capacity,
	}
}

func (f *FCS) Empty() bool {
	return len(f.a) == 0
}

func (f *FCS) Full() bool {
	return f.n == len(f.a)
}

func (f *FCS) Push(item int) error {
	if !f.Full() {
		f.a = append(f.a, item)
		return nil
	}
	return fmt.Errorf("Достигнут лимит стека")
}

func (f *FCS) Pop() int {
	var popped int
	if !f.Empty() {
		popped = f.a[len(f.a)-1]
		f.a = f.a[:len(f.a)-1]
	}
	return popped
}

func (f *FCS) Top() int {
	if !f.Empty() {
		return f.a[len(f.a)-1]
	}
	return -1
}

func main() {
	var n int
	fmt.Scan(&n)
	stack := &FCS{n: n}

	fmt.Println(stack.Empty())
	fmt.Println(stack.Full())

	var item int
	ans := "y"

	for ans == "y" && !stack.Full() {
		fmt.Scan(&item)
		stack.Push(item)
		fmt.Println("Стэк:", stack.a)
		fmt.Println("Добавить элемент в стэк? Если да, то напишите y")
		fmt.Scan(&ans)
	}

	fmt.Println("Стэк пуст?", stack.Empty())
	fmt.Println("Стэк полон?", stack.Full())

	fmt.Println("Удалить вершину стэка? Если да, то напишите y")
	fmt.Scan(&ans)
	for ans == "y" && !stack.Empty() {
		fmt.Println("Удаленный элемент: ", stack.Pop())
		fmt.Println(stack.a)
		fmt.Println("Удалить вершину стэка? Если да, то напишите y")
		fmt.Scan(&ans)
	}

	fmt.Println("Стэк пуст?", stack.Empty())
	fmt.Println("Стэк полон?", stack.Full())
}
