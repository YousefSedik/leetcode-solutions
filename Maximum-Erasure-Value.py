class Solution:
    def maximumUniqueSubarray(self, nums: List[int]) -> int:
        n = len(nums)        
        mp = defaultdict(int)
        mp[nums[0]] += 1
        p1, p2, answer, current_sum = 0, 0, nums[0], nums[0]
        while n > p2+1:
            p2 += 1
            if mp[nums[p2]] == 1:
                while mp[nums[p2]] == 1:
                    current_sum -= nums[p1]
                    mp[nums[p1]] -= 1
                    p1 += 1
                current_sum += nums[p2]
                mp[nums[p2]] = 1
            else:
                mp[nums[p2]] = 1
                current_sum += nums[p2]
            answer = max(answer, current_sum)

        return answer
