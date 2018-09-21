func removeDuplicates(nums []int) int {
	//two pointer
	left, right, size := 0, 1, len(nums)
	for ; right < size; right++ {
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left] = nums[right]
	}
	return left + 1
}