package main

import (
	"fmt"
)

func equal(x, y []byte) bool {
	// 两个分片的比较
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func reverse(inputStr string) []byte {
	var newSlice []byte
	for i := 0; i < len(inputStr); i++ {
		newSlice = append(newSlice, inputStr[len(inputStr)-i-1])
	}
	strSlice := string(newSlice)
	fmt.Printf("new: %s\n", strSlice)
	// fmt.Printf(strings.Join([]string{strSlice}, "-"))
	return newSlice
}

func main() {
	var inputStr string
	fmt.Scanf("%s", &inputStr)
	fmt.Println(equal([]byte(inputStr), reverse(inputStr)))
}
