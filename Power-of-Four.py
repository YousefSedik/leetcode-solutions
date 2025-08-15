class Solution:
    def isPowerOfFour(self, n: int) -> bool:
        n_binary = bin(n).removeprefix("0b")
        return n >= 0 and n_binary.count('0') % 2 == 0 and n_binary.count('1') == 1