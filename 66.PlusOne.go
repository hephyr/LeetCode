package main

import "fmt"

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] == 10 {
			digits[i] = 0
			carry = 1
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	a := []int{4,3,2,1}
	result := plusOne(a)
	fmt.Println(result)
}
