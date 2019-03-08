class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def __init__(self):
        self.result = list()

    def inorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if root is not None:
            self.inorderTraversal(root.left)
            self.result.append(root.val)
            self.inorderTraversal(root.right)
        return self.result

    def inorderTraversal2(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if root is None:
            return self.result
        stash = []
        pre = root
        while stash or pre:
            if pre is not None:
                stash.append(pre)
                pre = pre.left
            else:
                node = stash.pop()
                self.result.append(node.val)
                pre = node.right
        return self.result
