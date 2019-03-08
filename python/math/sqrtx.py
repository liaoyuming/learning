class Solution:
    def mySqrt(self, x):
        """
        :type x: int
        :rtype: int
        """
        if x == 0:
            return 0
        if x == 1:
            return 1
        return self.bs(x, 0, x/2)

    def bs(self, x, start, end):
        if start > end:
            return None
        mid = int((end - start)/2 + start)
        m = mid * mid
        n = (mid + 1) * (mid + 1)
        if m <= x < n:
            return mid
        elif x == n:
            return mid + 1
        elif x > n:
            return self.bs(x, mid + 1, end)
        else:
            return self.bs(x, start, mid)

if __name__ == '__main__':
    s = Solution()
    print(s.mySqrt(3))