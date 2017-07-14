package main

import "fmt"

func average(xyz ...float64) float64 {
	total := 0.0
	for _, v := range xyz {
		total += v
	}
	return total / float64(len(xyz))
}

func main() {

	fmt.Println(average(89.36, 68.75, 106.32, 286.32, 1006.89))
}
