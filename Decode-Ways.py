class Solution:
    def numDecodings(self, s: str) -> int:
        mem = dict()

        def solve(i: int) -> int:
            if i == len(s):
                return 1

            if s[i] == '0':
                return 0

            if i in mem:
                return mem[i]

            opt1 = solve(i + 1)

            opt2 = 0
            if i + 1 < len(s) and int(s[i:i+2]) <= 26:
                opt2 = solve(i + 2)

            mem[i] = opt1 + opt2
            return mem[i]

        return solve(0)
