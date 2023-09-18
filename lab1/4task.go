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

	fmt.Println(addElements(arr, add))

}

func addElements(arr, add []int) []int {
	mass := make([]int, 0)
	mass = append(mass, arr...)
	mass = append(mass, add...)
	return mass
}
