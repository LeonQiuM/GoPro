package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	var a int = 10
	switch a {
	case 0, 2, 3, 4: // 在任何一个满足，就 break， 可以写多个判断满足逻辑
		fmt.Println("a is equal 0")
	case 10:
		fmt.Println("a is equal 10")
		fmt.Println("match")
		fallthrough // 满足本次 case，继续向下执行
	default: // 前面都不满足，走 default
		fmt.Println("a is equal default")
	}
	guess()

}

func guess() {
	var n int
	rand.Seed(time.Now().Unix())
	n = rand.Intn(100)
	fmt.Println(n)

	for {
		var input int
		fmt.Scanf("%d\n", &input)
		fmt.Println(input, reflect.TypeOf(input))
		flag := false
		switch {
		case input == n:
			fmt.Println("You get thie")
			flag = true
		case input > n:
			fmt.Println("Too Upper")
		case input < n:
			fmt.Println("Too Lower")
		default:
			fmt.Println("Error in you input")
		}
		if flag {
			break
		}

	}

}
