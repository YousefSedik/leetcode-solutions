func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := len(nums)-3; i > -1; i-- {
		cond1 := nums[i]+nums[i+1] > nums[i+2]
		cond2 := nums[i+1]+nums[i+2] > nums[i]
		cond3 := nums[i+2]+nums[i] > nums[i+1]
		if cond1 && cond2 && cond3 {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}
	return 0
}