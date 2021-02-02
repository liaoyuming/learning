import numpy as np
from activefunc import sigmoid, softmax
from lostfunc import cross_entropy_error
from gradient import numerical_gradient
from mnist import load_mnist
import matplotlib.pyplot as plt


class TwoLayerNet:
    def __init__(self, input_size, hidden_size, output_size, weight_init_std=0.01):
        self.params = {}
        self.params['W1'] = weight_init_std * np.random.randn(input_size, hidden_size)
        self.params['b1'] = np.zeros(hidden_size)
        self.params['W2'] = weight_init_std * np.random.randn(hidden_size, output_size)
        self.params['b2'] = np.zeros(output_size)

    def predict(self, x):
        W1, W2 = self.params['W1'], self.params['W2']
        b1, b2 = self.params['b1'], self.params['b2']

        a1 = np.dot(x, W1) + b1
        z1 = sigmoid(a1)

        a2 = np.dot(z1, W2) + b2
        y = softmax(a2)

        return y

    def loss(self, x, t):
        y = self.predict(x)
        return cross_entropy_error(y, t)

    def accuracy(self, x, t):
        y = self.predict(x)
        y = np.argmax(y, axis=1)
        t = np.argmax(t, axis=1)

        accuracy = np.sum(y == t) / float(x.shape[0])
        return accuracy

    def numerical_gradient(self, x, t):
        loss_W = lambda W: self.loss(x, t)

        grads = {'W1': numerical_gradient(loss_W, self.params['W1']),
                 'b1': numerical_gradient(loss_W, self.params['b1']),
                 'W2': numerical_gradient(loss_W, self.params['W2']),
                 'b2': numerical_gradient(loss_W, self.params['b2'])}

        return grads


if __name__ == '__main__':

    (x_train, t_train), (x_test, t_test) = load_mnist(normalize=True, one_hot_label=True)

    train_loss_list = []

    iters_num = 10
    train_size = x_train.shape[0]
    batch_size = 100
    learning_rate = 0.1

    train_acc_list = []
    test_acc_list = []
    # iter_per_epoch = max(train_size / batch_size, 1)

    net = TwoLayerNet(input_size=784, hidden_size=50, output_size=10)

    for i in range(iters_num):
        batch_mask = np.random.choice(train_size, batch_size)
        x_batch = x_train[batch_mask]
        t_batch = t_train[batch_mask]

        grad = net.numerical_gradient(x_batch, t_batch)

        for key in ('W1', 'b1', 'W2', 'b2'):
            net.params[key] -= learning_rate * grad[key]

        loss = net.loss(x_batch, t_batch)
        train_loss_list.append(loss)

        # if i % iter_per_epoch == 0:
        train_acc = net.accuracy(x_train, t_train)
        test_acc = net.accuracy(x_test, t_test)
        train_acc_list.append(train_acc)
        test_acc_list.append(test_acc)

    print(train_loss_list)

    plt.plot(np.arange(len(train_loss_list)), train_loss_list)
    plt.show()

    plt.plot(np.arange(len(train_acc_list)), train_acc_list)
    plt.plot(np.arange(len(test_acc_list)), test_acc_list)
    plt.show()


