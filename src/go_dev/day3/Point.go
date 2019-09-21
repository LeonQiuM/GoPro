package main

import "fmt"

//  指针

/*

 */

func main() {
	var a int = 5
	var b *int = &a
	fmt.Println(a, b, *b)

	var x int = 111
	fmt.Println(x)
	test(&x)
	fmt.Println(x)

}

func test(a *int) {
	*a = 100
	return
}
