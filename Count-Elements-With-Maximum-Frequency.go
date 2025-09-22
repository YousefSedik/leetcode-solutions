func maxFrequencyElements(nums []int) int {
	freq := [101]int{}
	max_freq, answer := 0, 0
	for _, v := range nums {
		freq[v-1] += 1
		max_freq = max(max_freq, freq[v-1])
	}
    
	for i := range 101 {
		if freq[i] == max_freq {
			answer += freq[i]
		}
	}

	return answer
}
