class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

class Solution:
    def levelOrderBottom(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        r = dict()
        self.help(root, r, 0)
        res = list()
        l = len(r)
        while l > 0:
            res.append(r[l-1])
            l = l-1
        return res

    def help(self, root, result, l):
        if root is None:
            return
        if l in result:
            result[l].append(root.val)
        else:
            result[l] = [root.val]
        self.help(root.left, result, l+1)
        self.help(root.right, result, l+1)


    def toBST(self, arr, i=1):
        if not arr:
            return None
        if i > len(arr):
            return None
        if arr[i - 1] is None:
            return None
        root = TreeNode(arr[i - 1])
        root.left = self.toBST(arr, 2 * i)
        root.right = self.toBST(arr, 2 * i + 1)
        return root

if __name__ == '__main__':
    origin = [3,9,20, None, None,15,7]
    s = Solution()
    a = s.toBST(origin)
    print(s.levelOrderBottom(a))


