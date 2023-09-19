package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	var M int
	fmt.Scan(&M)

	add := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&add[i])
	}

	var k int
	fmt.Scan(&k)

	fmt.Println(addElementsFrom(arr, add, k))
}

func addElementsFrom(arr, add []int, k int) []int {
	mass := make([]int, 0)
	if k > len(arr) {
		k = len(arr)
	}
	mass = append(mass, arr[:k]...)
	mass = append(mass, add...)
	mass = append(mass, arr[k:]...)
	return mass
}
