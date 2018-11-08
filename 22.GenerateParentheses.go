package main

import "fmt"

func generateParenthesis(n int) []string {
	result := &[]string{}
	parenthesis(n, 0, "", result)
	return *result
}

func parenthesis(m, n int, s string, result *[]string) {
	if m == 0 && n == 0 {
		*result = append(*result, s)
	}
	if m > 0 {
		parenthesis(m - 1, n + 1, s+"(", result)
	}
	if n > 0 {
		parenthesis(m, n - 1, s+")", result)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}