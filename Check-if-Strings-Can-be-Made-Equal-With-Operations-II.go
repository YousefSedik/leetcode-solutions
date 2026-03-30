func checkStrings(s1 string, s2 string) bool {
	
    if s1 == s2 {
		return true
	}

	for i := 0; i < min(len(s1), 2); i++ {

		freqS1 := make(map[rune]int)
		freqS2 := make(map[rune]int)
		
		for j := i; j < len(s1); j+=2 {
			freqS1[rune(s1[j])]++
		}
		for j := i; j < len(s2); j+=2 {
			freqS2[rune(s2[j])]++
		}
		for k, v := range freqS1 {
			if freqS2[k] != v {
				return false
			}
		}
	}
	return true
}