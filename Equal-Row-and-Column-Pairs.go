func equalPairs(grid [][]int) int {
	row_mapper := make(map[string]int)
	col_mapper := make(map[string]int)
	row, col := "", ""
	answer := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			row += string(grid[j][i]) + " "
			col += string(grid[i][j]) + " "
		}
		row_mapper[row] += 1
		col_mapper[col] += 1

		row = ""
		col = ""
	}
	for key, value := range row_mapper {
		if col_mapper[key] != 0{
			answer += (value * col_mapper[key])
		}
	}
	
	return answer
}