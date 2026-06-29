func numOfStrings(patterns []string, word string) int {
	counter := 0 
	for i := 0; i < len(patterns); i++ {
		if strings.Contains(word, patterns[i]){
			counter++
		}
	}

	return counter
}