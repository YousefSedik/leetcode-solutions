class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        largest_number = max(nums)
        max_count = 0
        count = 0
        for i in range(len(nums)):
            if nums[i] == largest_number:
                count += 1
                max_count = max(max_count, count)
            else:
                count = 0

        return max_count  