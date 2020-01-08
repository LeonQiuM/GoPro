package main

import "fmt"

func main() {
	b := [...]int{1, 23, 543, 6, 8, 234, 77, 34, 767, 234}
	ssort(b[:])
	fmt.Println(b)
}

func ssort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

}
