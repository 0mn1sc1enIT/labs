package main

import (
	"fmt"
)

func main() {
	var M, N int
	fmt.Println("Enter M and N:")
	fmt.Scan(&M, &N)

	A := make([]int, M)
	B := make([]int, N)
	fmt.Println("Enter A[i]:")
	for i := 0; i < M; i++ {
		fmt.Scan(&A[i])
	}
	fmt.Println("Enter B[i]:")
	for i := 0; i < N; i++ {
		fmt.Scan(&B[i])
	}

	C := make([]int, 0)

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if A[i] == B[j] {
				C = append(C, A[i])
			}
		}
	}

	fmt.Println(deleteDupl(C))
}

func deleteDupl(arr []int) []int {
	mass := make([]int, 0)

	for _, val := range arr {
		isDuplicate := false
		for _, resVal := range mass {
			if val == resVal {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			mass = append(mass, val)
		}
	}

	return mass
}
