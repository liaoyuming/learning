package main

func addStrings(num1 string, num2 string) string {
	a := 0
	res := ""
	zero := "0"[0]
	t := 0
	for i, j:=len(num1)-1,len(num2)-1; i>=0 || j>=0; {
		if i < 0 {
			t = int(num2[j] - zero)
		} else if j < 0  {
			t = int(num1[i] - zero)
		} else {
			t = int(num1[i] - zero) + int(num2[j] - zero)
		}
		t += a
		a = t / 10
		res = string(t%10 + int(zero)) + res
		i--
		j--
	}
	if a != 0 {
		res = string(byte(a) + zero) + res
	}
	return res
}