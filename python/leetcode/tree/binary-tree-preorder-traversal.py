class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def __init__(self):
        self.result = []

    def preorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if root is None:
            return self.result
        self.result.append(root.val)
        self.preorderTraversal(root.left)
        self.preorderTraversal(root.right)
        return self.result

    def preorderTraversal2(self, root):
        if root is None:
            return self.result
        stash = [root]
        while stash:
            node = stash.pop()
            self.result.append(node.val)
            if node.right is not None:
                stash.append(node.right)
            if node.left is not None:
                stash.append(node.left)
        return self.result


