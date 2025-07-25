class Solution:
    def maxSum(self, nums: List[int]) -> int:
        largest = -100
        found = [False] * (101)
        answer = 0
        for i in nums:
            largest = max(largest, i)
            if i > 0 and not found[i]:
                answer += i
                found[i] = True
        
        return answer if answer > 1 else largest