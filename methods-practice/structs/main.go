package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func CircleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{x: 0, y: 0, r: 10}
	fmt.Println(CircleArea(c))

}
