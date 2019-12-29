func combinationSum(candidates []int, target int) [][]int {
    result := make([][]int, 0)
    solve(candidates, target, []int{}, &result)
    return result
}

func solve(candidates []int, rest int, present []int, result *[][]int) {
    if rest == 0 {
		temp := make([]int, len(present))
		copy(temp, present)
		*result = append(*result, temp)
        return
    } else if rest < 0 {
        return
    }
    if len(candidates) == 0 {
        return
    } else {
        solve(candidates, rest - candidates[0], append(present, candidates[0]), result)
        if len(candidates) > 1 {
            solve(candidates[1:], rest, present, result)
        }
    }
}