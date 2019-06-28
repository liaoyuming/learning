class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def __init__(self):
        self.result = []

    def postorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if root is not None:
            self.postorderTraversal(root.left)
            self.postorderTraversal(root.right)
            self.result.append(root.val)
        return self.result

    def postorderTraversal2(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if root is None:
            return self.result
        stash = [root]
        while stash:
            node = stash.pop()
            self.result = [node.val] + self.result
            if node.left is not None:
                stash.append(node.left)
            if node.right is not None:
                stash.append(node.right)
        return self.result

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
    origin = [1, 2, 2, 3]
    s = Solution()
    a = s.toBST(origin)
    print(s.postorderTraversal2(a))

