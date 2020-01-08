package main

import "fmt"

func main() {
	b := [...]int{1, 23, 543, 6, 8, 234, 77, 34, 767, 234}
	bosrt(b[:])
	fmt.Println(b)
}

func bosrt(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] > a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}

}
