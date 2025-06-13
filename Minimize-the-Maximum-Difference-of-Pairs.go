func minimizeMax(nums []int, p int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	answer := 0
	l, r := 0, nums[len(nums)-1]
	for l <= r {
		// can the answer be mid ?
		mid := l + (r-l)/2
		counter := 0
		for i := 0; i < len(nums)-1; i++ {
			if max(nums[i]-nums[i+1], (nums[i]-nums[i+1])*-1) <= mid {
				counter++
				if counter == p { // no need to continue
					break
				}
				i++
			}
		}
		if counter == p {
			answer = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return answer
}