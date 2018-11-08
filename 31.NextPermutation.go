package main

import "fmt"

func nextPermutation(nums []int)  {
	smaller := len(nums) - 1
	if smaller == 0 || smaller == 1 {
		return
	}
	for ; smaller > 0 && nums[smaller] <= nums[smaller-1]; smaller-- {}
	smaller--
	if smaller < 0 {
		reverse(nums)
	} else {
		greater := len(nums) -1
		for nums[smaller] >= nums[greater] && greater > 0 {
			greater--
		}
		temp := nums[smaller]
		nums[smaller] = nums[greater]
		nums[greater] = temp
		reverse(nums[smaller+1:])
	}
	fmt.Println(nums)
}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		i++
		j--
	}
}

func main() {
	nextPermutation([]int{1, 3, 2})
	nextPermutation([]int{3, 2, 1})
	nextPermutation([]int{2, 3, 1})
	nextPermutation([]int{1, 2, 1})
	nextPermutation([]int{1, 1, 1})
	nextPermutation([]int{5, 1, 1})
}
