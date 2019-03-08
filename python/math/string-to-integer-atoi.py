import re
from _testcapi import INT_MIN
from _testcapi import INT_MAX

class Solution:
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        s = re.search('^([\+\-]?)\d+', str.strip())
        if s:
            s = s.group()
        else:
            return 0
        s = int(s)
        if s > INT_MAX :
            return INT_MAX
        if s < INT_MIN:
            return INT_MIN
        return s

if __name__ == '__main__':
    s = Solution()
    print(s.myAtoi("-+1"))