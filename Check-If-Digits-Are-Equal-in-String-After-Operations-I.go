func hasSameDigits(s string) bool {
	temp := strings.Builder{}
	for len(s) != 2 { // 282
		for i := 0; i < len(s)-1; i++ {
			sum := ((int(s[i]) - '0') + (int(s[i+1]) - '0')) % 10
			temp.WriteRune(rune(sum + '0'))
		}
		s = temp.String()
		temp.Reset()
	}
	return s[0] == s[1]
}