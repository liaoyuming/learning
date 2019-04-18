package main

func isPowerOfTwo(n int) bool {
	t := 1
	for t <= n {
		if t == n {
			return true
		}
		t = t << 1
	}
	return false
}

func main() {
	println(isPowerOfTwo(4  ))
}