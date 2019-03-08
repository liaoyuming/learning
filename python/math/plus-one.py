class Solution:
    def plusOne(self, digits):
        """
        :type digits: List[int]
        :rtype: List[int]
        """
        length = len(digits)
        i = length - 1
        j = 1
        while i >= 0 and j == 1:
            t = digits[i] + j
            digits[i] = t % 10
            j = 0
            if t >= 10:
                j = 1
            i -= 1
        if j == 1:
            digits = [1] + digits
        return digits

if __name__ == '__main__':
    s = Solution()
    print(s.plusOne([9, 9, 9]))
