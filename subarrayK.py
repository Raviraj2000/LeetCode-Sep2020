class Solution(object):
    def numSubarrayProductLessThanK(self, nums, k):
        subarrays = [nums[i:i+j] for i in range(0,len(nums)) for j in range(1,len(nums)-i+1)]
        answer = []
        print(subarrays)
        for li in subarrays:
            product = 1
            for j in li:
                product = product * j
            if product < k:
                answer.append(li)
        print(answer)
        print(len(answer))
        return len(answer)



s = Solution()
nums = [10, 5, 2, 6]
s.numSubarrayProductLessThanK(nums, 100)
