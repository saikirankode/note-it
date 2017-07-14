package main

import (
	"fmt"
)

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}

func wish(finame, laname string) {
	fmt.Println(finame, laname)
}

func main() {
	greet("saikiran", "kode")
	wish("sandeep", "maganti")
}
