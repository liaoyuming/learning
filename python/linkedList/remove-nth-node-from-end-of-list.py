class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        fast = head
        slow = head
        i = n
        while i>0:
            fast = fast.next
            i = i-1
        if fast is None:
            return head.next
        while fast.next:
            fast = fast.next
            slow = slow.next
        slow.next = slow.next.next
        return head


if __name__ == '__main__':
    h = ListNode(5)
    for i in range(5):
        node = ListNode(i)
        node.next = h.next
        h.next = node
    s = Solution()
    a = s.removeNthFromEnd(h, 3)
    while a:
        print(a.val)
        a = a.next

