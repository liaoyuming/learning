class Node:
    def __init__(self, val, next, random):
        self.val = val
        self.next = next
        self.random = random

class Solution:
    def copyRandomList(self, head: 'Node') -> 'Node':
        if head is None:
            return None
        h = head
        res = Node(h.val, None, None)
        resHead = res
        d = {h: res}
        while h.next:
            h = h.next
            t = Node(h.val, None, None)
            res.next = t
            res = res.next
            d[h] = res
        while head:
            if head.random is None:
                d[head].random = None
            else:
                d[head].random = d[head.random]
            head = head.next
        return resHead