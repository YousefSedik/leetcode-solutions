# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def firstBadVersion(self, n: int) -> int:
        left, right, answer = 1, n, n
        while left <= right:
            mid = (left + right) // 2
            if isBadVersion(mid):
                answer = mid
                right = mid - 1
            else:
                left = mid + 1
        
        return answer
