class Solution:
    def largestGoodInteger(self, num: str) -> str:
        res = ""
        for i in range(len(num) - 2):
            if res and res[0] == '9':
                return '999'
            elif num[i] == num[i+1] == num[i+2]:
                res = max(res, num[i]*3)
        
        return res