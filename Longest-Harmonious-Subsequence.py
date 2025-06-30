class Solution:
    def findLHS(self, nums: List[int]) -> int:
        freq = defaultdict(lambda: 0)
        for i in nums:
            freq[i] += 1
        
        res = 0
        for k in freq:
            if freq.get(k+1, 0) != 0:
                res = max(res, freq[k] + freq[k+1])
        
        return res