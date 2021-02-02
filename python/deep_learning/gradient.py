import numpy as np
from lostfunc import cross_entropy_error
from activefunc import softmax


# 梯度
def numerical_gradient(f, x):
    h = 1e-4  # 0.0001
    grad = np.zeros_like(x)  # 生成和x形状相同的数组
    it = np.nditer(x, flags=['multi_index'], op_flags=['readwrite'])
    while not it.finished:
        idx = it.multi_index
        tmp_val = x[idx]
        x[idx] = float(tmp_val) + h
        fxh1 = f(x)  # f(x+h)

        x[idx] = tmp_val - h
        fxh2 = f(x)  # f(x-h)
        grad[idx] = (fxh1 - fxh2) / (2 * h)

        x[idx] = tmp_val  # 还原值
        it.iternext()

    return grad


# 梯度下降法
def gradient_descent(f, init_x, lr=0.01, step_num=100):
    x = init_x
    for i in range(step_num):
        grad = numerical_gradient(f, x)
        x -= lr * grad
        print(grad, x)
    return x


def function_2(x):
    return x[0] ** 2 + x[1] ** 2


class simpleNet:
    def __init__(self):
        self.W = np.random.randn(2, 3)  # 正态分布

    def predict(self, x):
        return np.dot(x, self.W)

    def loss(self, x, t):
        z = self.predict(x)
        y = softmax(z)
        loss = cross_entropy_error(y, t)
        return loss


if __name__ == '__main__':
    init_x = np.array([-3.0, 4.0])
    print(gradient_descent(function_2, init_x=init_x, lr=0.1, step_num=100))


