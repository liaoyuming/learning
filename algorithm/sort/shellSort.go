package main

func main() {

}

func shellSort(arr []int) []int {
	for gap:=len(arr)/2; gap>0; gap/=2 {
		for i:=0; i<gap; i++ {
			for j:=i+gap; j<len(arr); j+=gap {
				if arr[j] < arr[j-gap] {
				}
			}
		}
	}
	return arr
}