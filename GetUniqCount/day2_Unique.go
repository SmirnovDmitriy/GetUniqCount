package Smirnov

import "fmt"

func GetUniqCount(arr []int) int {
	uniq := 0
	x := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] && i != j {
				x++
			}
		}
		if x == 0 {
			uniq++
			//fmt.Println(uniq)
		}
		x = 0
	}
	return uniq
}

func main() {
	fmt.Println(GetUniqCount([]int{3, 8, 9, 7, 3, 5, 16, 8, 9, 6}))
}
