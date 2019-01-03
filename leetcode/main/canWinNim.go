package main

/**
	leetcode https://leetcode-cn.com/problems/nim-game/description/
 */

func main()  {
	for i:=1; i < 100; i++ {
		println(i, "", canwinnim(i), "")
	}

}

func canwinnim(n int) bool {
	return n % 4 != 0
	//return canwin(n, true);
}

func canwin(n int, flag bool) bool {
	if (n <= 3) {
		return flag
	}
	if (flag) {
		return canwin(n-1, !flag) || canwin(n-2, !flag) || canwin(n-3, !flag)
	}
	return canwin(n-1, !flag) && canwin(n-2, !flag) && canwin(n-3, !flag)
}