import numpy as np
import matplotlib.pyplot as plt


def sin():
    x = np.arange(90, 100, 0.01)
    y = np.sin(x)

    print(x, y)
    plt.plot(x, y)
    plt.show()


def cos():
    x = np.arange(90, 100, 0.01)
    y = np.cos(x)

    print(x, y)
    plt.plot(x, y)
    plt.show()


def cos_sin():
    x = np.linspace(-2 * np.pi, 2 * np.pi, 256)
    c, s = np.cos(2*x), np.sin(x)
    plt.plot(x, c)
    plt.plot(x, s)
    plt.show()


if __name__ == '__main__':
    cos_sin()
