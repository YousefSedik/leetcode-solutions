class Solution:
    def countSubmatrices(self, grid: List[List[int]], k: int) -> int:
        counter = 0
        rows, cols = len(grid), len(grid[0])
        partial_sum = [[0] * (cols + 1) for _ in range(rows + 1)]
        
        for i in range(1, rows + 1):
            for j in range(1, cols + 1):
                partial_sum[i][j] = (
                    grid[i - 1][j - 1]
                    + partial_sum[i - 1][j] 
                    + partial_sum[i][j - 1]
                    - partial_sum[i - 1][j - 1]
                )
        
        for i in range(rows):
            for j in range(cols):
                if partial_sum[i+1][j+1] <= k:
                    counter += 1
                else:
                    break
        
        return counter
