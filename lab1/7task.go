package main

import "fmt"

func main() {
	arr := []int{10, 5, 3, 6, 7, 8, 13, 55}
	p := &arr

	length := len(arr)
	l := &length

	p = noPrimeNum(p, l)
	fmt.Println(*p)
}

func noPrimeNum(p *[]int, l *int) *[]int {
	mass := make([]int, 0)
	for i := 0; i < *l; i++ {
		if !isPrime((*p)[i]) {
			mass = append(mass, (*p)[i])
		}
	}
	newP := &mass
	return newP
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
