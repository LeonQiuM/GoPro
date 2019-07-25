package gorouter

func Add(a, b int, c chan int) {
	sum := a + b
	c <- sum
}
