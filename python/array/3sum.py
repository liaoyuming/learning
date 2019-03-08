class Solution:
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = list()
        length = len(nums)
        nums.sort()
        i = 0
        while i < length-2 and nums[i] <= 0:
            if i > 0 and nums[i-1] == nums[i]:
                i = i+1
                continue
            j = i+1
            k = length-1
            while j<k and nums[k] >= 0:
                t = nums[i] + nums[j]
                if t > -nums[k]:
                    k = k - 1
                elif t < -nums[k]:
                    j = j + 1
                else:
                    res.append([nums[i], nums[j], nums[k]])
                    j = j+1
                    k = k-1
            i = i+1
        return list(set([tuple(t) for t in res]))

if __name__ == '__main__':
    # i = -100
    # nums = []
    # while i < 100:
    #     nums += [i]
    #     i = i+1
    nums = [-2, 0, 0, 2, 2]
    # print(nums)
    s = Solution()
    print(s.threeSum(nums))
