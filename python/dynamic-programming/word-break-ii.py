import copy
import collections

class Solution:
    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: List[str]
        :rtype: List[str]
        """
        if not wordDict or not s:
            return []
        length = len(s)
        dp = [[] for i in range(length+1)]
        i = 1
        while i < length+1:
            j = i-1
            while j >= 0:
                if (j == 0 or dp[j]) and s[j:i] in wordDict:
                    dp[i].append([s[j:i], dp[j]])
                j -= 1
            i += 1
        return dp[len(s)]

    def wordBreak2(self, s, wordDict):
        return self.dfs(s, wordDict, dict())

    def dfs(self, s, wordDict, tmp):
        if s in tmp:
            return tmp[s]
        if len(s) == 0:
            return ['']
        res = []
        for word in wordDict:
            if s.startswith(word):
                sublist = self.dfs(s[len(word):], wordDict, tmp)
                for sub in sublist:
                    res.append((word + " " + sub).strip())
        tmp[s] = res
        return res

if __name__ == '__main__':
    s = Solution()
    print(s.wordBreak4("aaaaaaaaaaaaaaaaaaaaaaaaa", ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]))
    # print(s.wordBreak("catsanddog", ["cat", "cats", "and", "sand", "dog"]))


