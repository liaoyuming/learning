class ListNode:
    def __init__(self, key, val):
        self.val = val
        self.key = key
        self.next = None
        self.pre = None

class LRUCache:

    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.capacity = capacity
        self.cacheMap = dict()
        self.cacheListHead = None
        self.cacheListLast = None

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        if key in self.cacheMap:
            self.__update(self.cacheMap[key])
            return self.cacheListHead.val
        return -1

    def __update(self, node):
        if node != self.cacheListHead:
            if node == self.cacheListLast:
                self.cacheListLast = self.cacheListLast.pre
                self.cacheListLast.next = None
            if node.pre is not None:
                node.pre.next = node.next
            if node.next is not None:
                node.next.pre = node.pre
            node.next = self.cacheListHead
            self.cacheListHead = node
            self.cacheListHead.pre = None
            if self.cacheListHead.next is not None:
                self.cacheListHead.next.pre = node

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: void
        """
        if key in self.cacheMap:
            self.cacheMap[key].val = value
            self.__update(self.cacheMap[key])
            return
        node = ListNode(key, value)
        self.__update(node)
        self.cacheMap[key] = node
        if self.cacheListLast is None:
            self.cacheListLast = node
        if len(self.cacheMap) > self.capacity:
            self.cacheMap.pop(self.cacheListLast.key)
            self.cacheListLast = self.cacheListLast.pre
            self.cacheListLast.next = None

    def printCache(self):
        mapStr = ''
        for i in self.cacheMap:
            mapStr += '|' + str(i) + ':' + str(self.cacheMap[i].val)
        print(mapStr)
        head = self.cacheListHead
        listStr = ''
        while head:
            listStr += '#' + str(head.val)
            head = head.next
        print(listStr)
# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)

if __name__ == '__main__':
    cache = LRUCache(3)
    print(
        cache.put(1, 1),
        cache.put(2, 2),
        cache.put(3, 3),
        cache.put(4, 4),
        cache.get(4),
        cache.get(3),
        cache.get(2),
        cache.get(1),
        cache.put(5, 5),
        cache.get(1),
        cache.get(2),
        cache.get(3),
        cache.get(4),
        cache.get(5))