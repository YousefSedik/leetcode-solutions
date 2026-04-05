func judgeCircle(moves string) bool {
	freq := make(map[rune]int)
	for _, move := range moves {
		freq[move]++
	}

	return freq['U'] == freq['D'] && freq['L'] == freq['R']
}