class Solution:
    def sumZero(self, n: int) -> List[int]:
        if n == 1:
            return [0]
        
        numbers = [1] * n
        for i in range(n-1):
            numbers[i] = i+1
        
        numbers[n-1] = -n*(n-1) // 2
        return numbers