package main

import (
	"bufio"
	"fmt"
	"os"
)

func process(lineStr string) (worldCount, spaceCount, numberCount, otherCount int) {
	strRune := []rune(lineStr)
	fmt.Println(strRune)
	for _, value := range strRune {
		switch {
		case (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z'):
			worldCount++
		case value == 32:
			spaceCount++
		case value >= '0' && value <= '9':
			numberCount++
		default:
			otherCount++
		}
	}
	return

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
	}
	worldCount, spaceCount, numberCount, otherCount := process(string(line))
	fmt.Printf("worldCount:%d\nspaceCount:%d\nnumberCount:%d\notherCount:%d\n", worldCount, spaceCount, numberCount, otherCount)
}
