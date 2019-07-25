package main

import (
	"go_dev/day1/package_example/calc"
	"fmt"
)

func main()  {
	sum:= calc.Add(300,100)
	sub:= calc.Sub(300,100)
	fmt.Println(sum,sub)
}