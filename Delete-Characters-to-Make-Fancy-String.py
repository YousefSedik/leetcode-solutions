class Solution:
    def makeFancyString(self, s: str) -> str:
        n = len(s)
        if n < 3: 
            return s
        
        res = []
        for i in range(n - 2):
            if s[i] == s[i + 1] and s[i + 1] == s[i + 2]:
                continue
            else:
                res.append(s[i])
            
        res.append(s[n-2])
        res.append(s[n-1])
        
        return ''.join(res)