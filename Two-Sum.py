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
16        answer = []
17        for i in range(len(nums)):
18            if nums_original[i] == fnum:
19                answer.append(i)
20                fnum = None
21            elif nums_original[i] == snum:
22                answer.append(i)
23                snum = None
24
25        return answer