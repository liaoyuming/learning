package main

import "io"

// book: go 程序设计语言
// 练习7.2:
// 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，
// 返回一个新的Writer类型把原来的Writer封装在里面和一个表示写入新的Writer字节数的int64类型指针


type CountWriter struct {
	W io.Writer
	Counter int64
}

func Constructor(w *io.Writer) *CountWriter {
	return &CountWriter{W:w}
}

func(c *CountWriter) Write(p []byte) (n int, err error) {
	n, err = c.W.Write(p)
	c.Counter += int64(n)
	return n, err
}

func CountingWriter(w *io.Writer) (io.Writer, *int64) {
	countWriter := Constructor(w)
	return countWriter, &countWriter.Counter
}

