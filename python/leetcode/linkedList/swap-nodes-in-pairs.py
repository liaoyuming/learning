class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def swapPairs(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if head is None or head.next is None:
            return head
        h = ListNode(None)
        h.next = head
        pre = h
        while pre.next and pre.next.next:
            n = pre.next
            m = n.next
            l = m.next

            pre.next = m
            m.next = n
            n.next = l
            pre = n
        self.printList(h.next)
        return h.next

    def printList(self, head):
        listStr = ''
        while head:
            listStr += '#' + str(head.val)
            head = head.next
        print(listStr)


if __name__ == '__main__':
    l = ListNode(1)
    l.next = ListNode(2)
    l.next.next = ListNode(3)
    # l.next.next.next = ListNode(4)
    s = Solution()
    s.swapPairs(l)
