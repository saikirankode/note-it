package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, v := range numbers {
		callback(v)
	}
}

func main() {
	visit([]int{10, 20, 30, 40, 50, 60}, func(v int) {
		fmt.Println(v)
	})
}
