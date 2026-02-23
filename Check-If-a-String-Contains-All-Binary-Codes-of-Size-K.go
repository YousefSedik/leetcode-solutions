func hasAllCodes(s string, k int) bool {
	leadingBitValue := int(math.Pow(2, float64(k-1)))
	maxCounterLength := int(math.Pow(2, float64(k)))
	counter := make(map[int]bool, 0)
	windowValue,count := 0, 1
	for i := 0; i < min(k, len(s)); i++ {
		if s[i] == '1' {
			windowValue += int(math.Pow(2, float64(k-1-i)))
		}
	}
	counter[windowValue] = true
	
	for i := k; i < len(s); i++ {
		if s[i-k] == '1' {
			windowValue -= leadingBitValue
		}
		windowValue *= 2
		if s[i] == '1' {
			windowValue++
		}
		if _, pres := counter[windowValue]; pres == false {
			counter[windowValue] = true
			count++
			if count == maxCounterLength {
				return true
			}
		}
	}
	return false
}