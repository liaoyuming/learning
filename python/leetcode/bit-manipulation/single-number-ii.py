class Solution:
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return int((3 * sum(set(nums))-sum(nums))/2)

if __name__ == '__main__':
    nums = [2, 2, 2, 2, 3]
    s = Solution()
    print(s.singleNumber(nums))