class Solution:
    is_common = set()

    def checkIfExistsAndAdd(self, n):
        if n in self.is_common:
            return True
        
        self.is_common.add(n)
        return False
    
    def findThePrefixCommonArray(self, A: List[int], B: List[int]) -> List[int]:
        n = len(A)
        counter = 0
        c = [0] * n
        for i in range(n):
            counter += 1 if self.checkIfExistsAndAdd(A[i]) else 0
            counter += 1 if self.checkIfExistsAndAdd(B[i]) else 0
            c[i] = counter
        
        self.is_common.clear()
        return c