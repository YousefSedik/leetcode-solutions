func isTrionic(nums []int) bool {
	decIndexStart, decIndexEnd := -1, -1 
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1]{
			decIndexStart = i
			break
		}
	}
	if decIndexStart == 0 || decIndexStart == len(nums) - 1 || decIndexStart == -1 {
		return false
	}
	for i := len(nums) -1 ; i > 0; i-- {
		if nums[i-1] >= nums[i]{
			decIndexEnd = i
			break
		}
	}
	if decIndexEnd == len(nums) - 1 {
		return false
	}
	for i := decIndexStart; i < decIndexEnd; i++ {
		if nums[i] <= nums[i+1]{
			return false
		}
	}
	return true
}