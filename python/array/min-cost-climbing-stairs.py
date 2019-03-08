class Solution:
    def minCostClimbingStairs(self, cost):
        """
        :type cost: List[int]
        :rtype: int
        """
        length = len(cost)
        if length == 1:
            return cost[0]
        if length == 2:
            return min(cost[0], cost[1])
        dp = dict()
        dp[0] = cost[0]
        dp[1] = cost[1]
        i = 2
        while i < length:
            dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
            i += 1
        return min(dp[length-1], dp[length-2])

if __name__ == '__main__':
    cost = [0, 1, 1, 0]
    s = Solution()
    print(s.minCostClimbingStairs(cost))