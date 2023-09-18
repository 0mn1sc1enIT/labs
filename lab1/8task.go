package main

import "fmt"

func main() {
	arr := [10]int{-5, -10, 3, 8, 7, 17, -1, 4, -15, 0}
	p := &arr

	length := len(arr)
	l := &length

	fmt.Println(sort(p, l))
}

func sort(p *[10]int, l *int) ([]int, []int, []int) {
	zero := []int{}
	pos := []int{}
	neg := []int{}

	for i := 0; i < *l; i++ {
		switch {
		case (*p)[i] > 0:
			pos = append(pos, (*p)[i])
		case (*p)[i] == 0:
			zero = append(zero, (*p)[i])
		case (*p)[i] < 0:
			neg = append(neg, (*p)[i])
		}
	}
	return neg, zero, pos
}
