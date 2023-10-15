package main

import (
	"fmt"
)

func close(s string) bool {
	return s == "}" || s == ")" || s == "]"
}

func open(s string) bool {
	return s == "{" || s == "(" || s == "["
}

func isValid(s string) bool {
	stack := make([]string, 0)
	m := map[string]string{"{": "}", "[": "]", "(": ")"}

	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if open(char) {
			stack = append(stack, char)
		}
		if close(char) {
			if len(stack) == 0 || char != m[stack[len(stack)-1]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	var line string
	fmt.Scan(&line)

	if isValid(line) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
