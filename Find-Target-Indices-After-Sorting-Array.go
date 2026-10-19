func targetIndices(nums []int, target int) []int {
	slices.Sort(nums)
	lowerBound := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})

	upperBound := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})

	answer := make([]int, upperBound-lowerBound)
	for i := lowerBound; i < upperBound; i++ {
		answer[i-lowerBound] = i
	}
	
	return answer
}
