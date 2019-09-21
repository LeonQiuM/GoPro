package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

//  错误处理
/*

 */

func ProcessIo() error {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read faild: %v", err)
		}
		fmt.Print(string(r))
	}
	return nil
}
func main() {
	// err := ProcessIo()
	// if err != nil {
	// 	fmt.Print(err)
	// }

	var a3 = [3]int{1, 2, 3}
	a4 := [...]int{1, 23, 44, 5, 76, 7, 8}
	var a5 []string
	var a6 = [3]int{4, 5, 6}

	fmt.Println(a4, a5, a6)
	fmt.Println(reflect.TypeOf(a3))
	m := a3[0:2]
	fmt.Println(m)
	fmt.Println(reflect.TypeOf(m))

}
