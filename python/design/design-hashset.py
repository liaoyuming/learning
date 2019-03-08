class MyHashSet:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.hashList = [[] for i in range(2)]

    def add(self, key):
        """
        :type key: int
        :rtype: void
        """
        k = self._getHashKey(key)
        l = self.hashList[k]
        for i in range(len(l)):
            if l[i] == key:
                return
        l.append(key)
        self.hashList[k] = l

    def _getHashKey(self, key):
        return key % 2

    def remove(self, key):
        """
        :type key: int
        :rtype: void
        """
        k = self._getHashKey(key)
        l = self.hashList[k]
        j = len(l)
        i = 0
        while i < j:
            if l[i] == key:
                l.pop(i)
                j -= 1
            else:
                i += 1
        self.hashList[k] = l

    def contains(self, key):
        """
        Returns true if this set contains the specified element
        :type key: int
        :rtype: bool
        """
        k = self._getHashKey(key)
        l = self.hashList[k]
        for i in range(len(l)):
            if l[i] == key:
                return True
        return False

if __name__ == '__main__':
    hashSet = MyHashSet()
    print(hashSet.add(1),
          hashSet.add(2),
          hashSet.add(3),
          hashSet.contains(1),
          hashSet.contains(2),
          hashSet.contains(3),
          hashSet.remove(1),
          hashSet.remove(3),
          hashSet.contains(1),
          hashSet.contains(3),
    hashSet.contains(2))
