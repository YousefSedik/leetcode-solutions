class Solution:
    def numSquares(self, n: int) -> int:
        options = [i ** 2 for i in range(1, 101)]
        mem = [-1] * (n + 1)
        def solve(sum):
            if sum == 0:
                return 1
            elif sum < 0:
                return 0
            if mem[sum] != -1:
                return mem[sum]
            mi = float("inf")
            for option in options:
                ret = solve(sum - option)
                if ret:
                    mi = min(mi, 1 + ret)
            mem[sum] = mi
            return mi

        return solve(n) - 1