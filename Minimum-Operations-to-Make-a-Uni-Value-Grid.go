func minOperations(grid [][]int, x int) int {
	answer := 0
	flat := make([]int, len(grid)*len(grid[0]))
	index := 0 
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			flat[index] = grid[i][j]
			index++
		}
	}
	slices.Sort(flat)
	median := flat[len(flat)/2]

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			answer += int(math.Abs(float64(median - grid[i][j]))) / x
			if (median - grid[i][j]) % x != 0 {
				return -1
			}
		}
	}
	return answer
}