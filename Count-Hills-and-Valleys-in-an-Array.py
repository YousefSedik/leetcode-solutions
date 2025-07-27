class Solution:
    def countHillValley(self, nums: list[int]) -> int:
        count = 0
        nums_2 = [nums[0]]
        for i in range(1, len(nums)):
            if nums_2[-1] != nums[i]:
                nums_2.append(nums[i])

        for i in range(len(nums_2) - 2):
            if nums_2[i] < nums_2[i + 1] > nums_2[i + 2]:
                count += 1
            elif nums_2[i] > nums_2[i + 1] < nums_2[i + 2]:
                count += 1
        
        return count
