class ListNode:
    def __init__(self, val):
        self.val = val
        self.next = None

class MyLinkedList:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.head = None
        self.length = 0

    def get(self, index):
        """
        Get the value of the index-th node in the linked list. If the index is invalid, return -1.
        :type index: int
        :rtype: int
        """
        node = self.head
        i = 0
        while i < index and node is not None:
            node = node.next
            i += 1
        if node is None:
            return -1
        return node.val

    def addAtHead(self, val):
        """
        Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
        :type val: int
        :rtype: void
        """
        node = ListNode(val)
        node.next = self.head
        self.head = node
        self.length += 1

    def addAtTail(self, val):
        """
        Append a node of value val to the last element of the linked list.
        :type val: int
        :rtype: void
        """
        node = ListNode(val)
        if self.length == 0:
            self.head = node
        else:
            head = self.head
            while head.next:
                head = head.next
            head.next = node
        self.length += 1

    def addAtIndex(self, index, val):
        """
        Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
        :type index: int
        :type val: int
        :rtype: void
        """
        if index > self.length:
            return
        elif index == 0:
            return self.addAtHead(val)
        elif index == self.length:
            return self.addAtTail(val)
        node = ListNode(val)
        head = self.head
        i = 0
        while i < index-1 and head is not None:
            head = head.next
            i += 1
        node.next = head.next
        head.next = node
        self.length += 1


    def deleteAtIndex(self, index):
        """
        Delete the index-th node in the linked list, if the index is valid.
        :type index: int
        :rtype: void
        """
        if index >= self.length:
            return
        elif index == 0:
            self.head = self.head.next
            self.length -= 1
            return
        head = self.head
        i = 0
        while i < index-1 and head is not None:
            head = head.next
            i += 1
        head.next = head.next.next
        self.length -= 1

    def printList(self):
        head = self.head
        listStr = ''
        while head:
            listStr += '#' + str(head.val)
            head = head.next
        print(listStr)

if __name__ == '__main__':
    linkList = MyLinkedList()
    print(
        linkList.addAtHead(0),
        linkList.printList(),
        linkList.addAtIndex(1, 1),
        linkList.printList(),
        linkList.addAtTail(2),
        linkList.addAtTail(2),
        linkList.printList(),
        linkList.addAtIndex(4, 3),
        # linkList.printList(),
        # linkList.deleteAtIndex(2),
        linkList.printList(),
        linkList.get(5),
        linkList.printList())
