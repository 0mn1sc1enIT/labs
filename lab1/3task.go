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
		var b bool
		for j := 0; j < N; j++ {
			if A[i] == B[j] {
				b = true
			}
		}
		if !b {
			C = append(C, A[i])
		}
	}

	for i := 0; i < N; i++ {
		var b bool
		for j := 0; j < M; j++ {
			if B[i] == A[j] {
				b = true
			}
		}
		if !b {
			C = append(C, B[i])
		}
	}
	fmt.Println(deleteDupl(C))
}

func deleteDupl(arr []int) []int {
	mass := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		var b bool
		for j := 0; j < len(mass); j++ {
			if arr[i] == mass[j] {
				b = true
				break
			}
		}
		if !b {
			mass = append(mass, arr[i])
		}
	}

	return mass
}
