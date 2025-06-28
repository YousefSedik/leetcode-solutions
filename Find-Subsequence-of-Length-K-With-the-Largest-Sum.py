class Solution:
    def maxSubsequence(self, nums: List[int], k: int) -> List[int]:
        pq = []
        for i in nums:
            if len(pq) < k:
                heapq.heappush(pq, i)
            else:
                smallest = heapq.heappop(pq)
                heapq.heappush(pq, max(smallest, i))
        
        mapper = {}
        for i in pq:
            mapper[i] = mapper.get(i, 0) + 1
            
        answer = [0] * k
        answerIndex = 0
        for i in nums:
            if mapper.get(i):
                answer[answerIndex] = i
                mapper[i] -= 1
                answerIndex += 1
                
        return answer
