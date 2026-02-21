func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func countPrimeSetBits(left int, right int) int {
	counter := 0
	for i := left; i <= right; i++ {
		if (isPrime(bits.OnesCount(uint(i)))){
			counter++
		}
	}
	return counter
}