package main

import "io"

/**
go 程序设计语言
练习7.5：
io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，并且返回另一 个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。实现这个LimitReader函数:
func LimitReader(r io.Reader, n int64) io.Reader
 */

type LimitingReader struct {
	R io.Reader
	Max int64
	Counter int64
}

func(r *LimitingReader)Read(p []byte) (n int, err error) {
	if r.Counter + int64(len(p)) > r.Max {
		return r.R.Read(p[:r.Max-r.Counter])
	}
	return r.R.Read(p)
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitingReader{r, n, 0}
}