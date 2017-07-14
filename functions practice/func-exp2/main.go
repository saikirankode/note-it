package main

import "fmt"

func main() {
	half := func(a int) (int, bool) {
		return a / 2, a%2 == 0
	}
	add := func(x int, y int) int {
		return (x + y)
	}
	fmt.Println(half(2))
	fmt.Println(add(3, 4))
}
