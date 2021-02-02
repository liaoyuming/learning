import numpy as np


# 感知机

def AndGate(x1, x2):
    x = np.array([x1, x2])
    w = np.array([0.5, 0.5])
    b = -0.7
    tmp = np.sum(w * x) + b
    print(tmp)
    if tmp <= 0:
        return 0
    elif tmp > 0:
        return 1


def NAndGate(x1, x2):
    x = np.array([x1, x2])
    w = np.array([-0.5, -0.5])
    b = 0.7
    tmp = np.sum(w * x) + b
    print(tmp)
    if tmp <= 0:
        return 0
    elif tmp > 0:
        return 1


def ORGate(x1, x2):
    x = np.array([x1, x2])
    w = np.array([0.5, 0.5])
    b = 0
    tmp = np.sum(w * x) + b
    print(tmp)
    if tmp <= 0:
        return 0
    elif tmp > 0:
        return 1


def XORGate(x1, x2):
    s1 = NAndGate(x1, x2)
    s2 = ORGate(x1, x2)
    y = AndGate(s1, s2)
    return y


if __name__ == '__main__':
    assert AndGate(1, 1) == 1
    assert AndGate(1, 0) == 0
    assert AndGate(0, 1) == 0
    assert AndGate(0, 0) == 0

    assert NAndGate(1, 1) == 0
    assert NAndGate(0, 1) == 1
    assert NAndGate(1, 0) == 1
    assert NAndGate(0, 0) == 1

    assert ORGate(1, 1) == 1
    assert ORGate(1, 0) == 1
    assert ORGate(0, 1) == 1
    assert ORGate(0, 0) == 0

    assert XORGate(1, 1) == 0
    assert XORGate(0, 1) == 1
    assert XORGate(1, 0) == 1
    assert XORGate(0, 0) == 0
