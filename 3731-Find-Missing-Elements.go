func findMissingElements(nums []int) []int {
	// range -> (start, end)
	start, end := slices.Min(nums), slices.Max(nums)
	// store values in a map 
	exists := make(map[int]bool, 0)

	for i := range(nums) {
		exists[nums[i]] = true
	}
	list := make([]int, 0, end - start + 1 - len(nums))

	for i := start+1; i < end; i++ {
		if exists[i] == false{
			list = append(list, i)
		}
	}
	return list
}
