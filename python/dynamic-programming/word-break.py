class Solution:
    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: List[str]
        :rtype: bool
        """
        length = len(s)
        dp = [False for i in range(length+1)]
        dp[0] = True
        i = 1
        while i < length + 1:
            j = i-1
            while j >= 0:
                if dp[j] and s[j:i] in wordDict:
                    dp[i] = True
                    break
                j -= 1
            i += 1
        return dp[len(s)]

if __name__ == '__main__':
    s = Solution()
    print(s.wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                       ["a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"]))
    # print(s.wordBreak('leetcode', ["leet", "code"]))

