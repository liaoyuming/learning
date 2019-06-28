class Solution:
    def getSum(self, a, b):
        """
        :type a: int
        :type b: int
        :rtype: int
        """
        m = a ^ b
        n = (a & b) << 1
        if n is not 0:
            return self.getSum(m, n)
        return m

if __name__ == '__main__':
    s = Solution()
    print(s.getSum(-1, 1))