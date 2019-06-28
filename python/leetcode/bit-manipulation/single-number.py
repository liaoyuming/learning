class Solution:
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        res = 0
        for n in nums:
            res ^= n
        return res

if __name__ == '__main__':
    nums = [2, 2, 2, 2, 2, 2, 3, 3, 3]
    s = Solution()
    print(s.singleNumber(nums))