func hasIncreasingSubarrays(nums []int, k int) bool {
	for i := 0; i < len(nums)-(2*k-1); i++ { // 0 => 3
		possible := true
		for j := i; j < i+k*2-1; j++ {
			if nums[j] < nums[j+1] || j == i+k-1 {
				continue
			} else {
				possible = false
				break
			}
		}
		if possible {
			return true
		}
	}
	return false
}
