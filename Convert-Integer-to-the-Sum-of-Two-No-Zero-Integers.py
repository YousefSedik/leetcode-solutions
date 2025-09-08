class Solution:
    def getNoZeroIntegers(self, n: int) -> List[int]:
        answer = [n - 1, 1]
        while str(answer[0]).count("0") or str(answer[1]).count("0"):
            answer[0] -= 1
            answer[1] += 1

        return answer
