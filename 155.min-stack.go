/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
    minpoint int
    values []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{minpoint: -1, values: make([]int, 0, 5)}
}


func (this *MinStack) Push(x int)  {
    if this.minpoint == -1 || this.values[this.minpoint] > x {
        this.minpoint = len(this.values)
    }
    this.values = append(this.values, x)
}


func (this *MinStack) Pop()  {
    this.values = this.values[:len(this.values) - 1]
    if this.minpoint == len(this.values) {
        this.minpoint = 0
        for i := 1; i < len(this.values); i++ {
            if this.values[i] < this.values[this.minpoint] {
                this.minpoint = i
            }
        }
	}
	if len(this.values) == 0 {
        this.minpoint = -1
    }
}


func (this *MinStack) Top() int {
    return this.values[len(this.values) - 1]
}


func (this *MinStack) GetMin() int {
    return this.values[this.minpoint]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

