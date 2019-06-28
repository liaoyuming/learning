class Solution:
    def nextGreatestLetter(self, letters, target):
        """
        :type letters: List[str]
        :type target: str
        :rtype: str
        """
        let = []
        for c in letters:
            let.append(ord(c))
        res = self.binarySearch(let, ord(target), 0, len(letters)-1)
        return chr(res)
            

    def binarySearch(self, letters, target, start, end):
        mid = int((end - start)/2) + start
        if letters[mid] <= target:
            if mid == len(letters) - 1:
                return letters[0]
            elif letters[mid+1] > target:
                return letters[mid+1]
            else:
                return self.binarySearch(letters, target, mid+1, end)
        else:
            if mid == 0:
                return letters[0]
            return self.binarySearch(letters, target, start, mid)





if __name__ == '__main__':
     s = Solution()
     print(s.nextGreatestLetter(["c", "f", "j"], "a"))