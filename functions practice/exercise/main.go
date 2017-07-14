package main

import "fmt"

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}
func main() {
	n, even := half(2)
	fmt.Println(n, even)
}
