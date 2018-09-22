func removeElement(nums []int, val int) int {
	//two pointer
	i := 0
	for _, e := range nums {
		if e != val {
			nums[i] = e
			i++
		}
	}
	return i
}