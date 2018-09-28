func generate(numRows int) [][]int {
	result := [][]int{}
	if numRows == 0 {
		return result
	}
	result = append(result, []int{1})
	if numRows == 1 {
		return result
	}
	for i := 1; i < numRows; i++ {
		result = append(result, nextRow(result[i-1]))
	}
	return result
}

func nextRow(previousRow []int) []int {
	next := []int{0}
	next = append(next, previousRow...)

	for i := 0; i < len(next)-1; i++ {
		next[i] += next[i+1]
	}
	return next
}