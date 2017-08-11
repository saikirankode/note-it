package main

import (
	"fmt"
)

func main() {

	var mygreeting  = map[string]string{}
	mygreeting["kode"] = "hello"
	mygreeting["sai"] = "Good Morning"

	fmt.Println(mygreeting)
	fmt.Println(mygreeting == nil)
}
