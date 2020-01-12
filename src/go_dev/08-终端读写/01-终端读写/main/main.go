package main

import "fmt"

//
/*

 */

var (
	firstName, lastName string
	i                   int
	f                   float64
	s                   string
	input               = "56.12/2312/Go"
	format              = "%f/%d/%s"
)

func main() {
	fmt.Println("Please input your full name")
	fmt.Scanln(&firstName, &lastName)
	fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("hello %s %s\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("from the string get:", f, i, s)
}
