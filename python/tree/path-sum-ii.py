import copy

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def __init__(self):
        self.stash = []

    def pathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: List[List[int]]
        """
        if root is None:
            return []
        result = []
        self.help(root, sum, result)
        return result

    def help(self, root, sum, result):
        self.stash.append(root.val)
        if root.left is None and root.right is None:
            if sum == root.val:
                result.append(self.stash)
            return
        if root.left is not None:
            self.help(root.left, sum - root.val, result)
        if root.right is not None:
            self.help(root.right, sum - root.val, result)
        self.stash.pop()

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
    origin = [5, 4, 8, 11, None, 13, 4, 7, 2, None, None, 5, 1]
    s = Solution()
    a = s.toBST(origin)
    print(s.pathSum(a, 22))
