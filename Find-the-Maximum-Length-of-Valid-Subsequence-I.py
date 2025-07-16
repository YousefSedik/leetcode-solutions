class Solution:
    def maximumLength(self, nums: List[int]) -> int:
        all_even = sum([1 if n % 2 == 0 else 0 for n in nums])
        all_odd = sum([1 if n % 2 == 1 else 0 for n in nums])
        odd_then_even, even_then_odd = 0, 0
        last_is_even, last_is_odd = True, True

        for i in nums:
            if i % 2 == 1 and last_is_even:
                odd_then_even += 1
                last_is_even = False
            elif i % 2 == 0 and not last_is_even:
                odd_then_even += 1
                last_is_even = True

            if i % 2 == 1 and not last_is_odd:
                even_then_odd += 1
                last_is_odd = True
            elif i % 2 == 0 and last_is_odd:
                even_then_odd += 1
                last_is_odd = False

        return max(all_even, all_odd, odd_then_even, even_then_odd)

print(Solution().maximumLength([1, 2, 3, 4]))
