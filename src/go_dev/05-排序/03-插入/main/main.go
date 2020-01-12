package main

import "fmt"

func main() {
	b := [...]int{1, 23, 543, 6, 8, 234, 77, 34, 767, 234}
	isort(b[:])
	fmt.Println(b)
}

func isort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[i], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
