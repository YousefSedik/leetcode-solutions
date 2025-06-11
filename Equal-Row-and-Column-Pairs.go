func equalPairs(grid [][]int) int {
	rowMapper := make(map[string]int)
	n := len(grid)

	for i := 0; i < n; i++ {
		var row strings.Builder
		for j := 0; j < n; j++ {
			row.WriteString(fmt.Sprint(grid[i][j]))
			row.WriteByte(',')
		}
		rowMapper[row.String()]++
	}

	answer := 0
	for i := 0; i < n; i++ {
		var col strings.Builder
		for j := 0; j < n; j++ {
			col.WriteString(fmt.Sprint(grid[j][i]))
			col.WriteByte(',')
		}
		answer += rowMapper[col.String()]
	}

	return answer
}
