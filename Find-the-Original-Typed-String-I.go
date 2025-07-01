func possibleStringCount(word string) int {
	answer := 1
	for i := 0; i < len(word)-1; i++ {
		if word[i] == word[i+1] {
			answer++
		}
	}
	return answer
}