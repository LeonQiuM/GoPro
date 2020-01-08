package main

/*
计算1000以内的完数

*/

import "fmt"

func getDivisor(x int) []int {
	var divisorList []int
	for i := 1; i < x; i++ {
		if x%i == 0 {
			divisorList = append(divisorList, i)
		}
	}
	return divisorList
}

func main() {
	for i := 2; i <= 10000; i++ {
		divisorList := getDivisor(i)
		sum := 0
		for _, item := range divisorList {
			sum += item

		}
		// fmt.Println(sum)
		if sum == i {
			fmt.Println(i)
		}

	}
}
