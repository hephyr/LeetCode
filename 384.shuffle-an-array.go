/*
 * @lc app=leetcode id=384 lang=golang
 *
 * [384] Shuffle an Array
 */

// @lc code=start
type Solution struct {
    nums []int
}


func Constructor(nums []int) Solution {
    return Solution{nums: nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.nums
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    result := make([]int, len(this.nums), len(this.nums))
    copy(result, this.nums)
    rand.Shuffle(len(result), func(i, j int) {
        result[i], result[j] = result[j], result[i]
    })
    return result
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

