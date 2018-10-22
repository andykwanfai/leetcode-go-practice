
import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	backtracking(candidates, []int{}, target, 0, &result)
	return result

}

func backtracking(candidates, solution []int, target, start int, result *[][]int) {
	if target < 0 {
		return
	} else if target == 0 {
		*result = append(*result, solution)
		return
	} else {
		for i := start; i < len(candidates); i++ {
			solution = solution[:len(solution):len(solution)]
			solution = append(solution, candidates[i])
			backtracking(candidates, solution, target-candidates[i], i, result)
			solution = solution[:len(solution)-1]
		}
	}
}