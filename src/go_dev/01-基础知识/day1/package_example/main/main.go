package main

import (
	"fmt"
	"go_dev/day1/package_example/calc"
)

func main() {
	sum := calc.Add(300, 100)
	sub := calc.Sub(300, 100)
	fmt.Println(sum, sub)
}
