import numpy as np
import matplotlib.pyplot as plt


def sigmoid(x):
    return 1 / (1 + np.exp(-x))


def show_sigmoid():
    X = np.arange(-10.0, 10.0, 0.1)
    Y = sigmoid(X)
    plt.plot(X, Y)
    plt.ylim(-0.1, 1.1)
    plt.show()

def identity_function(x):
    return x


def softmax(a):
    c = np.max(a)
    exp_a = np.exp(a - c)
    sum_exp_a = np.sum(exp_a)
    y = exp_a / sum_exp_a
    return y

def show_softmax():
    X = np.arange(1.0, 10.0, 0.1)
    Y = softmax(X)
    plt.plot(X, Y)
    plt.ylim(-0.1, 1.1)
    plt.show()

if __name__ == '__main__':
    show_softmax()
