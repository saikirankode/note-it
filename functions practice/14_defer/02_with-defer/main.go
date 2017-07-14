package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("World")
}

func main() {
	defer world()
	hello()
}
