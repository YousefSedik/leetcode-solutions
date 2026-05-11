class Solution:
    def separateDigits(self, nums: List[int]) -> List[int]:
        answer = []
        for number in nums:
            tmp = []
            while number > 0:
                tmp.append(number % 10)
                number //= 10
            answer.extend(tmp[::-1])
        return answer