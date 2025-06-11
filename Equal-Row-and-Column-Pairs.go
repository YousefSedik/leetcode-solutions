func equalPairs(grid [][]int) int {
	row_mapper := make(map[string]int)
	col_mapper := make(map[string]int)

	for i := 0; i < len(grid); i++ {
		var row strings.Builder
		var col strings.Builder
		for j := 0; j < len(grid); j++ {
			row.WriteString(fmt.Sprintf("%d ", grid[i][j]))
			col.WriteString(fmt.Sprintf("%d ", grid[j][i]))
		}

		row_mapper[row.String()] += 1
		col_mapper[col.String()] += 1

	}
	answer := 0
	
	for key, value := range row_mapper {
		answer += (value * col_mapper[key])
	}
	return answer
}