package main

/*
输入三个整数x,y,z，请把这三个数由小到大输出


吧最小的放到x上
*/
import (
	"fmt"
)

func main() {
	var item_x, item_y, item_z float64
	var tmp float64
	fmt.Printf("使用空格隔开：")
	fmt.Scan(&item_x, &item_y, &item_z)
	fmt.Println(item_x, item_y, item_z)
	if item_x > item_y {
		tmp = item_x
		item_x = item_y
		item_y = tmp
	}
	if item_x > item_z {
		tmp = item_z
		item_z = item_x
		item_x = tmp
	}
	if item_y > item_z {
		tmp = item_z
		item_z = item_y
		item_y = tmp
	}
	fmt.Println(item_x, item_y, item_z)
}
