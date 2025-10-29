func smallestNumber(n int) int {
	answer := 1
	for n != 0 {
		n /= 2
		answer *= 2
	}
	return answer - 1
}