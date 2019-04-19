package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	t1 := map[byte]int{}
	t2 := map[byte]int{}
	for i:=0; i<len(s1); i++ {
		if _, ok:= t1[s1[i]]; ok {
			t1[s1[i]] ++
		} else {
			t1[s1[i]] = 1
		}
		if _, ok:=t2[s2[i]]; ok {
			t2[s2[i]] ++
		} else {
			t2[s2[i]] = 1
		}
	}
	if isEqual(&t1, &t2) {
		return true
	}
	for i:=len(s1); i<len(s2); i++ {
		t2[s2[i-len(s1)]]--
		t2[s2[i]] ++
		if isEqual(&t1, &t2) {
			return true
		}
	}
	return false
}

func isEqual(t1 *map[byte]int, t2 *map[byte]int) bool {
	for key, _ := range *t1 {
		if (*t1)[key] != (*t2)[key]{
			return false
		}
	}
	return true
}

func main() {
	println(checkInclusion("rmqqh", "nrsqrqhrym"))
}
