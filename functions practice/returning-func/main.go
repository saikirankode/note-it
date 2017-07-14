package main

import "fmt"

func makeEvenGenerator() func() int {
	i := 0
	return func() int {
		i += 2
		return i
	}
}
func main() {
	na := makeEvenGenerator()
	fmt.Println(na())
	fmt.Println(na())
	fmt.Println(na())

	mas := makeEvenGenerator()
	fmt.Println(mas())
	fmt.Println(mas())
	fmt.Println(mas())
}
