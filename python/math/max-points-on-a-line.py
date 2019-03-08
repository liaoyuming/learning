class Point:
    def __init__(self, a=0, b=0):
        self.x = a
        self.y = b

class Solution:
    def maxPoints(self, points):
        """
        :type points: List[Point]
        :rtype: int
        """
        length = len(points)
        if length <= 2:
            return length
        a = 0
        max = 2
        while a < length - max:
            f = dict(dict())
            b = a + 1
            # 标记是否所有点坐标都相同
            flag = True
            # 标记坐标相同点个数
            t = 0
            tmax = 0
            while b < length:
                if points[a].x == points[b].x and points[a].y == points[b].y:
                    t += 1
                    b += 1
                    continue
                elif points[a].x == points[b].x:
                    # 于x轴平行，特殊标记下
                    k = 'x'
                    j = points[a].x
                    flag = False
                else:
                    k = (points[a].y - points[b].y)*10000/(points[a].x-points[b].x)*10000
                    j = points[a].y - k * points[a].x
                    flag = False
                if k in f and j in f[k]:
                    f[k][j] += 1
                elif k in f:
                    f[k][j] = 2
                else:
                    f[k] = {j : 2}

                if f[k][j] > tmax:
                    tmax = f[k][j]
                b += 1
            tmax += t
            if flag:
                tmax += 1
            if tmax > max:
                max = tmax
            a += 1
        return max

if __name__ == '__main__':
    points = [[0,0],[94911151,9491115],[94911152,94911151]]
    for i in range(len(points)):
        points[i] = Point(points[i][0], points[i][1])
    s = Solution()
    print(s.maxPoints(points))