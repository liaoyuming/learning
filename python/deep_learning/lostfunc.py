import numpy as np

# 损失函数

# 均方误差
def mean_squared_error(y, t):
    return 0.5 * np.sum((y-t)**2)

def test_mean_squared_error():
    # 将正确解标签表示为 1，其他标签表示为 0 的表示方法称 为 one-hot 表示
    # 此例中表示 2 是对的
    t = np.array([0, 0, 1, 0, 0, 0, 0, 0, 0, 0])
    y = np.array([0.1, 0.05, 0.6, 0.0, 0.05, 0.1, 0.0, 0.1, 0.0, 0.0])
    print(mean_squared_error(y, t))
    y = np.array([0.1, 0.05, 0.1, 0.0, 0.05, 0.1, 0.0, 0.6, 0.0, 0.0])
    print(mean_squared_error(y, t))


# 交叉熵误差
def cross_entropy_error(y, t):
    # 当出现 np.log(0) 时，np.log(0) 会变为负无限大 的 -inf，这样一来就会导致后续计算无法进行。作为保护性对策，添加一个 微小值可以防止负无限大的发生
    delta = 1e-7
    return -np.sum(t * np.log(y + delta))


def test_cross_entropy_error():
    t = np.array([2])
    y = np.array([[0.1, 0.05, 0.6, 0.0, 0.05, 0.1, 0.0, 0.1, 0.0, 0.0]])
    print(cross_entropy_error(y, t))
    t = np.array([2])
    y = np.array([[0.1, 0.05, 0.1, 0.0, 0.05, 0.1, 0.0, 0.6, 0.0, 0.0]])
    print(cross_entropy_error(y, t))


# def mini_batch_cross():




if __name__ == '__main__':

    test_mean_squared_error()
    test_cross_entropy_error()
