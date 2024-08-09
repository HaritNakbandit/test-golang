package main

import "fmt"

func findTheOddInt(arr []int) int {
	counts := make(map[int]int)

	for _, num := range arr {
		counts[num]++
	}
	for num, count := range counts {
		if count%2 != 0 {
			return num
		}
	}

	return -1
}

func main() {
	fmt.Println("findTheOddInt 1 : ", findTheOddInt([]int{7}))
	fmt.Println("findTheOddInt 2 : ", findTheOddInt([]int{0}))
	fmt.Println("findTheOddInt 3 : ", findTheOddInt([]int{1, 1, 2}))
	fmt.Println("findTheOddInt 4 : ", findTheOddInt([]int{0, 1, 0, 1, 0}))
	fmt.Println("findTheOddInt 5 : ", findTheOddInt([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
}
