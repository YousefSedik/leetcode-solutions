func kthCharacter(k int) byte {
	word := strings.Builder{}
	word.WriteRune(97)
	for len(word.String()) < k {
		wo := word.String()
		for i := 0; i < len(wo); i++ {
			if wo[i] == 'z' {
				word.WriteRune(97)
			} else {
				word.WriteRune(rune(wo[i] + 1))
			}
		}
	}
	return word.String()[k-1]
}