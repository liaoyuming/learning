import numpy as np
from tensorflow import keras
import matplotlib.pyplot as plt
import pickle
from activefunc import sigmoid, softmax


def _change_one_hot_label(X):
    T = np.zeros((X.size, 10))
    for idx, row in enumerate(T):
        row[X[idx]] = 1
    return T


def load_mnist(normalize=True, flatten=True, one_hot_label=False):
    (x_train, y_train), (x_test, y_test) = keras.datasets.mnist.load_data()

    if normalize:
        x_train = x_train.astype(np.float32) / 255.0
        x_test = x_test.astype(np.float32) / 255.0

    if flatten:
        x_train = x_train.reshape(*x_train.shape[:1], -1)
        x_test = x_test.reshape(*x_test.shape[:1], -1)

    if one_hot_label:
        y_train = _change_one_hot_label(y_train)
        y_test = _change_one_hot_label(y_test)

    return (x_train, y_train), (x_test, y_test)


def show_mnist():
    (x_train, t_train), (x_test, t_test) = load_mnist(flatten=False, normalize=False)
    img = x_train[0]

    plt.imshow(img)
    # show the plot
    plt.show()


def show():
    (x_train, y_train), (x_test, y_test) = keras.datasets.mnist.load_data()

    # Scale images to the [0, 1] range
    x_train = x_train.astype("float32") / 255
    x_test = x_test.astype("float32") / 255

    # Make sure images have shape (28, 28, 1)
    x_train = np.expand_dims(x_train, -1)
    x_test = np.expand_dims(x_test, -1)
    print("x_train shape:", x_train.shape)
    print(x_train.shape[0], "train samples")
    print(x_test.shape[0], "test samples")

    # convert class vectors to binary class matrices
    # y_train = keras.utils.to_categorical(y_train, num_classes)
    # y_test = keras.utils.to_categorical(y_test, num_classes)

    plt.subplot(221)
    plt.imshow(x_train[0], cmap=plt.get_cmap('gray'))
    # plt.subplot(222)
    # plt.imshow(x_train[1], cmap=plt.get_cmap('gray'))
    # plt.subplot(223)
    # plt.imshow(x_train[2], cmap=plt.get_cmap('gray'))
    # plt.subplot(224)
    # plt.imshow(x_train[3], cmap=plt.get_cmap('gray'))

    plt.subplot(221)
    plt.imshow(x_test[0], cmap=plt.get_cmap('gray'))
    # show the plot
    plt.show()


def init_network():
    with open('sample_weight.pkl', 'rb') as f:
        network = pickle.load(f)
    return network

def predict(network, x):

    W1, W2, W3 = network['W1'], network['W2'], network['W3']
    b1, b2, b3 = network['b1'], network['b2'], network['b3']
    a1 = np.dot(x, W1) + b1
    z1 = sigmoid(a1)
    a2 = np.dot(z1, W2) + b2
    z2 = sigmoid(a2)
    a3 = np.dot(z2, W3) + b3
    y = softmax(a3)
    return y

if __name__ == '__main__':
    (x_train, y_train), (x_test, y_test) = load_mnist()
    network = init_network()
    accuracy_cnt = 0
    for i in range(len(x_test)):
        y = predict(network, x_test[i])
        p = np.argmax(y)
        if p == y_test[i]:
            accuracy_cnt += 1
    print("accuracy: " + str(float(accuracy_cnt)/ len(x_test)))






