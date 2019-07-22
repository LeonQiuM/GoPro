package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	err := ProcessIo()
	if err != nil {
		fmt.Print(err)
	}
}
