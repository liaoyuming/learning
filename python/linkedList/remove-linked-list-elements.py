class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def removeElements(self, head, val):
        """
        :type head: ListNode
        :type val: int
        :rtype: ListNode
        """
        while head and head.val == val:
            head = head.next
        if head is None:
            return None
        h = head
        while h.next:
            if h.next.val == val:
                h.next = h.next.next
            else:
                h = h.next
        return head

if __name__ == '__main__':
    h = ListNode(2)
    for i in range(1):
        node = ListNode(i)
        node.next = h.next
        h.next = node
    s = Solution()
    # a = s.removeElements(h, 1)
    while h:
        print(h.val)
        h = h.next

