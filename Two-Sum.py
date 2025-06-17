class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        nums_original = nums.copy()
        nums.sort()
        p1, p2 = 0, len(nums)-1
        fnum, snum = 0, 0
        while p1 < p2:
            if nums[p1] + nums[p2] == target:
                fnum = nums[p1]
                snum = nums[p2]
                break
            elif nums[p1] + nums[p2] < target:
                p1 += 1
            else:
                p2 -= 1
        
        answer = []
        for i in range(len(nums)):
            if nums_original[i] == fnum:
                answer.append(i)
                fnum = None
        
            elif nums_original[i] == snum:
                answer.append(i)
                snum = None

        return answer