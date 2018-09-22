func searchInsert(nums []int, target int) int {
	//binary search
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	if target > nums[left] {
		return left + 1
	} else {
		return left
	}
}