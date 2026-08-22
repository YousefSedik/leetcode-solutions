class Solution:
    def checkDivisibility(self, n: int) -> bool:
        return (n % (sum(int(i) for i in str(n)) + math.prod(int(i) for i in str(n)))) == 0