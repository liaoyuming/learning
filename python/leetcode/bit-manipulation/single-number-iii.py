class Solution:
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        a = 0
        for n in nums:
            a = a ^ n
        c = 0
        while a & 1 == 0:
            a = a >> 1
            c += 1
        j = 0
        k = 0
        for n in nums:
            if n >> c & 1:
                j = j ^ n
            else:
                k = k ^ n
        return j, k

if __name__ == '__main__':
    nums = [2, 2, 3, 3, 5, 1]
    s = Solution()
    print(s.singleNumber(nums))