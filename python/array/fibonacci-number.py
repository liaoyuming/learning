class Solution:
    def fib(self, N):
        """
        :type N: int
        :rtype: int
        """
        dp = dict()
        dp[0] = 0
        dp[1] = 1
        i = 2
        while i <= N:
            dp[i] = dp[i-1] + dp[i-2]
            i += 1
        return dp[N]
