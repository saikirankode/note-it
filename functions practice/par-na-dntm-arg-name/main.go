package main

import "fmt"

func average(xyz []float64) float64 {
	total := 0.0
	for _, v := range xyz {
		total += v
	}
	return total / float64(len(xyz))
}

func main() {
	data := []float64{25, 45, 89, 36, 106}
	y := average(data)
	fmt.Println(y)
}
