package main

import "fmt"

func greatest(abc ...int) int {
	var largest int

	for _, v := range abc {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	fmt.Println(greatest(56, 45, 68, 1564, 6, 19, 16))
}
