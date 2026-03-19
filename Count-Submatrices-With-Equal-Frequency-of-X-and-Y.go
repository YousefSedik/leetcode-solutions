func numberOfSubmatrices(grid [][]byte) int {
	counter := 0
	rows, cols := len(grid), len(grid[0])
	partial_sum_x := make([][]int, rows+1)
	partial_sum_y := make([][]int, rows+1)


	for i := 0; i < rows+1; i++ {
        partial_sum_x[i] = make([]int, cols+1)
    }
	for i := 0; i < rows+1; i++ {
        partial_sum_y[i] = make([]int, cols+1)
    }
	for i := 1; i < rows+1; i++ {
		for j := 1; j < cols+1; j++ {
			X, Y := 0, 0
			if grid[i-1][j-1] == 'X'{
				X = 1
			}
			if grid[i-1][j-1] == 'Y'{
				Y = 1
			}
			
			partial_sum_x[i][j] = X + partial_sum_x[i-1][j] + partial_sum_x[i][j-1] - partial_sum_x[i-1][j-1]
			partial_sum_y[i][j] = Y + partial_sum_y[i-1][j] + partial_sum_y[i][j-1] - partial_sum_y[i-1][j-1]
			if partial_sum_x[i][j] == partial_sum_y[i][j] && partial_sum_x[i][j] != 0{
				counter++
			}
		}
	}

	return counter
}