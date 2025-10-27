func numberOfBeams(bank []string) int {
	row1, row2, answer := 0, 0, 0
	for i := 1; i < len(bank); i++ {
		if strings.Contains(bank[i], "1") {
			row1 = row2
			row2 = i
		}
		if row1 != row2 && row2 == i {
			answer += strings.Count(bank[row1], "1") * strings.Count(bank[row2], "1")
		}
	}
	
	return answer
}