from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        length = len(height)
        if length <= 2:
            return 0
        start = 0
        end = length - 1
        res = 0
        m = 0
        while start < end:
            if height[start] <= m:
                res += m - height[start]
                start += 1
            elif height[end] <= m:
                res += m - height[end]
                end -= 1
                continue
            else:
                m = min(height[start], height[end])
        return res

if __name__ == '__main__':
    s = Solution()
    print(s.trap([2, 0, 1, 0, 2]))

