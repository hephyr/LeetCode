package main

import "fmt"

func maxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		var area int
		if height[i] > height[j] {
			area = (j - i) * height[j]
			j--
		} else {
			area = (j - i) * height[i]
			i++
		}
		if area > max {
			max = area
		}
	}
	return max
}

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(a))
}
