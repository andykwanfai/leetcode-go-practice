package pascals_triangle_II

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	index := 1
	for i := 0; i <= rowIndex; i++ {
		result[i] = index
		index = index * (rowIndex - i) / (i + 1)
	}
	return result
}
