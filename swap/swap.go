package main

import "fmt"

func Swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}
func main() {
	fmt.Println(Swap(2, 3))
}
