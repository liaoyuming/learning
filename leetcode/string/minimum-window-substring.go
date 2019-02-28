package main

func minWindow(s string, t string) string {
	tMap := make(map[uint8]int, len(t))
	target := make([]int, len(t))
	for k,_ := range t {
		if _, ok := tMap[t[k]]; !ok {
			tMap[t[k]] = 0
		}
	}
	min := 0
	max := 0
	matchLen := 0
	for k, _:= range s {
		if _, ok := tMap[s[k]]; ok {
			if matchLen < len(t) {
				if target[s[k]] != 0 {
					matchLen ++
					target[s[k]] --
				}
				if matchLen == 1 {
					min = k
				}
				if matchLen == len(t) {
					max = k
				}
			} else {
				tmin := 0
				tmax := 0
				for _, j:= range tMap {
					if tmin == 0 || j < tmin {
						tmin = j
					}
					if tmax == 0 || j > tmax {
						tmax = j
					}
				}
				if min == 0 || tmax - tmin < max - min{
					min = tmin
					max = tmax
				}
			}
		}
	}
	if min == 0 {
		return ""
	} else {
		return s[min:max+1]
	}
}

func main() {
	println(minWindow("of_characters_and_as", "aas"))
}