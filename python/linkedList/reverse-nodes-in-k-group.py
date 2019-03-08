class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        cur = head
        if self.canRev(cur, k):
            revHead = None
            for _ in range(k):
                next = cur.next
                cur.next = revHead
                revHead = cur
                cur = next
            if cur is not None:
                head.next = self.reverseKGroup(cur, k)
            return revHead
        else:
            return head

    def canRev(self, head, k):
        i = 0
        while i < k:
            if head:
                head = head.next
                i += 1
            else:
                return False
        return True

    def printList(self, head):
        listStr = ''
        while head:
            listStr += '#' + str(head.val)
            head = head.next
        print(listStr)

if __name__ == '__main__':
    listArr = [1, 2, 3, 4];
    head = ListNode(None)
    t = head
    for i in range(len(listArr)):
        t.next = ListNode(listArr[i])
        t = t.next
    head = head.next
    s = Solution()
    s.printList(head)
    head = s.reverseKGroup(head, 3)
    s.printList(head)
