class Solution:
    def countMaxOrSubsets(self, nums: List[int]) -> int:
        max_or = 0
        for i in nums:
            max_or |= i

        def solve(i, _or):
            if i == len(nums):
                if max_or == _or:
                    return 1
                return 0
            return solve(i + 1, _or) + solve(i + 1, _or | nums[i])

        return solve(0, 0)

