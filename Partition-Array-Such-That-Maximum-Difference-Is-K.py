class Solution:
    def partitionArray(self, nums: List[int], k: int) -> int:
        nums.sort()
        start = nums[0]
        numberOfSubSeq = 0
        for i in range(1, len(nums)):
            if nums[i] - start > k: # start new subsequence
                numberOfSubSeq += 1
                start = nums[i]  
        if not numberOfSubSeq: # numberOfSubSeq == 0
            return 1
        return numberOfSubSeq + 1