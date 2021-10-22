class Solution(object):
    def findTheDifference(self, s, t):

        stringS = list(s)
        print(stringS)
        stringT = list(t)
        print(stringT)
        answer = [x for x in stringT if x not in stringS]
        if len(answer) == 0:
            return s
        return answer[0]

sol = Solution()
s = "ae"
t = "aea"

ans = sol.findTheDifference(s,t)
print(ans)
