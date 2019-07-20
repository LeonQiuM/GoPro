package main

import "fmt"

//  iota 常量生成器
/*
常量生成器：按照规律，然后自动为多个常量赋值

*/
func main() {
	const n = iota
	fmt.Println(n)

	type Weekday int

	const (
		Sunday Weekday = iota // 初始值为 0,后面依次加 1
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Saturday)

	const (
		A = iota + 3
		B
		C
		D
	)
	fmt.Println(D)

	const (
		_ = 1 << (10 * iota) //  1向左移动 10 位，变成了 1024
		KiB
		MiB
		GiB
	)
	fmt.Println(KiB)

	const (
		_ = 1024 * iota
		x
		y
		z
	)
	fmt.Println(x, y, z)

}
