func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := hashTable[target-num]; ok {
			return []int{j, i}
		}
		hashTable[num] = i
	}
	return nil
}

/*func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
}*/