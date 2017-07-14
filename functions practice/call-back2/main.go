package main

import (
	"fmt"
)

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, v := range numbers {
		if callback(v) {
			xs = append(xs, v)
		}
	}
	fmt.Printf("%T\n", callback)
	return xs
}
func main() {
	xs := filter([]int{1, 2, 3, 4, 5, 6}, func(v int) bool {
		return v > 1
	})
	fmt.Println(xs)
}
