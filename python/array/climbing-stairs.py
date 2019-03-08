class Solution:
    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        dp = dict()
        dp[0] = 1
        dp[1] = 1
        i = 2
        while i <= n:
            dp[i] = dp[i - 1] + dp[i - 2]
            i += 1
        return dp[n]