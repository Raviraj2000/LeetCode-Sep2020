class Solution(object):
    def largestOverlap(self, A, B):
        N = len(A)
        LA = [i / N * 100 + i % N for i in range(N * N) if A[i / N][i % N]]
        LB = [i / N * 100 + i % N for i in range(N * N) if B[i / N][i % N]]
        c = collections.Counter(i - j for i in LA for j in LB)
        return max(c.values() or [0])


A = [[1,1,0],
     [0,1,0],
     [0,1,0]]
B = [[0,0,0],
     [0,1,1],
     [0,0,1]]

s = Solution()
answer = s.largestOverlap(A,B)
print(answer)
