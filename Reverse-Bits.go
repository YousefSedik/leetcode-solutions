func reverseBits(n int) int {
	binaryString := fmt.Sprintf("%032b", n)
	answer := 0
	binaryValue := 1
	for i := 0; i < 32; i++ {
		if binaryString[i] == '1' {
			answer += binaryValue
		}
		binaryValue *= 2
	}
	return answer
}