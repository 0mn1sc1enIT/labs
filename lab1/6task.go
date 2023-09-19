package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	var k int
	fmt.Scan(&k)

	fmt.Println(deleteElem(arr, k))
}

func deleteElem(arr []int, k int) []int {
	return arr[:k]
}
