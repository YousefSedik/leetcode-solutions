func longestBalanced(nums []int) int {
	n := len(nums)
	globalAnswer := 0
	localAnswer := 0

	for i := 0; i < n; i++ {
		Exists := make(map[int]bool, n)
		evenCounter, oddCounter := 0, 0
		for j := i; j < n; j++ {
			if !Exists[nums[j]] {
				if nums[j]%2 == 0 {
					evenCounter++
				} else {
					oddCounter++
				}
				Exists[nums[j]] = true

			}
			if evenCounter == oddCounter {
				localAnswer = j - i + 1
			}
		}
		globalAnswer = max(globalAnswer, localAnswer)
	}
	return globalAnswer
}