class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isBalanced(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if not root:
            return True
        return self.dfs(root) != -1

    def dfs(self, root):
        if not root:
            return 0
        left = self.dfs(root.left)
        right = self.dfs(root.right)
        if left == -1 or right == -1 or abs(left-right) > 1:
            return -1
        return max(left, right) + 1

    def toBST(self, arr, i=1):
        if not arr:
            return None
        if i > len(arr):
            return None
        if arr[i-1] is None:
            return None
        root = TreeNode(arr[i-1])
        root.left = self.toBST(arr, 2*i)
        root.right = self.toBST(arr, 2*i+1)
        return root


if __name__ == '__main__':
    origin = [1, 2, 2, 3, None, None, 3, 4, None, None, 4]
    s = Solution()
    a = s.toBST(origin)
    print(s.isBalanced(a))

