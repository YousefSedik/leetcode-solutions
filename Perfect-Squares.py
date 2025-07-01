class Solution:
    def numSquares(self, n: int) -> int:
        options = [i ** 2 for i in range(1, 101)]
        dp = [float("inf")] * (n + 1)
        dp[0] = 0
        for i in range(1, n + 1):
            for option in options:
                if i - option < 0:
                    break
                
                dp[i] = min(dp[i], dp[i - 1] + 1, dp[i - option] + 1)

        return dp[n]