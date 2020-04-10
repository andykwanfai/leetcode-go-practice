package main

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for numbers[i]+numbers[j] != target {
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{i + 1, j + 1}
}
