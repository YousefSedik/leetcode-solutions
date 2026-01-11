1class Solution:
2    def twoSum(self, nums: list[int], target: int) -> list[int]:
3        nums_original = nums.copy()
4        nums.sort()
5        p1, p2 = 0, len(nums)-1
6        fnum, snum = 0, 0
7        while p1 < p2:
8            if nums[p1] + nums[p2] == target:
9                fnum = nums[p1]
10                snum = nums[p2]
11                break
12            elif nums[p1] + nums[p2] < target:
13                p1 += 1
14            else:
15                p2 -= 1
16        
17        answer = []
18        for i in range(len(nums)):
19            if nums_original[i] == fnum:
20                answer.append(i)
21                fnum = None
22        
23            elif nums_original[i] == snum:
24                answer.append(i)
25                snum = None
26
27        return answer